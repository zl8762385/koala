package main

import "fmt"

func test1(args ...string) {
	fmt.Printf("%+v", args)
}

func main() {

	s1 := "111"
	s2 := "222"
	test1(s1)
	test1(s2)
}
