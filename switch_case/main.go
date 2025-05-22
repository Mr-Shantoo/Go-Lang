package main

import (
	"fmt"
)
func main(){
	var day int
	fmt.Println("Enter a number (1-7):")
	fmt.Scanln(&day)
	switch day{
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("wednesday")
	case 6:
		fmt.Println("thursday")
	case 7:
		fmt.Println("Friday")
	}
//Fallthrough
grade:="B"
switch grade{
case "A":
	fmt.Println("Excellent")
	fallthrough
case "B":
	fmt.Println("very good")
	fallthrough
case "C":
	fmt.Println("good")
	fallthrough
case "D":
	fmt.Println("Needs Improvement")
default:
	fmt.Println("grade not recognized")
}
}