package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число N: ")
	fmt.Scan(&n)

	if n > 0 {
		fmt.Println("Число положительное")
	} else if n < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Число равно нулю")
	}

	if n%2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число нечётное")
	}

	if n%3 == 0 {
		fmt.Println("Делится на 3")
	} else {
		fmt.Println("Не делится на 3")
	}

	if n%5 == 0 {
		fmt.Println("Делится на 5")
	} else {
		fmt.Println("Не делится на 5")
	}
}
