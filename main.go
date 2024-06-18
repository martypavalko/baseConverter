package main

import (
	"errors"
	"fmt"
	"math"
	"os"
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

// TODO: Change return type from bool to float64 so that the output can be usable by whatever function receives it.
func validateInput(userInput, conversionType string) (bool, error) {
    var invalid bool = false
    var err error
    if conversionType == "binary" {
        for _, r := range(userInput) {
            if r != '0' && r != '1' {
                invalid = true
                err = errors.New("Invalid binary value")
            }
        }
    } else if conversionType == "decimal" {
        _, err := strconv.Atoi(userInput)
        if err != nil {
            invalid = true
            err = errors.New("Invalid decimal value")
        }
    }
    if invalid {
        return false, err
    }
    return true, nil
}

func betterValidation (userInput string) (string, error) {
    // Return numeral system (base type) and errors

    // Take input, determine what base it is, validate the format is correct.
    // Use Regex
    // If format is correct return, the original value and nil, if format is incorrect, return nil and the error 

    // Determine what base it is
    // Binary: input only contains 1's and 0's
    binaryPattern, err := regexp.Compile("^[0-1]+$")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    if binaryPattern.Match([]byte(userInput)) {
        return "binary", nil
    }
    // Hexadecimal: input only contains 0-9 and a-f
    hexPattern, err := regexp.Compile("^[0-9a-fA-F]+$")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    if hexPattern.Match([]byte(userInput)) {
        return "hexadecimal", nil
    }
    // Decimal: input only contains 0-9

    // Base64 -> I have no fucking clue
    return "shouldn't have hit this", nil
}

// TODO: Move all parts of input validation to the validateInput function and have it automatically detect what type was entered.

func binaryToDecimal(binaryValue string) string {
    isValid, err := validateInput(binaryValue, "binary")
    if err != nil || !isValid {
        fmt.Println(err)
        os.Exit(1)
    }
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
    isValid, err := validateInput(decimalValue, "decimal")
    if err != nil || !isValid {
        fmt.Println(err)
        os.Exit(1)
    }
    return "Binary!"
}

func main() {
    // fmt.Println(binaryToDecimal("010001"))
    // fmt.Println(binaryToDecimal("110000"))
    // fmt.Println(binaryToDecimal("11100001"))
    // fmt.Println(binaryToDecimal("000000011111"))
    str, err := betterValidation("11111111")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(str)
}
