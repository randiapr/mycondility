// A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

// For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

// Write a function:

// func Solution(N int) int

// that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

// For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

// Assume that:

// N is an integer within the range [1..2,147,483,647].
// Complexity:

// expected worst-case time complexity is O(log(N));
// expected worst-case space complexity is O(1).

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 561892 => 10001001001011100100
	// 74901729 => 100011101101110100011100001
	// 1376796946 => 1010010000100000100000100010010
	fmt.Println("Longest binary gap is\t", strconv.Itoa(Solution(561892)))
}

// Solution function to solve binary gap
func Solution(N int) int {
    n64 := int64(N)
    strN64 := strconv.FormatInt(n64,2)
    var(
        gap int
        cndGap int
    )
    sliceBin := strings.Split(strN64, "")
    for _, val := range sliceBin {
        if val == "0" {
            cndGap++
        } else {
            if cndGap > gap {
                gap = cndGap
            }
            cndGap = 0
        }
    }
    return gap
}
