package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	bits, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Allocate memory for the gray code bits.
	// For N bits, there will be 2^N values
	valueCount := intpower(2, bits)
	values := make([][]byte, valueCount)

	// Each value has N bits. Each bit in a value
	// kept as a Go byte
	for i := 0; i < valueCount; i++ {
		values[i] = make([]byte, bits)
	}

	// Fill in all values with appropriate bits,
	// one "column" at a time. Each loop fills in 2^(bit number + 1)
	// of the array of final values.
	for bitNo := 0; bitNo < bits; bitNo++ {
		vertical := intpower(2, bitNo+1)

		if bitNo > 0 {
			// Mirror the existing bits, which will only be
			// initialized for the first vertical/2 values
			for i, j := 0, vertical-1; i < j; i, j = i+1, j-1 {
				for k := 0; k < bitNo; k++ {
					values[j][k] = values[i][k]
				}
			}
		}

		// Fill in top half the current column of bits with zeros,
		// fill in the bottom half with 1s.
		for i, j := 0, vertical-1; i < vertical; i, j = i+1, j-1 {
			var value byte = 0
			if i > j {
				value = 1
			}
			values[i][bitNo] = value
		}
	}

	// Output of the 2^N values. Since bits get kept as Go bytes
	// there's a little bit of work to make human readable output.
	for i := 0; i < valueCount; i++ {
		fmt.Printf("%4d  ", i)
		for j := bits - 1; j >= 0; j-- {
			fmt.Printf("%d", values[i][j])
		}
		fmt.Println()
	}
}

// intpower raises base to the exponent via integer
// multiplication. Exists merely for the convenience.
func intpower(base, exponent int) int {
	retval := 1
	for i := 0; i < exponent; i++ {
		retval *= base
	}
	return retval
}
