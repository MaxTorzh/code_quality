package task_2

import(
	"fmt"
	"sync"
)
// Неправильный форматтер в Printf
func printExample() {
	name := "Max"
	age := 35
	fmt.Print("Name: %d, Age: %s", name, age) // %d для string, %s для int
}

// Бессмысленное присваивание
func uselessAssignment() {
	x := 5
	y := 10
	x = x // Бессмысленное присваивание самому себе
	z := x + y
	_ = z // Blank identifier
}

// Копирование мьютекса по значению
type Counter struct {
	mu sync.Mutex
	value int
}

func (c Counter) Increnemt() { // Передача по значению копирует мьютекс
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func rangeProblem() {
	var funcs []func()
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		funcs = append(funcs, func() {
			fmt.Printf("%d", v) // Все функции печатают последнее значение
		})
	}

	for _, f := range funcs {
		f()
	}
	fmt.Println()
}

// Некорректный тег структуры
type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:email`
}

