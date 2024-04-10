package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"unicode"
)

func hexToDec(hexValue string) int {
    var sum int
    for i, char := range hexValue {
        var conversion float64 = 0
        if float64(char - '0') > 9 {
            switch unicode.ToUpper(char) {
            case 'A':
                conversion = 10
            case 'B':
                conversion = 11
            case 'C':
                conversion = 12
            case 'D':
                conversion = 13
            case 'E':
                conversion = 14
            case 'F':
                conversion = 15
            }
        } else {
            conversion = float64(char - '0')
        }

        sum += int(conversion * math.Pow(16, float64(len(hexValue)-1-i)))
    }

    return sum

}

func hexValidate (hexValue string) bool {
    for _, char := range hexValue {
        var readyToValidate = int(unicode.ToUpper(rune(char)))
        // log.Println(readyToValidate)
        if 48 > readyToValidate || (57 < readyToValidate && 65 > readyToValidate) || 70 < readyToValidate {
            fmt.Println("false")
            return false
        }
    }
   return true 
}

func decToHex (decValue int) string {
    var floatDecValue = float64(decValue)
    var hexValue float64
    var sb strings.Builder
    log.Printf("decValue / 16 = %f\n", floatDecValue/16)
    for math.Floor(math.Mod(floatDecValue, 16)) > 0 {
        hexValue = math.Floor(math.Mod(floatDecValue, 16))
        log.Printf("Floored hexvalue = %f", hexValue)
        switch hexValue {
        case 10:
            sb.WriteString("A")
        case 11:
            sb.WriteString("B")
        case 12:
            sb.WriteString("C")
        case 13:
            sb.WriteString("D")
        case 14:
            sb.WriteString("E")
        case 15:
            sb.WriteString("F")
        default:
            sb.WriteString(fmt.Sprintf("%.0f", hexValue))
        }
        decValue = decValue / 16
    }
    var hexOutput string
    for _,v := range sb.String(){
        hexOutput = string(v) + hexOutput
    }
    return hexOutput
}

func userPrompt() {
    var answer string
    fmt.Println("Please select a mode:")
    fmt.Println("[1] Convert hexadecimal value to decimal")
    fmt.Println("[2] Convert decimal value to hexadecimal")
    for {
        fmt.Scanln(&answer)
        if answer == "1" {
            var hexValue string
            fmt.Println()
            hexToDec(hexValue)
            break
        } else if answer == "2" {
            var decValue string
            decToHex(decValue)
            break
        } else {
            log.Fatalln("Invalid Input! Try Again!")
        }
    }
    
    
}

func main() {
    var hexValue string
    var output int
    fmt.Println("Please enter a hexadecimal value: ")
    for {
    fmt.Scanln(&hexValue)
    output = hexToDec(hexValue)
        if hexValidate(hexValue) {
            fmt.Printf("%d\n", output)
            break
        } else {
            fmt.Println("Input a valid hexadecimal value!  Try again!")
        }
    }
    fmt.Println(decToHex(float64(output)))
}
