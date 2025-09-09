package main

import (
	"errors"
	"fmt"
)

func main() {
	a := Sum(1, 4)
	b := Subtraction(2, 6)
	c, err1 := Division(5.0, 0.0)
	d := Multiplication(2, 7)
	e, err2 := Degree(5, 0)

	fmt.Println(a)
	fmt.Println(b)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(c)
	}

	fmt.Println(d)

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(e)
	}
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

// Функция возведения в указанную степень
func Degree(n float64, e int) (float64, error) {
	if n == 0.0 {
		return 0.0, errors.New("значение не опредено")
	}

	if e == 0 {
		return 1.0, nil
	}

	result := 1.0

	for i := 1; i <= e; i++ {
		result *= n
	}

	return result, nil
}
