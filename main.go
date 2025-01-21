package main

import "fmt"

const (
	one = iota*2 + 1
	three
	five
	seven
	nine
	eleven
)

// 1 3 5 7 9 11
func main() {
	fmt.Println(one, three, five, seven, nine, eleven)
}
