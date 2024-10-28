package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func main() {
	var i int

	fmt.Print("Введите номер задания: ")
	fmt.Scanln(&i)

	// Цикл выбора задания
	for i > 0 {
		switch i {
		case 1:
			task_1()
		case 2:
			task_2()
		case 3:
			task_3()
		case 4:
			task_4()
		case 5:
			task_5()
		case 6:
			task_6()
		case 7:
			task_7()
		case 8:
			task_8()
		case 9:
			task_9()
		case 10:
			task_10()

		default:
			fmt.Println("Задание не найдено...")
		}

		fmt.Print("Введите номер задания: ")
		fmt.Scanln(&i)
	}

	fmt.Println("...Завершение программы...")
}

func task_1() {
	var base, height float64

	fmt.Println("Введите длину стороны и высоты треугольника: ")
	fmt.Scanln(&base, &height)

	triangleArea(base, height)
}
func triangleArea(base float64, height float64) {
	area := base * height / 2
	fmt.Println("Площадь треугольника = ", area)
}

func task_2() {
	var len int

	fmt.Println("Введите колличество элементов массива: ")
	fmt.Scanln(&len)

	var slice []int = make([]int, len)

	for index, r := range slice {
		fmt.Println("Введите ", index+1, " элемент массива: ")
		fmt.Scanln(&r)
		slice[index] = r
	}
	fmt.Println("Исходным массив: ", slice)

	sortArray(slice)

	fmt.Println("Отсортированный массив по абсолютным значениям:", slice)
}
func sortArray(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Сравниваем абсолютные значения
			if math.Abs(float64(arr[j])) > math.Abs(float64(arr[j+1])) {
				// Меняем местами
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println(arr)
			}
		}
	}
}

func task_3() {
	var number int

	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	sumOfSquares(number)
}
func sumOfSquares(n int) {
	var sum int

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum = sum + (i * i)
		}
	}

	fmt.Println("Сумма квадратов до заданного числа: ", sum)
}

func task_4() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку: ")
	if scanner.Scan() {
		input = scanner.Text()
	}

	isPalindrome(input)
}
func isPalindrome(input string) {
	var runes []string

	for _, rune := range input {
		if unicode.IsPunct(rune) || unicode.IsSpace(rune) {
			continue
		}
		runes = append(runes, string(rune))
	}

	var senur []string = make([]string, len(runes))
	for index, rune := range runes {
		senur[len(senur)-index-1] = rune
	}

	var senur1, runes1 string = "", ""

	for _, rune := range runes {
		runes1 += rune
	}

	for _, rune := range senur {
		senur1 = senur1 + rune
	}

	runes1 = strings.ToLower(runes1)
	senur1 = strings.ToLower(senur1)

	if runes1 == senur1 {
		fmt.Println("Эта строка является палиндромом")
	} else {
		fmt.Println("Эта строка не является палиднромом")
	}
}

func task_5() {
	var number int

	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	isPrime(number)
}
func isPrime(number int) {
	var count, del int

	for i := 2; i <= number; i++ {
		if number%i == 0 {
			count += 1
			if del == 0 {
				del = i
			}
		}
	}

	if count == 1 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func task_6() {
	var number2 int
	var arr []int

	fmt.Print("Введите лимит: ")
	fmt.Scanln(&number2)

	arr = generatePrimes(number2)
	fmt.Println("Все простые числа до лимита: ", arr)
}
func generatePrimes(number2 int) []int {
	var number1, k int
	var arr []int

	for i := number1; i <= number2; i++ {
		for n := 1; n <= number2; n++ {
			if i%n == 0 {
				k++
			}
		}
		if k == 2 {
			arr = append(arr, i)
			k = 0
		}
		k = 0
	}
	return arr
}

func task_7() {
	var number int

	fmt.Println("Введите число: ")
	fmt.Scanln(&number)
	fmt.Println("Число в двоичной системе счистления: ", toBinary(number))

}
func toBinary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		binary = string(n%2+'0') + binary
		n = n / 2
	}
	return binary
}

func task_8() {
	var len, element int

	fmt.Println("Введите размер массива: ")
	fmt.Scanln(&len)

	var arr0 []int

	for i := 1; i <= len; i++ {
		fmt.Println("Введите ", i, " элемент массива: ")
		fmt.Scanln(&element)
		arr0 = append(arr0, element)
	}

	findMax(arr0)
}
func findMax(arr []int) {
	var max int

	for _, r := range arr {
		if max < r {
			max = r
		}
	}

	fmt.Println("Максимальный элемент массива: ", max)
}

func task_9() {
	var a, b, greater, less int
	fmt.Print("Введите два числа для нахождения НОД: ")
	fmt.Scanln(&a, &b)

	if a > b {
		greater = a
		less = b
	} else {
		greater = b
		less = a
	}

	gcd(greater, less)
}
func gcd(greater int, less int) {
	var a int
	for i := 0; i < 1; {
		if greater%less == 0 {
			fmt.Println("Наибольший общий делитель = ", less)
			i = 1
		} else {
			a = greater
			greater = less
			less = a % less
		}
	}
}

func task_10() {
	var len, element int

	fmt.Println("Введите размер массива: ")
	fmt.Scanln(&len)

	var arr0 []int

	for i := 1; i <= len; i++ {
		fmt.Println("Введите ", i, " элемент массива: ")
		fmt.Scanln(&element)
		arr0 = append(arr0, element)
	}
	element = sumArray(arr0)
	fmt.Println("Сумма элементов массива: ", element)

}
func sumArray(arr []int) int {
	var i int
	for _, r := range arr {
		i += r
	}
	return i
}
