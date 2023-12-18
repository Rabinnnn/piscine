package piscine

// import "fmt"

func Unmatch(a []int) int {
	counts := make(map[int]int)

	// Count occurrences of each element in the slice
	for _, num := range a {
		counts[num]++
	}

	// Find the element with an odd count (no pair)
	for num, count := range counts {
		if count%2 != 0 {
			return num
		}
	}

	// If all elements have a pair, return -1
	return -1
}

/*
func main() {
	a := []int{1, 2, 1, 1, 4, 5, 5, 4, 1, 7}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
} */
