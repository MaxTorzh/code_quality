package main

import (
	"errors"
	"fmt"
	"os"
)

// Ошибка обработана
func readFileSafe() {
	data, err := os.ReadFile("config.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Config file not found, using defaults")
			return
		}
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println(string(data))
}

// Константа вместо магического числа
const discountRate = 0.15

func calculateDiscount(price float64) float64 {
	return price * discountRate
}

// Убрал неиспользуемый параметр
func processUser(name string, age int) string {
	return fmt.Sprintf("User: %s, Age: %d", name, age)
}

// Упростил сложность
func simplifiedFunction(items []int) int {
	result := 0
	for _, item := range items {
		if item <= 0 {
			continue
		}
		
		switch {
		case item > 100 && item%2 == 0:
			result += item * 2
		case item > 50:
			result += item
		case item > 75 && item%2 != 0:
			result += item * 3
		default:
			if item%2 == 0 {
				result += item / 2
			} else {
				result += item
			}
		}
	}
	return result
}

// Использую errors.Is
var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	if id < 0 {
		return ErrNotFound
	}
	return nil
}

func processItem(id int) {
	err := findItem(id)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Item not found")
	} else if err != nil {
		fmt.Printf("Unexpected error: %v\n", err)
	}
}

// Удалил неиспользуемую глобальную переменную

// Убрал Sleep из цикла или добавим контекст
func betterLoop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration %d\n", i)
	}
}

// Разбил длинную строку
const longString = "Это очень длинная строка, которая превышает стандартный лимит " +
	"в 120 символов. Мы разбили её на несколько частей, чтобы " +
	"удовлетворить требованиям линтера lll."

func main() {
	readFileSafe()
	
	discount := calculateDiscount(100.0)
	fmt.Printf("Discount: %.2f\n", discount)
	
	processUser("Alice", 30)
	
	items := []int{10, 25, 60, 120, 30}
	result := simplifiedFunction(items)
	fmt.Printf("Simplified result: %d\n", result)
	
	processItem(-1)
	
	betterLoop()
	
	fmt.Println(longString)
}