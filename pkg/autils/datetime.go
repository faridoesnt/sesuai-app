package autils

import (
	"Sesuai/internal/api/constants"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

func ParseTo24HourFormat(datetime string) string {
	result := datetime

	formattedHour := ""
	if strings.Contains(datetime, "PM") {
		addColon := false
		hour24format := 00
		formattedHour = datetime[11:13]

		if strings.Contains(formattedHour, ":") {
			addColon = true
			formattedHour = string(formattedHour[0])

			if tmp, err := strconv.Atoi(formattedHour); err == nil {
				hour24format = tmp + 12
			} else {
				logrus.Errorf("%s while parsing absent formatted hour to int", err)
			}
		} else {
			if tmp, err := strconv.Atoi(formattedHour); err == nil {
				hour24format = tmp + 12
			} else {
				logrus.Errorf("%s while parsing absent formatted hour to int", err)
			}
		}

		if hour24format == 24 {
			formattedHour = "12"
		} else if hour24format > 9 {
			formattedHour = strconv.Itoa(hour24format)
		} else {
			formattedHour = "0" + strconv.Itoa(hour24format)
		}

		if addColon {
			formattedHour += ":"
		}
	} else if strings.Contains(datetime, "AM") {
		formattedHour = datetime[11:13]

		if formattedHour == "12" {
			formattedHour = "00"
		}
	}

	if formattedHour != "" {
		result = datetime[0:11] + formattedHour + datetime[13:19]
	}

	return strings.TrimSpace(result)
}

func GetIntervalDeviceTime(deviceTime time.Time) (status string, message string) {
	deviceTime = deviceTime.Add(time.Hour * -7)
	duration := time.Since(deviceTime)
	interval := duration.Minutes()

	status = constants.SESSION_VALID
	if interval > 8 || interval < -8 {
		message = "invalid datetime with interval " + FloatToString(interval)
		status = constants.SESSION_INVALID
	}

	return
}

func GetDateTime(timeZone *time.Location) time.Time {
	if timeZone == nil {
		timeZone = DefaultTimeZone()
	}

	return time.Now().In(timeZone)
}

func DefaultTimeZone() *time.Location {
	tz := time.FixedZone("UTC+7", 7*60*60)
	if tz == nil {
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err == nil {
			tz = loc
		}
	}

	return tz
}
