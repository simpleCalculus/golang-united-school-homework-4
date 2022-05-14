package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func isEmpty(str string) bool {
	return len(strings.Trim(str, " \t\n")) == 0
}

func StringSum(input string) (output string, err error) {
	if isEmpty(input) {
		return "", fmt.Errorf("error: %w", errorEmptyInput)
	}

	input = strings.Trim(input, " \t\n")
	var firstNumSign = 1
	if input[0] == '+' {
		input = input[1:]
	} else if input[0] == '-' {
		input = input[1:]
		firstNumSign = -1
	}

	var sum Sum
	if strings.Contains(input, "+") {
		sum = calculate(input, '+')
	} else if strings.Contains(input, "-") {
		sum = calculate(input, '-')
	} else {
		return "", fmt.Errorf("error: %w", errorNotTwoOperands)
	}

	if sum.err != nil {
		return "", sum.err
	}

	return strconv.Itoa(firstNumSign*sum.firstNum + sum.secondNum), nil
}

func getNumber(str string) (int, error) {
	str = strings.Trim(str, " \t\n")
	return strconv.Atoi(str)
}

func calculate(input string, operand byte) (sum Sum) {
	strNums := strings.Split(input, string(operand))
	if len(strNums) != 2 {
		sum.err = errorNotTwoOperands
		return
	}

	num1, err := getNumber(strNums[0])
	if err != nil {
		sum.err = fmt.Errorf("The first number is not a valid integer, error = %w", err)
		return
	}

	num2, err := getNumber(strNums[1])
	if err != nil {
		sum.err = fmt.Errorf("The first number is not a valid integer, error = %w", err)
		return
	}

	sum.firstNum = num1
	if operand == '-' {
		sum.secondNum = -1 * num2
	} else {
		sum.secondNum = num2
	}
	return sum
}

type Sum struct {
	firstNum  int
	secondNum int
	err       error
}
