package main

import (
	"errors"
	"fmt"
)

func main() {
	a := Sum(1, 4)
	b := Subtraction(2, 6)
	c, err := Division(5.0, 0.0)
	d := Multiplication(2, 7)

	fmt.Println(a)
	fmt.Println(b)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}

	fmt.Println(d)
}

// Функция сложения
func Sum(values ...int) (sum int) {
	for _, value := range values {
		sum += value
	}

	return
}

// Функция вычитания
func Subtraction(a, b int) int {
	return a - b
}

// Функция деления
func Division(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("деление на 0 запрещено")
	}

	return a / b, nil
}

// Функция умножения
func Multiplication(a, b int) int {
	return a * b
}
