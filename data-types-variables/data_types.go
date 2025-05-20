package main

import (
	"fmt"
)
func main(){
	fmt.Println("Hello world")
	var name1 string="maine uddin shanto"
	fmt.Println("my name is",name1)
	//short variable declaration
	name:="maine uddin shanto"
	age:=22
	cgpa:=1.50
	fmt.Println("may name is",name," age=",age,"cgpa=",cgpa)
	//printf doesn't add a next line...
	fmt.Printf("age = %v", age) //%v is the place holder..
	fmt.Println("cgpa =",cgpa)


	//zero values
	var name2 string
	fmt.Println(name2)//null
	var count int
	fmt.Println("count",count)//0
	var price float64
	fmt.Printf("price %.1f\n",price)//0.0
	var insured bool
	fmt.Println("insured",insured)//false

	




}