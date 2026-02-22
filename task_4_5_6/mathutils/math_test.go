package mathutils

import (
	"testing"
)

//Testing

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    int
		wantErr bool
	}{
		{"zero", 0, 0, false},
		{"one", 1, 1, false},
		{"two", 2, 1, false},
		{"three", 3, 2, false},
		{"ten", 10, 55, false},
		{"negative", -1, 0, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fibonacci(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fibonacci() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name string
		n int
		want bool
	}{
		{"zero", 0, false},
		{"one", 1, false},
		{"two", 2, true},
		{"three", 3, true},
		{"four", 4, false},
		{"seventeen", 17, true},
		{"twenty", 20, false},
		{"ninety seven", 97, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortAndDeduplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{5}, []int{5}},
		{"no duplicates", []int{3, 1, 4, 2}, []int{1, 2, 3, 4}},
		{"with duplicates", []int{3, 1, 2, 3, 2, 4}, []int{1, 2, 3, 4}},
		{"all same", []int{1, 1, 1, 1}, []int{1}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortAndDeduplicate(tt.nums)
			
			if len(got) != len(tt.want) {
				t.Errorf("SortAndDeduplicate() length = %v, want %v", len(got), len(tt.want))
			}
			
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("SortAndDeduplicate()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
			
			original := make([]int, len(tt.nums))
			copy(original, tt.nums)
			for i := range tt.nums {
				if tt.nums[i] != original[i] {
					t.Errorf("Original slice was modified at index %d", i)
				}
			}
		})
	}
}

func TestAverage(t *testing.T) {
	tests := []struct {
		name    string
		nums    []float64
		want    float64
		wantErr bool
	}{
		{"empty", []float64{}, 0, true},
		{"single", []float64{5}, 5, false},
		{"two numbers", []float64{2, 4}, 3, false},
		{"three numbers", []float64{1, 2, 3}, 2, false},
		{"with decimals", []float64{1.5, 2.5, 3.5}, 2.5, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Average(tt.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Average() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Benchmarking

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Fibonacci(20)
		if err != nil {
			b.Fatalf("Fibonacci failed: %v", err)
		}
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := FibonacciRecursive(20)
		if err != nil {
			b.Fatalf("FibonacciRecursive failed: %v", err)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(97)
	}
}

func BenchmarkIsPrimeComposite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(100)
	}
}

func BenchmarkSortAndDeduplicate(b *testing.B) {
	data := []int{5, 2, 8, 2, 1, 9, 5, 3, 7, 8, 4, 6}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		SortAndDeduplicate(data)
	}
}

func BenchmarkAverage(b *testing.B) {
	data := []float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5, 9.5, 10.5}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, err := Average(data)
		if err != nil {
			b.Fatalf("Average failed: %v", err)
		}
	}
}