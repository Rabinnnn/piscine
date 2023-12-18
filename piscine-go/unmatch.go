package main

import "fmt"

func Unmatch(a []int) int {
	count := make([]int, len(a))
	for i := 0; i < len (a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count[i]++
			} 
		}
	}
	for j := 0; j < len(a); j++ {
		if count[j]%2 != 0 {
			return a[j]
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}
