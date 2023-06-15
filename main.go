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

	f1([]int{1, 3, 5, 4, 2, 6, 7, 8, 9})
	ar1, ar2 := f2([]int{1, 3, 1, 4, 3, 6, 8, 8, 9})
	fmt.Println(ar1, ar2)
}

// [1,3,5,4,2]
func f1(arr []int) {
	len := len(arr)
	for _, v := range arr {
		if v == len {
			fmt.Println(v + 1)
			break
		}
	}
}

func f2(arr []int) ([]int, []int) {
	maped := make(map[int]int)
	for _, v := range arr {
		if _, ok := maped[v]; ok {
			maped[v] += 1
		} else {
			maped[v] = 1
		}
	}
	unique := make([]int, 0)
	noUnique := make([]int, 0)

	for k, v := range maped {
		if v > 1 {
			noUnique = append(noUnique, k)
		} else {
			unique = append(unique, k)
		}
	}
	return unique, noUnique
}
