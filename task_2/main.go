package task_2

import (
	"fmt"
	"sync"
)

func printExample() {
	name := "Max"
	age := 35
	fmt.Printf("Name: %s, Age: %d", name, age) // Изменил на корректные
}

func uselessAssignment() {
	x := 5
	y := 1
	z := x + y
	_ = z
}

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increnemt() { // Добавил указатель
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func rangeProblem() {
	var funcs []func()
	values := []int{1, 2, 3, 4, 5}

	for _, v := range values {
		val := v
		funcs = append(funcs, func() {
			fmt.Printf("%d", val) // Создал копию переменной v
		})
	}

	for _, f := range funcs {
		f()
	}
	fmt.Println()
}

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

/*
$ go vet main.go
# command-line-arguments
.\main.go:18:2: self-assignment of x to x
.\main.go:29:9: Increnemt passes lock by value: command-line-arguments.Counter contains sync.Mutex
.\main.go:35:9: GetValue passes lock by value: command-line-arguments.Counter contains sync.Mutex
.\main.go:11:2: fmt.Print call has possible Printf formatting directive %d
.\main.go:61:2: struct field tag `json:email` not compatible with reflect.StructTag.Get: bad syntax for struct tag value
*/
