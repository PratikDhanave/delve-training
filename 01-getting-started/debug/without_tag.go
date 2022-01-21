//go:build !user_tag
// +build !user_tag

package main

func callme() {
	println("no user_tag provided")
}
