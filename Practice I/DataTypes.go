package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var second_s string
	// Read and save an integer, double, and String to your variables.
	scanner.Scan()

	if second_i, err := strconv.ParseUint(scanner.Text(), 10, 64); err == nil {
		fmt.Println(second_i + i)
	}
	scanner.Scan()

	if second_d, err := strconv.ParseFloat(scanner.Text(), 64); err == nil {
		fmt.Printf("%.1f\n", second_d+d)
	}
	scanner.Scan()
	second_s = scanner.Text()
	// Print the sum of both integer variables on a new line.

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + second_s)
}
