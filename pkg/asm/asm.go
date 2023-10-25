package asm

import (
	"encoding/base64"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"log"
)

func GetSecret(secretName, region string) (data map[string]string, err error) {
	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		return nil, err
	}

	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString

		err = json.Unmarshal([]byte(secretString), &data)
		if err != nil {
			return nil, err
		}
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		decodedLen, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			log.Printf("base64 decode error: %v", err)
			return nil, err
		}

		err = json.Unmarshal(decodedBinarySecretBytes[:decodedLen], &data)
		if err != nil {
			return nil, err
		}
	}

	return data, err
}
