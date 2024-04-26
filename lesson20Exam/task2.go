package main

import (
	"fmt"
)

var ones = [...]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var teens = [...]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = [...]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func numberToWords(num int) string {
	if num == 0 {
		return "zero"
	}

	var result string

	if num >= 1000000000 {
		result += numberToWords(num/1000000000) + " billion "
		num %= 1000000000
	}

	if num >= 1000000 {
		result += numberToWords(num/1000000) + " million "
		num %= 1000000
	}

	if num >= 1000 {
		result += numberToWords(num/1000) + " thousand "
		num %= 1000
	}

	if num >= 100 {
		result += numberToWords(num/100) + " hundred "
		num %= 100
	}

	if num > 0 {
		if result != "" {
			result += "and "
		}
		if num < 10 {
			result += ones[num]
		} else if num < 20 {
			result += teens[num-10]
		} else {
			result += tens[num/10]
			if num%10 != 0 {
				result += " " + ones[num%10]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(numberToWords(128))
}
