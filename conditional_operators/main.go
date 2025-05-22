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
}