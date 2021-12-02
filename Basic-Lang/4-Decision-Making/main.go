package main

import "fmt"

func main() {
	ifSample()
	ifElseSample()
	nestedIfSample()
	switchSample()
}

func ifSample() {
	/* local variable definition */
	var a int = 10

	/* check the boolean condition using if statement */
	if a < 20 {
		/* if condition is true then print the following */
		fmt.Printf("a is less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}

func ifElseSample() {
	/* local variable definition */
	var a int = 100

	/* check the boolean condition */
	if a < 20 {
		/* if condition is true then print the following */
		fmt.Printf("a is less than 20\n")
	} else {
		/* if condition is false then print the following */
		fmt.Printf("a is not less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)

	/* check the boolean condition */
	if a == 10 {
		/* if condition is true then print the following */
		fmt.Printf("Value of a is 10\n")
	} else if a == 20 {
		/* if else if condition is true */
		fmt.Printf("Value of a is 20\n")
	} else if a == 30 {
		/* if else if condition is true  */
		fmt.Printf("Value of a is 30\n")
	} else {
		/* if none of the conditions is true */
		fmt.Printf("None of the values is matching\n")
	}
	fmt.Printf("Exact value of a is: %d\n", a)
}

func nestedIfSample() {
	/* local variable definition */
	var a int = 100
	var b int = 200

	/* check the boolean condition */
	if a == 100 {
		/* if condition is true then check the following */
		if b == 200 {
			/* if condition is true then print the following */
			fmt.Printf("Value of a is 100 and b is 200\n")
		}
	}
	fmt.Printf("Exact value of a is : %d\n", a)
	fmt.Printf("Exact value of b is : %d\n", b)
}

func switchSample() {
	/* local variable definition */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)
}
