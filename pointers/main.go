package main 
import (
	"fmt"
)
func sayHelloPointer(s*string){
	*s="hello world"
}

func sayHello(s string){
	s="hello world"
}

func main(){
	var greeting string= "hello, go"
	sayHello(greeting)
	fmt.Println("after calling say hello:",greeting)//doesn't update the greeting
	sayHelloPointer(&greeting)
	fmt.Println("after calling sayHelloPointer:",greeting)//updates the greeting

}