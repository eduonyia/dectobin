package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int64

	fmt.Print("Enter a Decimal Number: ")
	fmt.Scan(&num)

	result := strconv.FormatInt(num, 2)
	fmt.Print("Output: ", result)
	fmt.Println()
}
