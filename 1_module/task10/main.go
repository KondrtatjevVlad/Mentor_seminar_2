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

	sum := x
	max, min := x, x
	evenCount, oddCount := 0, 0

	if x%2 == 0 {
		evenCount++
	} else {
		oddCount++
	}

	for i := 2; i <= n; i++ {
		fmt.Printf("Введите число %d: ", i)
		fmt.Scan(&x)

		sum += x
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
		if x%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	avg := float64(sum) / float64(n)
	fmt.Println("\n--- Статистика последовательности ---")
	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее значение: %.2f\n", avg)
	fmt.Printf("Максимальное значение: %d\n", max)
	fmt.Printf("Минимальное значение: %d\n", min)
	fmt.Printf("Количество чётных: %d\n", evenCount)
	fmt.Printf("Количество нечётных: %d\n", oddCount)
}
