package main

import (
	"fmt"
	"math"
)

func main() {
	var temp float64 = 12.30            // Example temperature value in Celsius
	fmt.Println(math.Round(temp))       // Rounds to nearest integer: 12.0
	fmt.Println(math.Ceil(temp))        // Rounds up: 13.0
	fmt.Println(math.Floor(temp))       // Rounds down: 12.0
	fmt.Println(math.Pow(2, 3))         // 2 raised to the power 3: 8.0
	fmt.Println(math.Abs(float64(-10))) // Absolute value: 10.0
	fmt.Println(math.Sqrt(16))          // Square root: 4.0
}
