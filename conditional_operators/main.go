package main
import "fmt"

func main(){
	age:=20
	isCitizen:=true
	if age>=18 && isCitizen{
		fmt.Println("you are eligible to vote")
	}else if age<18 && isCitizen{
		fmt.Println("you are a citizen but you aren't eligible to vote")
		
	}else if !isCitizen{
		fmt.Println("you are not a citizen")
	}
	var isweekend bool=true
	var isholiday bool=false
	isdayoff:=isweekend||isholiday
	if isdayoff{
		fmt.Println("it's a offday")
	}
}