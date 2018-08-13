// A non-empty array A consisting of N integers is given. Array A represents numbers on a tape.

// Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].

// The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|

// In other words, it is the absolute difference between the sum of the first part and the sum of the second part.

// For example, consider array A such that:

//   A[0] = 3
//   A[1] = 1
//   A[2] = 2
//   A[3] = 4
//   A[4] = 3
// We can split this tape in four places:

// P = 1, difference = |3 − 10| = 7
// P = 2, difference = |4 − 9| = 5
// P = 3, difference = |6 − 7| = 1
// P = 4, difference = |10 − 3| = 7
// Write a function:

// func Solution(A []int) int

// that, given a non-empty array A of N integers, returns the minimal difference that can be achieved.

// For example, given:

//   A[0] = 3
//   A[1] = 1
//   A[2] = 2
//   A[3] = 4
//   A[4] = 3
// the function should return 1, as explained above.

// Assume that:

// N is an integer within the range [2..100,000];
// each element of array A is an integer within the range [−1,000..1,000].
// Complexity:

// expected worst-case time complexity is O(N);
// expected worst-case space complexity is O(N) (not counting the storage required for input arguments).

package main

import (
	"fmt"
)

func main() {
	arr := [5]int{6, 1, 1, 5, 3}
	fmt.Println("Array = ", arr)
	fmt.Println("minimal difference = ", Solution(arr))
}

// Solution function
func Solution(A [5]int) int {
	var tape int
	difference := 1
	B := A[:]
	for i := 0; i < len(A); i++ {
		tape = i + 1
		var part1, part2 int
		if tape == len(A) {
			break
		}
		for i := 0; i < tape; i++ {
			part1 += B[i]
		}
		for j := tape; j < len(A); j++ {
			part2 += B[j]
		}
		var cndDiff int
		if part1 > part2 {
			cndDiff = part1 - part2
		} else {
			cndDiff = part2 - part1
		}
		if cndDiff <= difference {
			difference = cndDiff
		}
		fmt.Println("P = ", tape, ", difference : ", part1, "(part1) -(+) ", part2, "(part2) = ", cndDiff)

	}
	return difference
}
