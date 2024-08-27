// package main

// import "fmt"

// func main() {
// 	firstSlice := []int{1, 2, 3}
// 	secondSlice := make([]int, cap(firstSlice))
// 	firstArray := [3]int{1, 2, 3}

// 	copy(secondSlice, firstSlice)
// 	// firstSlice[0] = 100
// 	changeAll(firstSlice)
// 	// fmt.Println(secondSlice)
// 	fmt.Println(firstSlice)


// 	changeAllForArray(&firstArray)

// 	fmt.Println(firstArray)
// }


// func changeItem(slc []int) {
// 	slc[0] = 100000
// }



// func changeAll(slc []int) {
// 	for i, item := range slc{
// 		slc[i] = item * 3
// 	}
// }




// func changeAllForArray(arr *[3]int) {
// 	for i, item := range arr{
// 		arr[i] = item * 3
// 	}
// }