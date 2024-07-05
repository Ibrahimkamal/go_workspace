package main

import "fmt"

func main() {
	arr := []int{}

	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
	}

	for _, val := range arr {
		if val%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
