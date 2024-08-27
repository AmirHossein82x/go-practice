package main

import "fmt"

// "bufio"
// "fmt"
// "os"

// // "reflect"
// // "strconv"
// "strings"
// "unicode"

func main() {
	// res, err := strconv.Atoi("123")
	// if err != nil{
	// 	fmt.Println("error")
	// }else{
	// 	fmt.Println(res, reflect.TypeOf(res))
	// }
	// sample()
	hello()
	fmt.Println("end")
}

func sample() {

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter usernames (separated by space):")
	// input, _ := reader.ReadString('\n')
	// usernames := strings.Fields(input)
	// fmt.Println("Usernames are:", usernames)

}



func hello(){
	defer fmt.Println("yessssssssssssssssssss")
	fmt.Println("hello")
}
