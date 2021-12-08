package main

import "fmt"

func main() {
	ch := (chan bool)(nil)
	select {
	case v, ok := <-ch:
		fmt.Println(v, ok)
	default:
		fmt.Println("hello")
	}
}
