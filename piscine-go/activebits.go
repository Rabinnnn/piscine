package piscine

// ActiveBits returns the number of active bits (bits with the value 1) in the binary representation of an integer number.
func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		// Increment count if the least significant bit is 1
		count += n & 1
		// Right shift the number to check the next bit
		n >>= 1
	}
	return count
}
