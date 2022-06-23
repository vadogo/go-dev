package main

import (
	"fmt"
)

func main() {
	input := []int{4, 1, 4, -4, 6, 3, 8, 8}
	m := make(map[int]int)
	var result []int
	for _, v := range input {
		if _, ok := m[v]; ok {
			continue
		}
		m[v]++
		result = append(result, v)
	}

	fmt.Println(result)

}
