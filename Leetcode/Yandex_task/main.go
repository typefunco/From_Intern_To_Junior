package main

import (
	"fmt"
)

// Функция для вычисления НОД
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Функция для вычисления НОК
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	var x, y int
	_, _ = fmt.Scan(&x, &y)

	// Счетчик подходящих пар (p, q)
	count := 0

	// Перебираем все возможные значения p, которые делятся на x и p <= y
	for p := x; p <= y; p += x {
		if y%p == 0 { // Проверка, что y делится на p без остатка
			q := y / p
			if gcd(p, q) != x {
				count++
			}
		}
	}

	fmt.Println(count)
}
