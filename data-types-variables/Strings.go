package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " Hello World"
	W_index := strings.Index(str, "W")
	fmt.Println(W_index)
	fmt.Println(str[W_index:])
	fmt.Println(str[:W_index])
	str2 := "apple"
	fmt.Println(strings.Replace(str2, "a", "b", 1))
	str3 := "go programming"
	fmt.Println(strings.ToUpper(str3))
	str4 := strings.ToUpper(str3)
	fmt.Println(strings.ToLower(str4))
	fmt.Println(strings.Contains(str3, "go"))
	fmt.Println(strings.Contains(str3, "golang"))
}
