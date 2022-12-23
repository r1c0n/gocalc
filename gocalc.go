package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var variables map[string]float64

func main() {
	fmt.Println("Welcome to the calculator!")
	variables = make(map[string]float64)

	for {
		// Read a line of input from the user
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		// Split the line into tokens
		tokens := strings.Split(line, " ")

		// Check the first token to determine the type of expression
		switch tokens[0] {
		case "def":
			// This is a variable definition
			if len(tokens) != 3 {
				fmt.Println("Invalid definition")
				continue
			}
			variables[tokens[1]] = parseNumber(tokens[2])
		case "fn":
			// This is a function definition
			if len(tokens) < 4 {
				fmt.Println("Invalid function definition")
				continue
			}
			variables[tokens[1]] = parseFunction(tokens[2:])
		default:
			// This is a simple expression
			result := parseExpression(tokens)
			fmt.Printf("Result: %.2f\n", result)
		}
	}
}

func parseNumber(str string) float64 {
	// Try to parse the string as a number
	num, err := strconv.ParseFloat(str, 64)
	if err == nil {
		// The string is a number, return it
		return num
	}

	// The string is not a number, check if it is a variable
	num, ok := variables[str]
	if ok {
		// The string is a variable, return its value
		return num
	}

	// The string is not a number or a variable, return 0
	return 0
}

func parseFunction(tokens []string) float64 {
	// Get the function name and arguments
	name := tokens[0]
	args := tokens[1:]

	// Check the function name and perform the appropriate calculation
	switch name {
	case "sin":
		// Calculate the sine of the first argument
		return math.Sin(parseNumber(args[0]))
	case "cos":
		// Calculate the cosine of the first argument
		return math.Cos(parseNumber(args[0]))
	case "mean":
		// Calculate the mean of all the arguments
		var sum float64
		for _, arg := range args {
			sum += parseNumber(arg)
		}
		return sum / float64(len(args))
	}

	// If the function name is not recognized, return 0
	return 0
}

func parseExpression(tokens []string) float64 {
	// Initialize the result to the first number in the expression
	result := parseNumber(tokens[0])

	// Iterate through the rest of the tokens, performing the appropriate operation
	for i := 1; i < len(tokens); i += 2 {
		op := tokens[i]
		num := parseNumber(tokens[i+1])
		switch op {
		case "+":
			result += num
		case "-":
			result -= num
		case "*":
			result *= num
		case "/":
			result /= num
		}
	}

	return result
}
