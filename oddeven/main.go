package main

import "fmt"

func main() {

	odd_even_slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range odd_even_slice {
		if val%2 == 0 {
			fmt.Println(val, "is even")
		} else {
			fmt.Println(val, "is odd")
		}
	}
}
