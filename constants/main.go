package main
 import (
	"fmt"
 )
 func main(){
const pie float64 =3.14
fmt.Println(pie)
//multiple constants
const(
	name="maine uddin shanto"
	height="5'8"
)
fmt.Printf("name:%s, height:%s",name ,height)
//iota
const(
	
	fullname=iota
	age
	class
)
fmt.Println(fullname)
fmt.Println(age)
fmt.Println(class)
  }