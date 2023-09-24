package libs

import (
	"bytes"
	"crypto/tls"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"mime/multipart"
	"net/http"
)

const (
	maxPartSize = int64(2048 * 1000)
	maxRetries  = 3
)

func AWSMultipartUpload(bucket, access_key, secret, key string, file multipart.File, info *multipart.FileHeader) (filename string, err error) {
	creds := credentials.NewStaticCredentials(access_key, secret, "")
	if _, err = creds.Get(); err != nil {
		logrus.Errorf("%s while getting aws credentials", err)
		return
	}

	sess := session.Must(session.NewSession())

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	svc := s3.New(sess, &aws.Config{
		HTTPClient:  client,
		Region:      aws.String(endpoints.ApSoutheast3RegionID),
		Credentials: creds,
	})

	buffer := make([]byte, info.Size)
	contentType := http.DetectContentType(buffer)
	if _, err = file.Read(buffer); err != nil {
		logrus.Errorf("%s while reading file", err)
		return
	}

	filename = key

	input := &s3.CreateMultipartUploadInput{
		Bucket:       aws.String(bucket),
		Key:          aws.String(filename),
		ContentType:  aws.String(contentType),
		ACL:          aws.String("public-read"),
		StorageClass: aws.String("STANDARD"),
	}

	resp := &s3.CreateMultipartUploadOutput{}

	resp, err = svc.CreateMultipartUpload(input)
	if err != nil {
		logrus.Errorf("%s while creating multipart upload request", err)
		return
	}

	var curr, partLength int64
	var remaining = info.Size
	var completedParts []*s3.CompletedPart

	partNumber := 1
	for curr = 0; remaining != 0; curr += partLength {
		if remaining < maxPartSize {
			partLength = remaining
		} else {
			partLength = maxPartSize
		}

		completedPart := &s3.CompletedPart{}
		completedPart, err = uploadPart(svc, resp, buffer[curr:curr+partLength], partNumber)
		if err != nil {
			logrus.Errorf("%s while uploading part", err)

			err = abortMultipartUpload(svc, resp)
			if err != nil {
				logrus.Errorf("%s while aborting multipart upload", err)
			}

			return
		}

		remaining -= partLength
		partNumber++
		completedParts = append(completedParts, completedPart)
	}

	_, err = completeMultipartUpload(svc, resp, completedParts)
	if err != nil {
		logrus.Errorf("%s while completeMultipartUpload ", err)
		return
	}

	return
}

func completeMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, completedParts []*s3.CompletedPart) (*s3.CompleteMultipartUploadOutput, error) {
	completeInput := &s3.CompleteMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
		MultipartUpload: &s3.CompletedMultipartUpload{
			Parts: completedParts,
		},
	}
	return svc.CompleteMultipartUpload(completeInput)
}

func uploadPart(svc *s3.S3, resp *s3.CreateMultipartUploadOutput, fileBytes []byte, partNumber int) (part *s3.CompletedPart, err error) {
	tryNum := 1
	partInput := &s3.UploadPartInput{
		Body:          bytes.NewReader(fileBytes),
		Bucket:        resp.Bucket,
		Key:           resp.Key,
		PartNumber:    aws.Int64(int64(partNumber)),
		UploadId:      resp.UploadId,
		ContentLength: aws.Int64(int64(len(fileBytes))),
	}

	for tryNum <= maxRetries {
		uploadResult, err := svc.UploadPart(partInput)
		if err != nil {
			logrus.Errorf("%s while uploading file part", err)

			if tryNum == maxRetries {
				if aerr, ok := err.(awserr.Error); ok {
					logrus.Errorf("%s maximum retry reached", aerr)
					return nil, aerr
				}

				return nil, err
			}

			tryNum++
		} else {
			return &s3.CompletedPart{
				ETag:       uploadResult.ETag,
				PartNumber: aws.Int64(int64(partNumber)),
			}, nil
		}
	}

	return nil, nil
}

func abortMultipartUpload(svc *s3.S3, resp *s3.CreateMultipartUploadOutput) (err error) {
	abortInput := &s3.AbortMultipartUploadInput{
		Bucket:   resp.Bucket,
		Key:      resp.Key,
		UploadId: resp.UploadId,
	}

	_, err = svc.AbortMultipartUpload(abortInput)
	if err != nil {
		logrus.Errorf("%s while aborting multipart upload", err)
		return
	}

	return
}
