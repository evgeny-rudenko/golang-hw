package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	str := "Hello, OTUS!"
	reversedstr := stringutil.Reverse(str)
	fmt.Print(reversedstr)
}
