package GoPractice

import "fmt"

func Sum() {
	sum, limit := 0, 100
	for i := 0; true; i++ {
		if i%2 != 0 {
			continue // переход к следующему числу, так как i — нечётное
		}

		if sum+i > limit {
			break // выход из цикла, так как сумма превысит заданный предел
		}

		sum += i
	}
	fmt.Println(sum)
}
