package main

import "fmt"

/////////////////////////
// 1. Basic Function
/////////////////////////

// calculateDiscount computes the final price after applying a discount.
func calculateDiscount(price float64, discount float64) float64 {
	return price - (price * discount / 100)
}

/////////////////////////
// 2. Multiple Return Values
/////////////////////////

// login authenticates a user and returns a success status and message.
func login(username string, password string) (bool, string) {
	if username == "admin" && password == "123" {
		return true, "Login successful"
	}
	return false, "Invalid credentials"
}

/////////////////////////
// 3. Named Return Values
/////////////////////////

// transactionFee calculates a 1.5% transaction fee for a given amount.
func transactionFee(amount float64) (fee float64) {
	fee = amount * 0.015
	return
}

/////////////////////////
// 4. Variadic Function
/////////////////////////

// total sums up any number of float64 values.
func total(items ...float64) (sum float64) {
	for _, item := range items {
		sum += item
	}
	return
}

/////////////////////////
// 5. First-Class Functions
/////////////////////////

// processNumber applies an operation (function) to a given number.
func processNumber(num int, operation func(int) int) int {
	return operation(num)
}

// double returns the double of a number.
func double(num int) int {
	return num * 2
}

// half returns half of a number.
func half(num int) int {
	return num / 2
}

/////////////////////////
// Main Function
/////////////////////////

func main() {
	fmt.Println("=== Function Examples in Go ===\n")

	// 1. Basic Function
	fmt.Println("Discounted price is:", calculateDiscount(100, 10))

	// 2. Multiple Return Values
	success, message := login("admin", "123")
	fmt.Println("Login result:", success, "-", message)

	// 3. Named Return Values
	fmt.Println("Transaction fee on $100:", transactionFee(100))

	// 4. Variadic Function
	fmt.Println("Total bill from individual items:", total(5.5, 10.10, 20.5))

	// 4. Using Variadic Function with a Slice
	fmt.Println("Using variadic function with a slice:")
	items := []float64{12.5, 10.5, 10.10, 20.5}
	fmt.Println("Total bill from slice:", total(items...))

	// 5. First-Class Functions
	fmt.Println("First-class function examples:")
	fmt.Println("Double of 10:", processNumber(10, double))
	fmt.Println("Half of 10:", processNumber(10, half))
}
