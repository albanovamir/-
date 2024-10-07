package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	fmt.Print("Введите номер задания: ")
	fmt.Scan(&i)

	// Цикл выбора задания
	for i > 0 {
		switch i {
		case 1_1:
			task_1_1()
		case 1_2:
			task_1_2()
		case 1_3:
			task_1_3()
		case 1_4:
			task_1_4()
		case 1_5:
			task_1_5()
		case 2_1:
			task_2_1()
		case 2_2:
			task_2_2()
		case 2_3:
			task_2_3()
		case 2_4:
			task_2_4()
		case 2_5:
			task_2_5()
		case 3_1:
			task_3_1()
		case 3_2:
			task_3_2()
		case 3_3:
			task_3_3()
		case 3_4:
			task_3_4()
		case 3_5:
			task_3_5()

		default:
			fmt.Println("Задание не найдено...")
		}

		fmt.Print("Введите номер задания: ")
		fmt.Scan(&i)
	}

	fmt.Println("...Завершение программы...")
}

/////////////////////
// Функции заданий //
/////////////////////

func task_1_1() {
	var n int
	fmt.Print("Введите четырехзначное число: ")
	fmt.Scan(&n)
	fmt.Print("Сумма цифр числа = ")
	fmt.Println((n / 1000 % 10) + (n / 100 % 10) + (n / 10 % 10) + (n % 10))
}

func task_1_2() {
	var choice int
	var temperature float64

	fmt.Println("Выберите вариант преобразования:")
	fmt.Println("1. Цельсий в Фаренгейт")
	fmt.Println("2. Фаренгейт в Цельсий")
	fmt.Print("Введите ваш выбор (1 или 2): ")
	fmt.Scan(&choice)

	// Массив функций преобразования
	conversions := [2]func(float64) float64{
		celsiusToFahrenheit,
		fahrenheitToCelsius,
	}

	// Массив строк формата результата
	results := [2]string{"°C = %.2f °F\n", "°F = %.2f °C\n"}

	// Вызываем соответствующую функцию.
	fmt.Print("Введите температуру: ")
	fmt.Scan(&temperature)
	result := conversions[choice-1](temperature)
	fmt.Printf(results[choice-1], result)
}

func task_1_3() {
	var a, b, c, d int

	fmt.Print("Введите первый элемент массива: ")
	fmt.Scan(&a)
	fmt.Print("Введите второй элемент массива: ")
	fmt.Scan(&b)
	fmt.Print("Введите третий элемент массива: ")
	fmt.Scan(&c)
	fmt.Print("Введите четвертый элемент массива: ")
	fmt.Scan(&d)

	list := [4]int{a * 2, b * 2, c * 2, d * 2}

	fmt.Println("Массив с удвоенными значениями: ")
	fmt.Print(list)
}

func task_1_4() {
	var name, second_name, third_name, full_name string
	fmt.Print("Введите вашу фамилию: ")
	fmt.Scan(&second_name)
	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите ваше отчество: ")
	fmt.Scan(&third_name)
	full_name = second_name + " " + name + " " + third_name
	fmt.Println("Ваше полное имя: ", full_name)

}

func task_1_5() {
	var x1, x2, y1, y2 float64
	var distance float64

	fmt.Print("Введите значение координаты x1: ")
	fmt.Scan(&x1)
	fmt.Print("Введите значение координаты y1: ")
	fmt.Scan(&y1)
	fmt.Print("Введите значение координаты x2: ")
	fmt.Scan(&x2)
	fmt.Print("Введите значение координаты y2: ")
	fmt.Scan(&y2)

	distance = math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
	fmt.Println("Расстояние между точкачи = ", distance)
}

func task_2_1() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number%2 == 0 {
		fmt.Println("Число четное")
	} else {
		fmt.Println("Число нечетное")
	}
}

func task_2_2() {
	var year int

	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if year%100 == 0 && year%400 == 0 {
		fmt.Println("Этот год високосный")
	} else if year%4 == 0 && year%100 != 0 {
		fmt.Println("Этот год високосный")
	} else {
		fmt.Println("Этот год не является високосным")
	}
}

func task_2_3() {
	var a, b, c int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&c)

	if a > b && a > c {
		fmt.Println("Число ", a, "наибольшее")
	}
	if b > a && b > c {
		fmt.Println("Число ", b, "наибольшее")
	}
	if c > b && c > a {
		fmt.Println("Число ", c, "наибольшее")
	}
}

// До 14 лет человек считается ребенком, с 14 по 18 лет - подростком, с 18 до 50 лет - взрослым, выше 50 - пожилым
func task_2_4() {
	var age int
	fmt.Print("Введите возраст человека: ")
	fmt.Scan(&age)
	if age < 14 {
		fmt.Sprintln("Это ребенок")
	}
	if age >= 14 && age < 18 {
		fmt.Println("Это подросток")
	}
	if age >= 18 && age <= 50 {
		fmt.Println("Это взрослый")
	}
	if age > 50 {
		fmt.Println("Это пожилой человек")
	}
}

func task_2_5() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Это число делится на 3 и 5 без остатка")
	} else {
		fmt.Println("Это число не делится одновременно на 3 и 5")

	}
}

func task_3_1() {
	var number int

	fmt.Print("Введите число для вычисления его факториала: ")
	fmt.Scan(&number)

	if number < 0 {
		fmt.Println("Факториал отрицательного числа не определен.")
		return
	}

	factorial := 1

	for i := 1; i <= number; i++ {
		factorial *= i
	}

	fmt.Printf("Факториал %d равен %d\n", number, factorial)
}

func task_3_2() {
	var a, b, c, number int
	a = 0
	b = 1
	c = 0

	fmt.Print("Введите колличество колличество чисел Фибоначи: ")
	fmt.Scan(&number)

	if number == 1 {
		fmt.Println(a)
	}

	if number == 2 {
		fmt.Println(a)
		fmt.Println(b)
	}

	if number > 2 {
		fmt.Println(a)
		fmt.Println(b)

		for i := number - 2; i > 0; i-- {
			c = b
			b = a + b
			a = c
			fmt.Println(b)
		}
	}
}

func task_3_3() {

	var (
		array  [5]int
		array2 [5]int = [5]int{0, 0, 0, 0, 0}
		n      int    = 4
	)

	for i := 0; i < 5; i++ {
		fmt.Print("Введите ", i+1, " элемента массива: ")
		fmt.Scan(&array[i])
	}

	for i := 0; i < 5; i++ {
		array2[n] = array[i]
		n--
	}
	fmt.Println("Инвертированный массив: ")
	fmt.Println(array2)
}

func task_3_4() {
	var number int
	var k int = 0

	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Println("Все простые числа до ", number)

	for i := 1; i <= number; i++ {
		for n := 1; n <= number; n++ {
			if i%n == 0 {
				k++
			}
		}
		if k == 2 {
			fmt.Println(i)
			k = 0
		}
		k = 0
	}
}

func task_3_5() {
	var sum int
	var array [5]int

	for i := 0; i < 5; i++ {
		fmt.Print("Введите ", i+1, " элемент массива: ")
		fmt.Scan(&array[i])
	}

	for i := 0; i < 5; i++ {
		sum += array[i]
	}

	fmt.Println("Сумма всех элементов массива = ", sum)
}

////////////////////////////
// Дополнительные функции //
////////////////////////////

// Функции перевода температуры
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
