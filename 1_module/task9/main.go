package main

import "fmt"

func main() {
	var balance float64
	fmt.Print("Введите начальный баланс: ")
	fmt.Scan(&balance)

	if balance < 0 {
		fmt.Println("Ошибка: баланс не может быть отрицательным.")
		return
	}

	var choice int
	fmt.Println("\nВыберите действие:")
	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Пополнить баланс")
	fmt.Println("3. Снять средства")
	fmt.Print("Ваш выбор (1-3): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Текущий баланс: %.2f\n", balance)
	case 2:
		var amount float64
		fmt.Print("Введите сумму пополнения: ")
		fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Ошибка: сумма пополнения должна быть больше нуля.")
		} else {
			balance += amount
			fmt.Printf("Успешно пополнено. Новый баланс: %.2f\n", balance)
		}
	case 3:
		var amount float64
		fmt.Print("Введите сумму снятия: ")
		fmt.Scan(&amount)

		if amount <= 0 {
			fmt.Println("Ошибка: сумма снятия должна быть больше нуля.")
		} else if amount > balance {
			fmt.Println("Ошибка: недостаточно средств на счете!")
		} else {
			balance -= amount
			fmt.Printf("Успешно снято. Новый баланс: %.2f\n", balance)
		}
	default:
		fmt.Println("Ошибка: неверный пункт меню.")
	}
}
