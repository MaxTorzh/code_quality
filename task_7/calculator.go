package calculator

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
)

type Operation struct {
	Type  string
	Value float64
}

type Calculator struct {
	history   []string
	lastValue float64
	ops       []Operation
}

func NewCalculator() *Calculator {
	return &Calculator{
		history:   make([]string, 0),
		lastValue: 0,
		ops:       make([]Operation, 0),
	}
}

func (c *Calculator) Add(a, b float64) (float64, error) {
	if (b > 0 && a > math.MaxFloat64-b) || (b < 0 && a < -math.MaxFloat64-b) {
		return 0, errors.New("addition overflow")
	}
	
	result := a + b
	result = c.round(result)
	c.lastValue = result
	c.addToHistory(fmt.Sprintf("%f + %f = %f", a, b, result))
	c.ops = append(c.ops, Operation{Type: "add", Value: b})
	return result, nil
}

func (c *Calculator) Subtract(a, b float64) (float64, error) {
	return c.Add(a, -b)
}

func (c *Calculator) Multiply(a, b float64) (float64, error) {
	if math.Abs(a) > math.MaxFloat64/math.Abs(b) {
		return 0, errors.New("multiplication overflow")
	}
	
	result := a * b
	result = c.round(result)
	c.lastValue = result
	c.addToHistory(fmt.Sprintf("%f * %f = %f", a, b, result))
	c.ops = append(c.ops, Operation{Type: "multiply", Value: b})
	return result, nil
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	
	result := a / b
	result = c.round(result)
	c.lastValue = result
	c.addToHistory(fmt.Sprintf("%f / %f = %f", a, b, result))
	c.ops = append(c.ops, Operation{Type: "divide", Value: b})
	return result, nil
}

func (c *Calculator) Start(value float64) *Calculator {
	c.lastValue = value
	c.ops = make([]Operation, 0)
	return c
}

func (c *Calculator) AddValue(val float64) (*Calculator, error) {
	result, err := c.Add(c.lastValue, val)
	if err != nil {
		return nil, err
	}
	c.lastValue = result
	return c, nil
}

func (c *Calculator) SubtractValue(val float64) (*Calculator, error) {
	return c.AddValue(-val)
}

func (c *Calculator) MultiplyValue(val float64) (*Calculator, error) {
	result, err := c.Multiply(c.lastValue, val)
	if err != nil {
		return nil, err
	}
	c.lastValue = result
	return c, nil
}

func (c *Calculator) DivideValue(val float64) (*Calculator, error) {
	result, err := c.Divide(c.lastValue, val)
	if err != nil {
		return nil, err
	}
	c.lastValue = result
	return c, nil
}

func (c *Calculator) Result() float64 {
	return c.lastValue
}

func (c *Calculator) SaveHistory(filename string) error {
	data, err := json.Marshal(c.history)
	if err != nil {
		return fmt.Errorf("failed to marshal history: %w", err)
	}
	
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save history: %w", err)
	}
	
	return nil
}

func (c *Calculator) LoadHistory(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read history: %w", err)
	}
	
	err = json.Unmarshal(data, &c.history)
	if err != nil {
		return fmt.Errorf("failed to unmarshal history: %w", err)
	}
	
	return nil
}

func (c *Calculator) Percentage(value, percent float64) float64 {
	result := value * percent / 100
	result = c.round(result)
	c.addToHistory(fmt.Sprintf("%f%% of %f = %f", percent, value, result))
	return result
}

func (c *Calculator) Pow(base, exponent float64) (float64, error) {
	if base < 0 && math.Floor(exponent) != exponent {
		return 0, errors.New("negative base with fractional exponent")
	}
	
	result := math.Pow(base, exponent)
	if math.IsInf(result, 0) || math.IsNaN(result) {
		return 0, errors.New("power operation resulted in overflow or invalid value")
	}
	
	result = c.round(result)
	c.addToHistory(fmt.Sprintf("%f ^ %f = %f", base, exponent, result))
	return result, nil
}

func (c *Calculator) round(value float64) float64 {
	return math.Round(value*1e10) / 1e10
}

func (c *Calculator) addToHistory(entry string) {
	c.history = append(c.history, entry)
}

func (c *Calculator) GetHistory() []string {
	return c.history
}

func (c *Calculator) ClearHistory() {
	c.history = make([]string, 0)
}