package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NAMO = "not a math operation(write the expression with whitespaces between numbers and operands)"
	DIFF = "different numeric styles"
	TMO  = "too many operands"
	NEG  = "roman numerals cant be negative"
	ZERO = "0 in roman numeric style doesnt exist"
	BORD = "enter the numbers from 1 to 10 or form I to X"
)

var roman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}

var arab = []int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

var ss []string
var a, b int
var err, err1 error

func ssConv(ss []string) {
	a, err = strconv.Atoi(ss[0])
	if err != nil {
		panic(err)
	}
	if a < 1 || a > 10 {
		panic(BORD)
	}
	b, err1 = strconv.Atoi(ss[2])
	if err1 != nil {
		panic(err1)
	}
	if b < 1 || b > 10 {
		panic(BORD)
	}
}

func arabConv(result int) {
	var romanNumber string
	if result == 0 {
		panic(ZERO)
	} else if result < 0 {
		panic(NEG)
	}
	for result > 0 {
		for _, d := range arab {
			for i := d; i <= result; {
				for index, value := range roman {
					if value == d {
						romanNumber += index
						result -= d
					}
				}
			}
		}
	}
	fmt.Print(romanNumber)
}

func main() {

	fmt.Print("enter the expression for calculation using roman or arabic numbers from 1 to 10 and operators + - * / with whitespaces:  ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	ss = strings.Fields(strings.ToUpper(s))

	if len(ss) < 3 {
		panic(NAMO)
	}
	if len(ss) > 3 {
		panic(TMO)
	}

	defer func() {
		if x := recover(); x != nil {
			switch ss[0] {
			case "I":
				a = 1
			case "II":
				a = 2
			case "III":
				a = 3
			case "IV":
				a = 4
			case "V":
				a = 5
			case "VI":
				a = 6
			case "VII":
				a = 7
			case "VIII":
				a = 8
			case "IX":
				a = 9
			case "X":
				a = 10
			default:
				if a < 1 || a > 10 {
					panic(BORD)
				}
				panic(DIFF)
			}

			switch ss[2] {
			case "I":
				b = 1
			case "II":
				b = 2
			case "III":
				b = 3
			case "IV":
				b = 4
			case "V":
				b = 5
			case "VI":
				b = 6
			case "VII":
				b = 7
			case "VIII":
				b = 8
			case "IX":
				b = 9
			case "X":
				b = 10
			default:
				if b < 0 || b > 10 {
					panic(BORD)
				}
				panic(DIFF)
			}

			switch ss[1] {
			case "+":
				arabConv(a + b)
			case "-":
				arabConv(a - b)
			case "*":
				arabConv(a * b)
			case "/":
				arabConv(a / b)
			default:
				panic(NAMO)
			}
		}
	}()

	ssConv(ss)

	switch ss[1] {
	case "+":
		fmt.Print(a + b)
	case "-":
		fmt.Print(a - b)
	case "*":
		fmt.Print(a * b)
	case "/":
		fmt.Print(a / b)
	default:
		panic(NAMO)
	}
}
