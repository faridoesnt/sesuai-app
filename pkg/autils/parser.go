package autils

import "strconv"

func FloatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 6, 64)
}
