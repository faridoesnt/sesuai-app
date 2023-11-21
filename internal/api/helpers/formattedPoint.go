package helpers

import (
	"Sesuai/internal/api/entities"
	"strings"
)

func FormattedPoint(results []entities.Result) []entities.Result {
	for index := range results {
		if strings.Contains(results[index].Point, ".") {
			parts := strings.Split(results[index].Point, ".")

			if len(parts) == 2 {
				decimalPart := parts[1]

				if len(decimalPart) > 1 {
					decimalPart = decimalPart[:1]
				}

				results[index].Point = parts[0] + "." + decimalPart
			}
		}
	}

	return results
}
