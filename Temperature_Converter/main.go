package main

import (
	"fmt"
	"strings"
)
func main(){
	var temp float64
	var fromUnit,toUnit string
	fmt.Println("Enter Temperature:")
	fmt.Scanln(&temp)
	fmt.Println("Enter the current Unit(C,F,K):")
	fmt.Scanln(&fromUnit)
	fromUnit=strings.ToUpper(fromUnit)
	fmt.Println("Enter the UNit you want to conver into:")
	fmt.Scanln(&toUnit)
	toUnit=strings.ToUpper(toUnit)
	convertedTemp:=0.00
	switch fromUnit{
	case "C":
		if toUnit=="F"{
			convertedTemp=temp*9/5+32  
		}else if toUnit=="K"{
			convertedTemp=temp+273.15
		}
	case "F":
		if toUnit=="C"{
			convertedTemp=(temp-32)*5/9
		}else if toUnit=="K"{
			convertedTemp=temp+273.15
		}
	case "K":
		if toUnit=="C"{
			convertedTemp=temp-273.15
		}else if toUnit=="F"{
			convertedTemp=(temp-273.15)*9/5+32
		}
	}

fmt.Printf("Coonverted Temperature:%.2f %s\n",convertedTemp,toUnit)
}