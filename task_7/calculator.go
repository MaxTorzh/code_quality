package calculator

import (
	"errors"
	"fmt"
)

type Calculator struct {
	history []string
}

func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]string, 0),
	}
}

func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.history = append(c.history, fmt.Sprintf("%f + %f = %f", a, b, result))
	return result
}

func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.history = append(c.history, fmt.Sprintf("%f - %f = %f", a, b, result))
	return result
}

func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.history = append(c.history, fmt.Sprintf("%f * %f = %f", a, b, result))
	return result
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	result := a / b
	c.history = append(c.history, fmt.Sprintf("%f / %f = %f", a, b, result))
	return result, nil
}

func (c *Calculator) GetHistory() []string {
	return c.history
}

func (c *Calculator) ClearHistory() {
	c.history = make([]string, 0)
}