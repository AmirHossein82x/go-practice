// package main

// import "fmt"

// func main() {
// 	// i, j := 2, 3
// 	i := 2
// 	// fmt.Println(&i)

// 	// // fmt.Print(i, j)
// 	// // fmt.Println(i, j)
// 	// var ip *int = &i
// 	// *ip += 5

// 	// fmt.Println(ip)
// 	// fmt.Println(*ip)
// 	// fmt.Printf("i: %d", i)
// 	fmt.Println("outside num: %d", &i)

// 	// changeValueByRef(&i)
// 	changeValueByValue(i)
// 	fmt.Println(i)
// }



// func changeValueByRef(num *int) {
// 	fmt.Println("inside num: %d", num)
// 	*num += 4
// }

// func changeValueByValue(num int) {
// 	fmt.Println("inside num: %d", &num)
// 	num += 4
// }