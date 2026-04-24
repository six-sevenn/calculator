package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("╔════════════════════════════════╗")
	fmt.Println("║       ПРОСТОЙ КАЛЬКУЛЯТОР       ║")
	fmt.Println("╚════════════════════════════════╝")
	fmt.Println("\nДоступные операции: +, -, *, /")
	fmt.Println("Формат ввода: число пробел оператор пробел число")
	fmt.Println("Пример: 5 + 3")
	fmt.Println("Введите 'exit' для выхода\n")
	
	for {
		fmt.Print("Введите выражение: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "exit" {
			fmt.Println("\nДо свидания!")
			break
		}
		
		if input == "" {
			continue
		}
		
		result, err := calculate(input)
		if err != nil {
			fmt.Printf("❌ Ошибка: %v\n\n", err)
		} else {
			fmt.Printf("✅ Результат: %g\n\n", result)
		}
	}
}

func calculate(expression string) (float64, error) {
	// Разделяем выражение на части
	parts := strings.Fields(expression)
	
	if len(parts) != 3 {
		return 0, fmt.Errorf("неверный формат. Используйте: число оператор число")
	}
	
	// Парсим первое число
	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' - не число", parts[0])
	}
	
	// Получаем оператор
	operator := parts[1]
	
	// Парсим второе число
	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' - не число", parts[2])
	}
	
	// Выполняем операцию
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль невозможно")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("оператор '%s' не поддерживается. Используйте +, -, *, /", operator)
	}
}
