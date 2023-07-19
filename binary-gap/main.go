package main

// you can also use imports, for example:
import (
	"fmt"
)

func main() {
	fmt.Printf("val=%d", findLongestBinaryGap(1041))
}

func findLongestBinaryGap(N int) int {
	var result int      // Variable to store the length of the longest binary gap found.
	var result_temp int // Variable to store the length of the current potential binary gap.
	var calc bool       // Boolean variable to track whether we are currently calculating a binary gap.

	for N > 0 {
		// Check the least significant bit (LSB) of N (1 if odd, 0 if even).
		if N%2 == 1 {
			if !calc { // If this is the first '1' encountered, it indicates the start of a binary gap.
				calc = true // Set calc to true to start counting the binary gap.
			} else { // If this is not the first '1', it indicates the end of a binary gap.
				if result_temp > result {
					result = result_temp // Update result with the length of the current binary gap if it's larger.
				}
				result_temp = 0 // Reset result_temp to 0 to start counting the next potential binary gap.
			}
		} else {
			if calc { // If we are currently counting a binary gap (inside a sequence of 1s).
				result_temp++ // Increment result_temp to keep track of the length of the current binary gap.
			}
		}
		// Move to the next bit by right shifting N (equivalent to dividing N by 2).
		N = N / 2
	}

	return result // Return the length of the longest binary gap found in N's binary representation.
}
