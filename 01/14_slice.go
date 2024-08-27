// package main

// import "fmt"

// func main() {
// 	firstSlice := []int{1, 2, 3}
// 	secondSlice := []int{4, 5, 6}
// 	// p := &firstSlice
// 	// fmt.Println(p)
// 	firstSlice = append([]int{0}, firstSlice...)
// 	fmt.Printf("%p\n", &firstSlice)
// 	firstSlice = append(firstSlice, secondSlice...)
// 	// fmt.Println(p)
// 	fmt.Printf("%p\n", &firstSlice)

// 	firstSlice = append(firstSlice[:2], firstSlice[3:]...)

// 	fmt.Println(firstSlice)
// }
