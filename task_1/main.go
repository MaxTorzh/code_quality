package task_1

import (
	"fmt"
)

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{10, 20, 30, 40, 50}

	sum1 := sumArray(arr1)
	sum2 := sumArray(arr2)

	fmt.Printf("First array sum is: %d\n", sum1)
	fmt.Printf("Second array sum is: %d\n", sum2)

	total := sum1 + sum2

	fmt.Printf("Total arrays sum is: %d\n", total)

	maxSum := findMaxSum(arr1, arr2)

	fmt.Printf("Maximum sum is: %d\n", maxSum)
}

func sumArray(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func findMaxSum(arr1, arr2 []int) int {
	maxSum := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; i < len(arr2); i++ {
			currentSum := arr1[i] + arr2[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}
