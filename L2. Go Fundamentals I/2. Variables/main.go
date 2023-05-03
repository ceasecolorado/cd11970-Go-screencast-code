package main

import "fmt"

func main() {
	// var product string = "T-shirt"
	// var cost int = 20

	product := "T-shirt"
	cost := 20

	fmt.Println("product's valus is:", product)
	fmt.Printf("product's type is: %T\n", product)

	cost = 15

	fmt.Println("cost's valus is:", cost)
	fmt.Printf("cost's type is: %T\n", cost)
}
