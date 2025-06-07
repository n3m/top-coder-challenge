// main.go
// CLI app: takes 3 args (int, int, float64), outputs a single float64 result.
//
// Usage: go run main.go <int1> <int2> <float1>
//
// Example: go run main.go 2 3 4.5

package main

import (
	"fmt"
	"os"
	"strconv"
)

// main is the entry point for the CLI application.
func main() {
	// Check for correct number of arguments.
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <int1> <int2> <float1>\n", os.Args[0])
		os.Exit(1)
	}

	// Parse first integer argument.
	int1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: first argument must be an integer\n")
		os.Exit(1)
	}

	// Parse second integer argument.
	int2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: second argument must be an integer\n")
		os.Exit(1)
	}

	// Parse third float argument.
	float1, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: third argument must be a float\n")
		os.Exit(1)
	}

	// Example operation: multiply all three values.
	result := float64(int1) * float64(int2) * float1

	// Output the result as a float.
	fmt.Printf("%f\n", result)
}
