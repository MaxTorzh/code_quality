package mathutils

import (
	"fmt"
	"os"
	"strconv"
)

func WriteToFile(filename, data string) (err error) {
    f, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    
    defer func() {
        closeErr := f.Close()
        if closeErr != nil && err == nil {
            err = fmt.Errorf("failed to close file: %w", closeErr)
        }
    }()
    
    _, err = f.WriteString(data)
    if err != nil {
        return fmt.Errorf("failed to write to file: %w", err)
    }
    
    return nil
}

func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return string(data), nil
}

func ParseNumbers(strs []string) ([]int, error) {
	result := make([]int, 0, len(strs))
	for i, s := range strs {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number at index %d: %w", i, err)
		}
		result = append(result, num)
	}
	return result, nil
}

func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return fmt.Errorf("failed to delete file %s: %w", filename, err)
	}
	return nil
}