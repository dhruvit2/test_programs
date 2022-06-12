package main

import (
	"errors"
	"fmt"
)

// map[string]int
// i - 1
// V - 5
// X = 10
// L = 50
// C = 100

// LC comes at 1st char then - x value

const (
	ROMAN_CHAR_I = "I"
	ROMAN_CHAR_V = "V"
	ROMAN_CHAR_X = "X"
	ROMAN_CHAR_L = "L"
	ROMAN_CHAR_C = "C"

	ROMAN_VAL_I = 1
	ROMAN_VAL_V = 5
	ROMAN_VAL_X = 10
	ROMAN_VAL_L = 50
	ROMAN_VAL_C = 100
)

var romanToNumberMap = make(map[string]int)

func ConvertTheRomanNumberToNumber(romanNumber string) (int, error) {
	resultVal := 0
	tmpPrevChar := ""
	for _, ch := range romanNumber {
		curChar := string(ch)
		numVal, ok := romanToNumberMap[curChar]
		if !ok {
			return 0, errors.New("Invalid roman number")
		}

		if tmpPrevChar == "" {
			resultVal += numVal
			tmpPrevChar = curChar
		} else {
			tmpVal := romanToNumberMap[curChar]
			if romanToNumberMap[string(tmpPrevChar)] < tmpVal {
				resultVal += (tmpVal - romanToNumberMap[string(tmpPrevChar)] - romanToNumberMap[string(tmpPrevChar)])
			} else {
				resultVal += tmpVal
			}
			tmpPrevChar = curChar
		}
	}

	return resultVal, nil
}

func InitTheRomanCharToNumber() {
	romanToNumberMap[ROMAN_CHAR_I] = ROMAN_VAL_I
	romanToNumberMap[ROMAN_CHAR_V] = ROMAN_VAL_V
	romanToNumberMap[ROMAN_CHAR_X] = ROMAN_VAL_X
	romanToNumberMap[ROMAN_CHAR_L] = ROMAN_VAL_L
	romanToNumberMap[ROMAN_CHAR_C] = ROMAN_VAL_C

}
func main() {
	inputRomanStr := "LIX"
	InitTheRomanCharToNumber()
	numVal, err := ConvertTheRomanNumberToNumber(inputRomanStr)
	if err != nil {
		fmt.Printf("Invalid roman string input\n")
	}

	fmt.Printf("[%v] = [%v] \n", inputRomanStr, numVal)
}
