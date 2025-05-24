package main 
import "fmt"
func pointtracker() func (int)int{
	points:=0
	return func(p int)int{
		points+=p
		return points
	}
}
func fib()func()int{
	a,b:=0,1
	return func()int{
		a,b=b,a+b
		return a
	}
}
func main(){
	addpoints:=pointtracker()
	fmt.Println(addpoints(10))
	fmt.Println(addpoints(20))
	fmt.Println("Fibonacci sequence using closure:")
	count:=5
	Fibonacci:=fib()
for i := 0; i < count; i++ {
	fmt.Println(Fibonacci())
}

	
}