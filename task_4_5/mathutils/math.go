package mathutils

import (
	"errors"
	"sort"
)

func Fibonacci(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n cannot be negative")
	}
	if n == 0 {
		return 0, nil
	}
	if n == 1 {
		return 1, nil
	}
	
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b, nil
}

func FibonacciRecursive(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("n cannot be negative")
	}
	if n <= 1 {
		return n, nil
	}
	a, _ := FibonacciRecursive(n - 1)
	b, _ := FibonacciRecursive(n - 2)
	return a + b, nil
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func SortAndDeduplicate(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := make([]int, len(nums))
	copy(result, nums)

	sort.Ints(result)
	j := 0
	for i := 0; i < len(result); i++ {
		if result[j] != result[i] {
			j++
			result[j] = result[i]
		}
	}
	return result[:j+1]
}

func Average(nums []float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("empty slice")
	}
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums)), nil
}
