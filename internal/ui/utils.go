package ui

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatGermanNumber(value float64, decimals int) string {
	formatted := fmt.Sprintf("%."+strconv.Itoa(decimals)+"f", value)

	// Replace decimal point with comma
	formatted = strings.Replace(formatted, ".", ",", 1)

	// Add thousand separators
	parts := strings.Split(formatted, ",")
	intPart := parts[0]

	// Add dots as thousand separators
	if len(intPart) > 3 {
		var result strings.Builder
		for i, digit := range intPart {
			if i > 0 && (len(intPart)-i)%3 == 0 {
				result.WriteString(".")
			}
			result.WriteRune(digit)
		}
		intPart = result.String()
	}

	if len(parts) > 1 {
		return intPart + "," + parts[1]
	}
	return intPart
}

func ParseGermanNumber(str string) (float64, error) {
	// Remove thousand separators (dots)
	cleaned := strings.ReplaceAll(str, ".", "")

	// Replace decimal comma with point
	cleaned = strings.Replace(cleaned, ",", ".", 1)

	return strconv.ParseFloat(cleaned, 64)
}

func FormatCurrency(value float64) string {
	return FormatGermanNumber(value, 2) + " €"
}

func FormatPercentage(value float64) string {
	return FormatGermanNumber(value, 1) + " %"
}

func FormatKilometers(value float64) string {
	return FormatGermanNumber(value, 0) + " km"
}

func FormatLiters(value float64) string {
	return FormatGermanNumber(value, 1) + " L"
}

func FormatKWh(value float64) string {
	return FormatGermanNumber(value, 1) + " kWh"
}

func FormatConsumption(value float64, unit string) string {
	return FormatGermanNumber(value, 1) + " " + unit + "/100km"
}
