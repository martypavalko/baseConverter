package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

/*
   Convert from:
       binary <-> decimal
       hexadecimal <-> decimal
       binary <-> hexadecimal

       **BONUS**
       base64 encoding and decoding
*/

/*
Steps for Binary conversion:
  1. Core function which is conversion
  2. Take input
  3. Input validation
*/

func betterValidation (userInput string) (string, error) {
    // Return numeral system (base type) and errors

    // Take input, determine what base it is, validate the format is correct.

    // Determine what base it is

    // Binary: input only contains 1's and 0's
    binaryPattern, err := regexp.Compile("^[0-1]+$")
    if err != nil {
        return "", err
    }
    if binaryPattern.Match([]byte(userInput)) {
        return "binary", nil
    }
    // Decimal: input only contains 0-9
    decPattern, err := regexp.Compile("^[0-9]+$")
    if err != nil {
        return "", err
    }
    if decPattern.Match([]byte(userInput)) {
        return "decimal", nil
    }
    // Hexadecimal: input only contains 0-9 and a-f
    hexPattern, err := regexp.Compile("^[0-9a-fA-F]+$")
    if err != nil {
        return "", err
    }
    if hexPattern.Match([]byte(userInput)) {
        return "hexadecimal", nil
    }
    // TODO: Add base64 validation
    return "shouldn't have hit this", nil
}

func binaryToDecimal(binaryValue string) string {
    size := len(binaryValue)
    var sum float64
    for exp, bit := range(binaryValue) {
        if (bit != '0') {
            sum += math.Pow(2, float64(size-exp-1))
        }
    }
    return strconv.Itoa(int(sum))
}

func decimalToBinary(decimalValue string) string {
    return decimalValue
}

func main() {
    str, err := betterValidation("11111111")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(str)
}
