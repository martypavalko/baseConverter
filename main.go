package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"math"
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
    // Binary return type: string
    // Decimal return type: int/float
    // Hexadecimal return type: string
    // Base64 return type: string

    // Take input, determine what base it is, validate the format is correct.
    // If format is correct return, the original value and nil, if format is incorrect, return nil and the error 
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
    fmt.Println(binaryToDecimal("010001"))
    fmt.Println(binaryToDecimal("110000"))
    fmt.Println(binaryToDecimal("11100001"))
    fmt.Println(binaryToDecimal("000000011111"))
}
