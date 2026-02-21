package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

// Необработанная ошибка
func readFileUnsafe() {
	data, _ := ioutil.ReadFile("config.txt")
	fmt.Println(string(data))
}

// Магическое число
func calculateDiscount(price float64) float64 {
	return price * 0.15
}

//Неиспользуемый параметр (unparam)
func processUser(name string, age int, unused string) string {
	return fmt.Sprintf("User: %s, Age: %d", name, age)
}

//Избыточная сложность цикла
func complexFunction(items []int) int {
	result := 0
	for i := 0; i < len(items); i++ {
		if items[i] > 0 {
			if items[i]%2 == 0 {
				if items[i] > 100 {
					result += items[i] * 2
				} else if items[i] > 50 {
					result += items[i]
				} else {
					result += items[i] / 2
				}
			} else {
				if items[i] > 75 {
					result += items[i] * 3
				} else {
					result += items[i]
				}
			}
		}
	}
	return result
}

//Сравнение с ошибкой напрямую
var ErrNotFound = errors.New("not found")

func findItem(id int) error {
	if id < 0 {
		return ErrNotFound
	}
	return nil
}

func processItem(id int) {
	err := findItem(id)
	if err == ErrNotFound { // Сравнение с конкретным значением ошибки
		fmt.Println("Item not found")
	}
}

// Небезопасный вызов Sleep в цикле
func badLoop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Iteration %d\n", i)
		time.Sleep(1 * time.Second) // Небезопасно в проде
	}
}

const longString = "Это очень длинная строка, которая превышает стандартный лимит в 120 символов и должна быть обнаружена линтером lll, который проверяет длину строк в коде"

func main() {
	readFileUnsafe()
	
	discount := calculateDiscount(100.0)
	fmt.Printf("Discount: %.2f\n", discount)
	
	processUser("Alice", 30, "unused")
	
	items := []int{10, 25, 60, 120, 30}
	result := complexFunction(items)
	fmt.Printf("Complex result: %d\n", result)
	
	processItem(-1)
	
	badLoop()
	
	fmt.Println(longString)
}