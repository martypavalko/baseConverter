package main

import (
	"errors"
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
    // BUG: This is not validating hex properly.  It needs to parse better

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
    return "", errors.New("Not a valid format!")
}

func binaryToDecimal(userInput string) string {
    size := len(userInput)
    var sum float64
    for exp, bit := range(userInput) {
        if (bit != '0') {
            sum += math.Pow(2, float64(size-exp-1))
        }
    }
    return strconv.Itoa(int(sum))
}

func decimalToBinary(userInput string) (string, error) {
    decimalValue, err := strconv.Atoi(userInput)
    if err != nil {
        return "", err
    }
    var binaryString string
    for {
        originalDecValue := decimalValue
        decimalValue = decimalValue / 2
        remainder := math.Mod(float64(originalDecValue), 2)
        binaryString += strconv.Itoa(int(remainder))
        if decimalValue == 0 {
            break
        }
    }
    binaryRunes := []rune(binaryString)
    for i := 0; i < len(binaryString) / 2; i++ {
        binaryRunes[i], binaryRunes[len(binaryString)-1-i] = binaryRunes[len(binaryString)-1-i], binaryRunes[i] 
    }
    return string(binaryRunes), nil
}

func hexadecimalToDecimal (userInput string) (string, error) {
    var hexBitsSeparated []float64
    var sum float64 = 0
    for _, r := range(userInput) {
        // Below should ensure everything is converted to upper case
        if int(r) > 70 {
            r = rune(int(r) - 32)
        }
        switch r {
        case 65:
        hexBitsSeparated = append(hexBitsSeparated, 10)
        case 66:
        hexBitsSeparated = append(hexBitsSeparated, 11)
        case 67:
        hexBitsSeparated = append(hexBitsSeparated, 12)
        case 68:
        hexBitsSeparated = append(hexBitsSeparated, 13)
        case 69:
        hexBitsSeparated = append(hexBitsSeparated, 14)
        case 70:
        hexBitsSeparated = append(hexBitsSeparated, 15)
        default:
        hexBitsSeparated = append(hexBitsSeparated, float64(r - '0'))
        }
    }
    for i, hexBits := range hexBitsSeparated {
        fmt.Printf("\nindex: %v, bit: %.0f\n", i, hexBits)
        fmt.Printf("Exp: %v\n", len(hexBitsSeparated)-i-1)
        sum += hexBits * math.Pow(16, float64(len(hexBitsSeparated)-i-1))
        fmt.Printf("Sum: %v\n", sum)
    }
    return strconv.Itoa(int(sum)), nil
}

func main() {
    // fmt.Printf("This is what the hex letters look like with int casting:\n- A: %v\n-a: %v\n-F: %v\n-f: %v\n-r-'0': %v\n", int('A'), int('a'), int('F'), int('f'), int('9'-'0'))
    // TODO: Ask the user what type of conversion to perform
    // Because the validation cannot tell the difference between dec and hex when dealing with only 0-9
    // Maybe accept them as cmdline args when calling the program

    fmt.Println("Enter the number you want to convert")
    var userInput string
    fmt.Scanln(&userInput)
    _, err := betterValidation(userInput)
    if err != nil {
        fmt.Println(err)
    }
    // fmt.Println(base)
    switch base {
    case "decimal":
        conversion, err := decimalToBinary(userInput)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(conversion)
    case "binary":
        conversion := binaryToDecimal(userInput)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(conversion)
    case "hexadecimal":
        conversion, err := hexadecimalToDecimal(userInput)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(conversion)
    }

}
