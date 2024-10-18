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
		fmt.Scanln(&i)
	}

	fmt.Println("...Завершение программы...")
}

// ///////////////////////////////////////////////////////
// ////////////////// Функции заданий ////////////////////
// ///////////////////////////////////////////////////////
func task_1_1() {
	var number string
	var baseFrom, baseTo int

	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	fmt.Print("Введите исходную систему счисления (2-36): ")
	fmt.Scanln(&baseFrom)
	fmt.Print("Введите конечную систему счисления (2-36): ")
	fmt.Scanln(&baseTo)

	if baseFrom < 2 || baseFrom > 36 || baseTo < 2 || baseTo > 36 {
		fmt.Println("Система счисления должна быть в диапазоне от 2 до 36.")
		return
	}

	result, err := convertBase(number, baseFrom, baseTo)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Результат: ", result)
}
func convertBase(number string, baseFrom int, baseTo int) (string, error) {
	decimalValue := 0
	for i, digit := range strings.ToUpper(number) {
		var value int
		if digit >= '0' && digit <= '9' {
			value = int(digit - '0')
		} else if digit >= 'A' && digit <= 'Z' {
			value = int(digit - 'A' + 10)
		} else {
			return "", fmt.Errorf("недопустимый символ: %c", digit)
		}

		if value >= baseFrom {
			return "", fmt.Errorf("цифра %c недопустима в системе счисления %d", digit, baseFrom)
		}

		decimalValue += value * int(math.Pow(float64(baseFrom), float64(len(number)-i-1)))
	}

	if decimalValue == 0 {
		return "0", nil
	}
	result := ""
	for decimalValue > 0 {
		remainder := decimalValue % baseTo
		var digit rune
		if remainder < 10 {
			digit = '0' + rune(remainder)
		} else {
			digit = 'A' + rune(remainder-10)
		}
		result = string(digit) + result
		decimalValue /= baseTo
	}

	return result, nil
}

func task_1_2() {
	var a, b, c float64
	fmt.Println("Введите переменные a, b, c квадратного уравнения: ")
	fmt.Scanln(&a, &b, &c)

	d := (b * b) - (4 * a * c)
	if d == 0 {
		fmt.Println("Корень квадратного уравнения: ", -b/2*a)
	}
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		fmt.Println("Корни квадратного уравнения: x1=", x1, ", x2=", x2)
	}
	if d < 0 {
		realPart := -b / (2 * a)
		imaginaryPart := math.Sqrt(-d) / (2 * a)
		fmt.Printf("Уравнение имеет два комплексных корня: x1 = %.2f + %.2fi, x2 = %.2f - %.2fi\n", realPart, imaginaryPart, realPart, imaginaryPart)
	}
}

func task_1_3() {
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

	bubbleSortByAbsolute(slice)

	fmt.Println("Отсортированный массив по абсолютным значениям:", slice)
}

func bubbleSortByAbsolute(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Сравниваем абсолютные значения
			if math.Abs(float64(arr[j])) > math.Abs(float64(arr[j+1])) {
				// Меняем местами
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func task_1_4() {
	var len1, len2 int

	fmt.Println("Введите колличество элементов массива 1: ")
	fmt.Scanln(&len1)

	var slice1 []int = make([]int, len1)

	for index, r := range slice1 {
		fmt.Println("Введите ", index+1, " элемент массива: ")
		fmt.Scanln(&r)
		slice1[index] = r
	}

	fmt.Println("Введите колличество элементов массива 2: ")
	fmt.Scanln(&len2)

	var slice2 []int = make([]int, len2)

	for index, r := range slice2 {
		fmt.Println("Введите ", index+1, " элемент массива: ")
		fmt.Scanln(&r)
		slice2[index] = r
	}

	slice2 = append(slice2, slice1...)

	bubbleSortByAbsolute(slice2)

	fmt.Println("Отсортированный массив по абсолютным значениям:", slice2)
}

func task_1_5() {
	var a, b string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку: ")
	if scanner.Scan() {
		a = scanner.Text()
	}
	fmt.Println("Введите подстроку: ")
	if scanner.Scan() {
		b = scanner.Text()
	}

	result := strings.Index(a, b)
	fmt.Printf("Первое вхождение подстроки \"%s\" в строке \"%s\" : %d\n", b, a, result)
}

func task_2_1() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите оператор: ")
	fmt.Scan(&operator)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	switch operator {
	case "+":
		fmt.Println("a + b = ", a+b)
	case "-":
		fmt.Println("a - b = ", a-b)
	case "/":
		if b == 0 {
			fmt.Println("Деление на ноль невозвожно")
		} else {
			fmt.Println("a / b = ", a/b)
		}
	case "*":
		fmt.Println("a * b = ", a*b)
	case "%":
		fmt.Println("a % b = ", int(a)%int(b))
	case "^":
		fmt.Println("a ^ b = ", math.Pow(a, b))

	default:
		fmt.Println("Недопустимый оператор")
	}
}

func task_2_2() {
	var input string
	var runes []string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строку: ")
	if scanner.Scan() {
		input = scanner.Text()
	}
	fmt.Println(input)

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
		fmt.Println("Эта стркоа является палиндромом")
	} else {
		fmt.Println("Эта строка не является палиднромом")
	}
}

func task_2_5() {
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
	var x1a, x1b, x2a, x2b, x3a, x3b int

	fmt.Println("Введите координаты первого отрезка (x1a, x1b): ")
	fmt.Scan(&x1a, &x1b)
	fmt.Println("Введите координаты второго отрезка (x2a, x2b): ")
	fmt.Scan(&x2a, &x2b)
	fmt.Println("Введите координаты третьего отрезка (x3a, x3b): ")
	fmt.Scan(&x3a, &x3b)

	if x1a > x1b {
		x1a, x1b = x1b, x1a
	}
	if x2a > x2b {
		x2a, x2b = x2b, x2a
	}
	if x3a > x3b {
		x3a, x3b = x3b, x3a
	}

	intersectionStart := x1a
	if x2a > intersectionStart {
		intersectionStart = x2a
	}
	if x3a > intersectionStart {
		intersectionStart = x3a
	}

	intersectionEnd := x1b
	if x2b < intersectionEnd {
		intersectionEnd = x2b
	}
	if x3b < intersectionEnd {
		intersectionEnd = x3b
	}

	if intersectionStart <= intersectionEnd {
		fmt.Println("Область пересечения: xa= ", intersectionStart, " xb= ", intersectionEnd)
	} else {
		fmt.Println("Пересечение не существует.")
	}
}

func task_2_4() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите предложение: ")
	if scanner.Scan() {
		input = scanner.Text()
	}

	longest := longestWord(input)
	if longest != "" {
		fmt.Printf("Самое длинное слово: %s\n", longest)
	} else {
		fmt.Println("Нет слов в предложении.")
	}

}

func task_3_1() {
	var a, b, c, number int
	a = 0
	b = 1
	c = 0

	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Println("Количество чисел Фибоначи до заданного числа: ")

	if number == 0 {
		fmt.Println(0)
	}

	if number == 1 {
		fmt.Println(0)
		fmt.Println(1)
	}

	if number > 1 {
		fmt.Println(a)
		fmt.Println(b)

		for i := number - 2; i > 0; i-- {
			c = b
			b = a + b
			a = c

			if b > number {
				break
			}

			fmt.Println(b)
		}
	}
}

func task_3_2() {
	var number1 int
	var number2 int

	var k int = 0

	fmt.Print("Введите перввое число: ")
	fmt.Scan(&number1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&number2)
	fmt.Println("Все простые числа от ", number1, " до ", number2)

	for i := number1; i <= number2; i++ {
		for n := 1; n <= number2; n++ {
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

func task_3_3() {
	var a, b, sum int
	list_of_numbers := []int{}

	fmt.Print("Введите диапозон: ")
	fmt.Scan(&a, &b)
	fmt.Println("Числа Армстронга в заданном диапазоне: ")

	for i := a; i <= b; i++ {
		j := i
		for j > 0 {
			last_number := j % 10
			list_of_numbers = append(list_of_numbers, last_number)
			j /= 10
		}
		size := len(list_of_numbers)

		for _, r := range list_of_numbers {
			sum += int(math.Pow(float64(r), float64(size)))
		}

		if sum == i {
			fmt.Println(i)
		}

		sum = 0
		list_of_numbers = nil
	}
}

func task_3_4() {
	var input string

	fmt.Println("Введите строку: ")
	fmt.Scan(&input)

	runes := make([]rune, len(input))

	for i, r := range input {
		runes[len(input)-1-i] = r
	}

	fmt.Println("Строка в обратном порядке: ", string(runes))
}

func task_3_5() {
	var a, b, greater, less int
	fmt.Print("Введите два числа для нахождения НОД: ")
	fmt.Scan(&a, &b)

	if a > b {
		greater = a
		less = b
	} else {
		greater = b
		less = a
	}

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

// vlad mishustin

/////////////////////////////////////////////////////////
//////////////////// Дополнительные функции /////////////
/////////////////////////////////////////////////////////

func isPunctuation(char rune) bool {
	return unicode.IsPunct(char) || unicode.IsSymbol(char)
}

func longestWord(sentence string) string {
	// Удаляем знаки препинания и разбиваем строку на слова
	words := strings.FieldsFunc(sentence, func(r rune) bool {
		return isPunctuation(r) || unicode.IsSpace(r)
	})

	longest := ""
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}

	return longest
}
