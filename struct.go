package main

import "fmt"

type customer struct {
	fname            string
	lname            string
	Age              int
	Subscriber       bool
	HomeAddress      string
	Phone            int
	CreditAvailable  float32
	CurrentCartCost  float32
	CurrentOrderCost float32
}

func main() {
	customer_1 := customer{
		fname:            "Annakin",
		lname:            "Skywalker",
		Age:              45,
		Subscriber:       true,
		HomeAddress:      "Death Star",
		Phone:            1234567,
		CreditAvailable:  10000.00,
		CurrentCartCost:  0.00,
		CurrentOrderCost: 0.00}

	customer_2 := customer{
		fname:            "Han",
		lname:            "Solo",
		Age:              50,
		Subscriber:       false,
		HomeAddress:      "Tatooine",
		Phone:            4321765,
		CreditAvailable:  20000.00,
		CurrentCartCost:  0.00,
		CurrentOrderCost: 0.00}

	fmt.Print(customer_1, customer_2)
}
