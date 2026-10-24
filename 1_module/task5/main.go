package main

import "fmt"

func main() {
	var n int
	fmt.Print("Сколько чисел введёте? ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Ошибка: количество чисел должно быть больше нуля.")
		return
	}

	var x int
	fmt.Print("Введите число 1: ")
	fmt.Scan(&x)
	max, min := x, x

	for i := 2; i <= n; i++ {
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&x)
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Минимальное значение: %d\n", min)
}
