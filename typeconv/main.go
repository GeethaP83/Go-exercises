package main

import (
	"fmt"
	"math"
	"strconv"
)

func stringToInt(str string) int {
	// convert to int
	stringToInt, _ := strconv.Atoi(str)
	return stringToInt
}

func stringToFloat(str string) float64 {
	// convert to float with 4 digits of precision
	stringToFloat, _ := strconv.ParseFloat(str, 64)
	result := math.Floor(stringToFloat*10000) / 10000
	return result
}

func FloatToString(value float64) string {
	// convert to float with 2 digits of precision
	FloatToString := fmt.Sprintf("%.2f", value)
	return FloatToString
}

func main() {
	isFailed := false
	if stringToInt("10") != 10 {
		fmt.Println("Failed: stringToInt")
		isFailed = true
	}

	if stringToFloat("123.33333333333") != 123.3333 {
		fmt.Println("Failed: stringToFloat")
		isFailed = true
	}

	if FloatToString(1.0/3) != "0.33" {
		fmt.Println("Failed: FloatToString")
		isFailed = true
	}

	if !isFailed {
		fmt.Println("All tests passed")
	}
}
