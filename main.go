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
}
