package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting application...")
	
	result, err := divide(10, 0)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	
	log.Printf("Result: %d", result)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}