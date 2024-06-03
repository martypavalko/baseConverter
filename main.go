package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func getInputBase(input string) string {
    // Determine input type: bin, hex, dec, base64
    // Binary input: 1010110
    // Hex input: bf43
    // Dec input: 34765

    conv, err := strconv.Atoi(input)
    if err != nil {
        // Must be bin or dec
        pattern, _ := regexp.Compile("[0-1]*")
        if ! pattern.MatchString(input) {
            return "not binary"
        }
    }
    // Else must be hex or base64

    return "AOK"
} 

func main() {
    fmt.Println("Base conversion")
}
