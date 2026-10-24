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

	posCount, negCount, zeroCount := 0, 0, 0

	for i := 1; i <= n; i++ {
		var x int
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&x)

		if x > 0 {
			posCount++
		} else if x < 0 {
			negCount++
		} else {
			zeroCount++
		}
	}

	fmt.Printf("Положительных чисел: %d\n", posCount)
	fmt.Printf("Отрицательных чисел: %d\n", negCount)
	fmt.Printf("Нулей: %d\n", zeroCount)
}
