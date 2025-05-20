package main

import (
	"fmt"
	"strconv"
)

func main() {
	var tempint int=10
	var tempfloat float64=float64(tempint)
	fmt.Println(tempfloat)
	//int to string 
	str:= strconv.Itoa(123)
	fmt.Println(str)
	//string to  int
	var str1 string="123"
	var num,_= strconv.Atoi(str1)
	fmt.Println(num)
	//string to float64
	var str2 string="12.30"
	var num2,_=strconv.ParseFloat(str2,64)
	fmt.Println(num2)
	
}