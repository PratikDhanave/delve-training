package main

import "os"

func f1() {
	println("f1")
	f2()
}

func f2() {
	println("f2")
	f3()
}

func f3() {
	println("f3")
}

func main() {
	f1()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println(dir)
	callme()
}
