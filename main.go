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

    _, err := strconv.Atoi(input)
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

func inputTester() {
    test1 := getInputBase("10001")
    test2 := getInputBase("1234")
    test3 := getInputBase("b34f")

    fmt.Printf("Bin: %s\n",test1)
    fmt.Printf("Dec: %s\n",test2)
    fmt.Printf("Hex: %s\n",test3)

}

func main() {
    fmt.Println("Base conversion")
    inputTester()
}
