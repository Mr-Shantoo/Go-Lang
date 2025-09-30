package main

import (
	"fmt"
	"strings"
	
)

func main() {
	s:="Hello World"
	fmt.Println(len(s))
	W_index:=strings.Index(s,"W")
	fmt.Println(W_index)
	substr:=s[:4]
	fmt.Println(substr)
	substr2:=s[2:]
	fmt.Println(substr2)
	string1:="go programming"
	fmt.Println(strings.Replace(string1,"go","went",1))

	fmt.Println(strings.ToLower(string1))
	fmt.Println(strings.ToUpper(string1))

	str:="University"
	fmt.Println(strings.HasPrefix(str,"Uni"))
	fmt.Println(strings.HasSuffix(str,"ty"))

	str2:="A,B,C"
	parts:=strings.Split(str2,",")
	fmt.Println(parts)
	fmt.Println(strings.Join(parts,"-"))
	s1:="ha"
	fmt.Println(strings.Repeat(s1,3))





	
}
