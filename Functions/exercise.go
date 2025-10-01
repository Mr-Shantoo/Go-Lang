package main

import (
	"fmt"
	"math"
)
func isPrime(n int)bool{
	if n<=1{
		return false
	}
	for i:=2;i<n;i++{
		if n%i==0{
			return false
		}
	}
	return true
}

func safeDivide(a,b float64)(float64, string){
	if b==0{
		return 0, "Error: division by zero"
	}
	return a/b, "success"
}

func simpleInterest(principal,rate,time float64)(interest float64){
	interest=(principal*rate*time)/100
	return
}

func max(nums ...int)(largest int){
	if len(nums)==0{
		return 0
	}
	largest=nums[0]
	for _,v:= range(nums){
		if v>largest{
			largest=v
		}

	}
	return
}

func applyoperation(n int, op func(int)int)int{
	return op(n)
}
func square(n int)int{
	return n*n
}

func cube(n int)int{
	return n*n*n
}



func main(){
	fmt.Println("=== Function Examples in Go (New Exercise) ===\n")

	// 1. Basic Function
	fmt.Println("Is 7 prime?", isPrime(7))
	fmt.Println("Is 10 prime?", isPrime(10))

	// 2. Multiple Return Values
	result, msg := safeDivide(10, 2)
	fmt.Println("Safe Divide (10/2):", result, "-", msg)
	result, msg = safeDivide(5, 0)
	fmt.Println("Safe Divide (5/0):", result, "-", msg)

	
	// 3. Named Return Values
	fmt.Println("Simple Interest (P=1000, R=5%, T=2yrs):", simpleInterest(1000, 5, 2))

	// 4. Variadic Function
	fmt.Println("Maximum of (3, 7, 2, 9, 5):", max(3, 7, 2, 9, 5))
	nums:=[]int{100,200,300,400,600,}
	// 4. Variadic Function
	fmt.Println("Maximum:",max(nums...))
	//5 first class function
	fmt.Println("square of 5:" ,applyoperation(5 ,square))
	fmt.Println("cube of 3:",applyoperation(3,cube))
	power:=applyoperation(5, func(i int)int{
		return int(math.Pow(float64(i),2))
	})
	fmt.Println(power)
	res:= func(n int, f func(int)int)int{
		return f(n)
	}(5,func (n int)int{
		return n*10
	})
	fmt.Println(res)



}