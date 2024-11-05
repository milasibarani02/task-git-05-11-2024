package main

import (
	"fmt"
	"os"
	"strconv"
	"task-git/utils"
)

func main() {
	// Check for command-line arguments
	n := 3 // Default value
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1]) // Convert first argument to int
		if err != nil {
			fmt.Println("Invalid number, using default value 3.")
			n = 3
		}
	}

	// Print results of functions
	fmt.Println("MagicSum:", utils.MagicSum(n))
	fmt.Println("MagicPow:", utils.MagicPow(n))
	fmt.Println("Magicodd:", utils.Magicodd(n))
	fmt.Println("MagicGrade:", utils.MagicGrade(n))
	fmt.Println("MagicName:", utils.MagicName(n))
	fmt.Println("MagicTria:", utils.MagicTria(n))

	// MagicChange example
	x := n
	fmt.Println("Before MagicChange:", x)
	utils.MagicChange(&x)
	fmt.Println("After MagicChange:", x)

	// Using MagicNumber
	magicNum := utils.MagicNumber{Number: n}
	fmt.Println("Before Multiply:", magicNum.Number)
	magicNum.Multiply(n)
	fmt.Println("After Multiply:", magicNum.Number)
}
