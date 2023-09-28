package helpers

import "time"

func GetHoroscope(t time.Time) string {
	month := t.Month()
	date := t.Day()

	if (month == time.March && date >= 21) || (month == time.April && date <= 19) {
		return "Aries"
	} else if (month == time.April && date >= 20) || (month == time.May && date <= 20) {
		return "Taurus"
	} else if (month == time.May && date >= 21) || (month == time.June && date <= 20) {
		return "Gemini"
	} else if (month == time.June && date >= 21) || (month == time.July && date <= 22) {
		return "Cancer"
	} else if (month == time.July && date >= 23) || (month == time.August && date <= 22) {
		return "Leo"
	} else if (month == time.August && date >= 23) || (month == time.September && date <= 22) {
		return "Virgo"
	} else if (month == time.September && date >= 23) || (month == time.October && date <= 22) {
		return "Libra"
	} else if (month == time.October && date >= 23) || (month == time.November && date <= 21) {
		return "Scorpio"
	} else if (month == time.November && date >= 22) || (month == time.December && date <= 21) {
		return "Sagitarius"
	} else if (month == time.December && date >= 22) || (month == time.January && date <= 19) {
		return "Capricorn"
	} else if (month == time.January && date >= 20) || (month == time.February && date <= 18) {
		return "Aquarius"
	} else {
		return "Pisces"
	}
}
