package utils

import (
	"fmt"
	"strconv"
)

func round(num float64, decimalPlaces int) float64 {
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	str := fmt.Sprintf(format, num)
	newNum, _ := strconv.ParseFloat(str, 64)
	return newNum
}
