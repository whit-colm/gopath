package main

import "fmt"
import "os"

var someString string

func main() {
	if len(os.Args) != 3 || os.Args[1] != "-s" {
		fmt.Println("No.")
	} else {
		someString = os.Args[2]
		fmt.Println(someString)
	}
}
