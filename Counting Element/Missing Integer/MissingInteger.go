// This is a demo task.

// Write a function:

// func Solution(A []int) int

// that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

// For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

// Given A = [1, 2, 3], the function should return 4.

// Given A = [−1, −3], the function should return 1.

// Assume that:

// N is an integer within the range [1..100,000];
// each element of array A is an integer within the range [−1,000,000..1,000,000].
// Complexity:

// expected worst-case time complexity is O(N);
// expected worst-case space complexity is O(N) (not counting the storage required for input arguments).

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 6, 4, 1, 2}
	fmt.Println("Given Array = ", arr)
	fmt.Println("Smallest positive int out array = ", Solution(arr))
}

// Solution function
func Solution(A []int) int {
	sort.Ints(A)
	fmt.Println("Sorted array = ", A)
	smallInt := 1
	var cndSmallInt int
	for i := 0; i < len(A); i++ {
		if i < len(A)-1 {
			j := A[i+1] - A[i]
			if j >= 2 {
				cndSmallInt = A[i] + 1
			}
		}
		// if (cndSmallInt == 0) && (i == len(A)) {
		// 	for {
		// 		cndSmallInt++
		// 		if cndSmallInt > 0 {
		// 			break
		// 		}
		// 	}
		// }
		if cndSmallInt < smallInt {
			smallInt = cndSmallInt
		}
		fmt.Println("Index = ", i, ", cnd = ", cndSmallInt, ", smallest = ", smallInt)
	}
	return smallInt
}
