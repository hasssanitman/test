package main

import "fmt"

func main() {
	arr := []int{1, 6, 1, 3, 3}
	maped := make(map[int]int)
	for _, v := range arr {
		if _, ok := maped[v]; ok {
			maped[v] += 1
		} else {
			maped[v] = 1
		}
	}
	for k, v := range maped {
		if v == 1 {
			fmt.Println(k)
		}
	}

	f1([]int{1, 3, 5, 4, 2, 6, 7, 8})
}

// [1,3,5,4,2]
func f1(arr []int) {
	len := len(arr)
	for _, v := range arr {
		if v == len {
			fmt.Println(v + 1)
		}
	}
}
