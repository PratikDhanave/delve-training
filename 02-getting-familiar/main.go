package main

import (
	"fmt"
	"os"
)

func callme(i int, pi *int, s ...string) {
	fmt.Println("got called with the following arguments:")
	fmt.Println(i, s, pi)
}

func main() {
	vv := 0
	myvar := &vv
	callme(1, &vv, os.Args...)
	fmt.Println(myvar)
}
