// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var mySlice []int = []int{1, 0, 3, 4, 5}
// 	target := 4
// 	fmt.Println(find2Sum(mySlice, target))

	
// }


// func find2Sum(mySlice []int, target int) []int{
// 	myMap := map[int]int{}
// 	for idx, number := range mySlice{
// 		diff := target - number
// 		if index, ok := myMap[diff]; ok {
// 			return []int{idx, index}

// 		}else{
// 			myMap[number] = idx
// 		}
// 	}
// 	return []int{0, 0}
// }