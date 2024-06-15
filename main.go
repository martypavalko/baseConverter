package main

import (
	"errors"
	"fmt"
	"os"

	// "log"
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

func validateInput(userInput, conversionType string) (bool, error){
    var invalid bool = false
    var err error
    if conversionType == "binary" {
        for r := range(userInput) {
            if r != '0' && r != '1' {
                invalid = true
                err = errors.New("Invalid binary string")
            }
        }
    }
    if invalid {
        return false, err
    }
    return true, nil
}

// CURRENTLY RETURNING ERROR
func binaryToDecimal(binaryValue string) float64 {
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
    return sum
}

func main() {
    fmt.Println(binaryToDecimal("010001"))
    fmt.Println(binaryToDecimal("110000"))
    fmt.Println(binaryToDecimal("11100001"))
    fmt.Println(binaryToDecimal("000000011111"))
}
