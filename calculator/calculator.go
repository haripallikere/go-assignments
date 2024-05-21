package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	// TODO: Implement Calculator function
	switch operand {
	case "+":
		return value1 + value2, nil
	case "-":
		return value1 - value2, nil
	case "*":
		return value1 * value2, nil
	case "/":
		if value2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return value1 / value2, nil
	default:
		return 0, fmt.Errorf("Operand not supported")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	// TODO: Check if exact 3 parts are given. If not, throw an error
	if len(parts) != 3 {
		fmt.Println("invalid")
		return
	}

	// TODO: Take all 3 values and pass to calculator function
	value1, err := strconv.ParseFloat(parts[0], 64)
	operand := parts[1]
	value2, err := strconv.ParseFloat(parts[2], 64)

	result, error := Calculator(value1, value2, operand)

	// TODO: Print results
	fmt.Printf("Result: %v\n", result, error)
}
