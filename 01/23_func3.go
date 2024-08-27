// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// copyFile("/home/amirhossein/Desktop/go-practice/new.txt", "/home/amirhossein/Desktop/go-practice/sample.txt")
// 	// copyFile("item.txt", "sample.txt")
// 	readFile("sample.txt")

// }

// func copyFile(destination string, source string) (written int) {
// 	file, err := os.Open(source)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	newFile, err := os.Create(destination)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	io.Copy(newFile, file)
// 	defer newFile.Close()
// 	return

// }


// func readFile(src string) {
// 	file, err := os.ReadFile(src)
// 	if err != nil{
// 		return
// 	}

// 	// fmt.Println(string(len(file)))
// 	// for idx, line := range file{
// 	// 	fmt.Println(string(line), idx)
// 	// }
	
// }