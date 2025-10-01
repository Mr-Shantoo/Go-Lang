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
fmt.Printf("name:%s, height:%s\n",name ,height)
//iota= sets value of constants from 0 if we want to start from 1 thn before 
// the constants we can use _=iota 
const (
        _ = iota
        fullname
        age
        class
    )

    fmt.Println("fullname:", fullname)
    fmt.Println("age:", age)
    fmt.Println("class:", class)
  }