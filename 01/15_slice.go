// package main

// import "fmt"

// func main() {
// 	firstSlice := []int{1, 2, 3, 5, 6}
// 	// firstSlice = append(firstSlice[:3], 0)
// 	// firstSlice = append(firstSlice[:3], firstSlice[2:]... )
// 	// firstSlice[3] = 0
// 	// insert(&firstSlice, 0, 0)
// 	firstSlice = insert(firstSlice, 5, 0)
// 	// firstSlice = append([]int{0}, firstSlice...)
// 	fmt.Println(firstSlice)
	
// }



// // func insert(slice []int, index int, value int) {
// // 	if index == 0{
// // 		return append([]int{value}, slice...)
		
// // 	} else{
// // 		return 0
// // 	}
// // }




// func insert(a []int, index int, value int) []int {
//     if len(a) == index { // nil or empty slice or after last element
//         return append(a, value)
//     }
//     a = append(a[:index+1], a[index:]...) // index < len(a)
//     a[index] = value
//     return a
// }



// // func insert2(a *[]int, index int, element int) {
// // 	if len(*a) == index{
// // 		*a = append(*a, element)
// // 	}else{
// // 		*a = append(*a[:index+1], *a[index:]...)
// // 	}
// // }


// func newInsert(slc []int, index int, element int) [] int{
// 	if len(slc) == index {
// 		return append(slc, element)
// 	}else {
// 		slc = append(slc[:index + 1], slc[index:]...)
// 		slc[index] = element
// 		return slc
// 	}
// }