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

type RomanToInt struct {
	romanToNumberMap map[string]int
}

func (ri RomanToInt) ConvertTheRomanNumberToNumber(romanNumber string) (int, error) {
	resultVal := 0
	tmpPrevChar := ""
	for _, ch := range romanNumber {
		curChar := string(ch)
		numVal, ok := ri.romanToNumberMap[curChar]
		if !ok {
			return 0, errors.New("Invalid roman number")
		}

		if tmpPrevChar == "" {
			resultVal += numVal
			tmpPrevChar = curChar
		} else {
			tmpVal := ri.romanToNumberMap[curChar]
			if ri.romanToNumberMap[string(tmpPrevChar)] < tmpVal {
				resultVal += (tmpVal - ri.romanToNumberMap[string(tmpPrevChar)] - ri.romanToNumberMap[string(tmpPrevChar)])
			} else {
				resultVal += tmpVal
			}
			tmpPrevChar = curChar
		}
	}

	return resultVal, nil
}

func (ri *RomanToInt) InitTheRomanCharToNumber() {
	ri.romanToNumberMap = make(map[string]int)

	ri.romanToNumberMap[ROMAN_CHAR_I] = ROMAN_VAL_I
	ri.romanToNumberMap[ROMAN_CHAR_V] = ROMAN_VAL_V
	ri.romanToNumberMap[ROMAN_CHAR_X] = ROMAN_VAL_X
	ri.romanToNumberMap[ROMAN_CHAR_L] = ROMAN_VAL_L
	ri.romanToNumberMap[ROMAN_CHAR_C] = ROMAN_VAL_C

}

func main() {
	inputRomanStr := "LIX"
	ri := RomanToInt{}
	ri.InitTheRomanCharToNumber()
	numVal, err := ri.ConvertTheRomanNumberToNumber(inputRomanStr)
	if err != nil {
		fmt.Printf("Invalid roman string input\n")
	}

	fmt.Printf("[%v] = [%v] \n", inputRomanStr, numVal)
}
