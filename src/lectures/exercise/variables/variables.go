//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favColour string = "green"
	fmt.Println("my favourite colour is ", favColour)

	birthYear, ageYears := 2002, 22
	fmt.Println("i was born in ", birthYear, ", i am ", ageYears, "years old")

	var (
		firstInitial = "E"
		lastInitial  = "K"
	)
	fmt.Println("my initials are ", firstInitial, lastInitial)

	var ageDays int
	ageDays = ageYears * 365
	fmt.Println("i've been alive for ", ageDays, "days.")
}
