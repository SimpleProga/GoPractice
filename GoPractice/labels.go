package GoPractice

import "fmt"

func Label1() {
outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerLoopLabel
		}
	}
	fmt.Println("End")
}

func Label2() {
outerLoopLabel:
	for i := 0; i < 5; i++ {
		for k := 0; k < 5; k++ {
			fmt.Printf("[%d, %d]\n", i, k)
			break outerLoopLabel
		}
	}
	fmt.Println("End")
}
