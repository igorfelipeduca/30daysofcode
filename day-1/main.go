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

	var (
		integer int64
		floatN  float64
		stringV string
	)

	// Read and save an integer, double, and String to your variables.

	fmt.Print("Insert an integer: ")
	scanner.Scan()
	integer, _ = strconv.ParseInt(scanner.Text(), 10, 32)

	fmt.Print("Insert a float: ")
	scanner.Scan()
	floatN, _ = strconv.ParseFloat(scanner.Text(), 64)

	fmt.Print("Insert a string: ")
	scanner.Scan()
	stringV = scanner.Text()

	// Print the sum of both integer variables on a new line.

	fmt.Println(int64(integer) + int64(i))

	// Print the sum of the double variables on a new line.

	fmt.Printf("%.1f\n", floatN+d)

	// Concatenate and print the String variables on a new line

	fmt.Println(s + stringV)

	// The 's' variable above should be printed first.

}
