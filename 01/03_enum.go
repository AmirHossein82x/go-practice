// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// const apiUrl = "https://api.github.com/users/"

// type Order struct {
// 	ID     int
// 	price  int
// 	status OrderStatus
// }

// type OrderStatus int

// const (
// 	created           OrderStatus = 0
// 	Processing                    = 1
// 	WaitingForPayment             = 2
// 	PaymentDone                   = 3
// 	Issue                         = 4
// 	Refund                        = 5
// )

// func main() {
// 	first_order := Order{ID: 1, price: 100, status: WaitingForPayment}
// 	second_order := Order{ID: 2, price: 200, status: Issue}
// 	// fmt.Printf("first order", first_order, second_order)
// 	// fmt.Print(first_order)

// 	firstOrderJson , _ := json.Marshal(first_order)
// 	fmt.Printf(string(firstOrderJson))
// 	fmt.Printf("\n%+v", second_order)

// }
