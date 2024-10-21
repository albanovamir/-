package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
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
		case 11:
			task_11()
		case 12:
			task_12()
		case 13:
			task_13()
		case 14:
			task_14()
		case 15:
			task_15()

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

func task_1() {
	var number, del, count int
	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	for i := 2; i <= number; i++ {
		if number%i == 0 {
			count += 1
			if del == 0 {
				del = i
			}
		}
	}

	if count == 1 {
		fmt.Println("Введенное число является простым")
	} else {
		fmt.Println("Введенное число не является простым")
		fmt.Println("Первый найденный делитель: ", del)
	}
}

func task_2() {
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

func task_3() {
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
				fmt.Println(arr)
			}
		}
	}
}

func task_4() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}
}

func task_5() {
	var number int

	fmt.Println("Введите номер числа фибоначи: ")
	fmt.Scanln(&number)
	fmt.Println(fibonacci(number))
}
func fibonacci(n int) int {
	var memo = make(map[int]int)
	if val, ok := memo[n]; ok {
		return val
	}
	if n <= 1 {
		return n
	}
	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func task_6() {
	var number string
	var reverse []string
	var rev_string string = ""

	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	for _, r := range number {
		reverse = append(reverse, string(r))
	}

	for i := 0; i < len(reverse); i++ {
		rev_string = rev_string + reverse[len(reverse)-1-i]
	}

	fmt.Println("Число в обратном порядке: ", rev_string)
}

func task_7() {
	var levels int
	fmt.Print("Введите уровень треугольника Паскаля: ")
	fmt.Scanln(&levels)

	pascalTriangle := make([][]int, levels)

	for i := 0; i < levels; i++ {
		pascalTriangle[i] = make([]int, i+1)
		pascalTriangle[i][0] = 1
		pascalTriangle[i][i] = 1

		for j := 1; j < i; j++ {
			pascalTriangle[i][j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}
	}

	for i := 0; i < levels; i++ {
		for j := 0; j < levels-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", pascalTriangle[i][j])
		}
		fmt.Println()
	}
}

func task_8() {
	var number, count int
	var slice1 []int

	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	for i := 0; i == 0; {
		if number/10 > 0 {
			slice1 = append(slice1, number%10)
			number /= 10
		}
		if number/10 == 0 {
			slice1 = append(slice1, number%10)
			i = 1
		}
	}

	var slice2 []int = make([]int, len(slice1))

	for index, r := range slice1 {
		slice2[len(slice2)-1-index] = r
	}

	for index, r := range slice1 {
		if r == slice2[index] {
			count++
		}
	}

	if count == len(slice1) {
		fmt.Println("Число является палиндромом")
	} else {
		fmt.Println("Число не является палиндромом")
	}
}

func task_9() {
	var len, max, min int

	fmt.Println("Введите колличество элементов массива: ")
	fmt.Scanln(&len)

	var slice []int = make([]int, len)

	for index, r := range slice {
		fmt.Println("Введите ", index+1, " элемент массива: ")
		fmt.Scanln(&r)
		slice[index] = r
	}

	max = slice[0]
	min = slice[0]

	for _, r := range slice {
		if max < r {
			max = r
		}
		if min > r {
			min = r
		}
	}

	fmt.Println("Маибольший элемент массива: ", max)
	fmt.Println("Наименьший элемент массива: ", min)
}

func task_10() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 10

	fmt.Println("Угадайте число от 1 до 100! У вас есть", attempts, "попыток.")

	for i := 1; i <= attempts; i++ {
		var guess int
		fmt.Printf("Попытка %d: Введите ваше число: ", i)
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Ошибка ввода. Пожалуйста, введите число.")
			i--
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("Число должно быть в диапазоне от 1 до 100.")
			i--
			continue
		}

		if guess < target {
			fmt.Println("Слишком мало! Попробуйте больше.")
		} else if guess > target {
			fmt.Println("Слишком много! Попробуйте меньше.")
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d с %d попытки!", target, i)
			fmt.Println("")
			return
		}
	}

	fmt.Printf("Вы исчерпали все попытки! Загаданное число было %d.n", target)
	fmt.Println("")
}

func task_11() {
	var number, j, sum int
	list_of_numbers := []int{}

	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	j = number

	for j > 0 {
		last_number := j % 10
		list_of_numbers = append(list_of_numbers, last_number)
		j /= 10
	}
	size := len(list_of_numbers)

	for _, r := range list_of_numbers {
		sum += int(math.Pow(float64(r), float64(size)))
	}

	if sum == number {
		fmt.Println("Данное число является числом Армстронга")
	} else {
		fmt.Println("Данное число не является числом Армстронга")
	}
}

func task_12() {
	fmt.Print("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	countUniqueWords(inputString)
}
func countUniqueWords(inputStr string) {
	words := strings.Fields(strings.ToLower(inputStr))

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	uniqueWordCount := len(wordCount)
	fmt.Println("Количество уникальных слов:", uniqueWordCount)
}

func task_13() {
	board := [][]bool{
		{true, false, false, false, false, true, false, false, false, false},
		{true, false, true, true, false, false, false, false, true, false},
		{false, true, true, true, false, false, false, true, false, false},
		{true, true, true, false, false, true, true, false, false, false},
		{false, false, true, false, false, false, false, false, false, false},
		{false, true, true, true, false, false, true, false, true, false},
		{false, false, false, false, true, false, false, false, true, false},
		{false, false, true, false, false, true, false, false, true, false},
		{true, false, false, false, false, false, true, false, true, true},
		{false, false, false, false, true, false, false, false, true, false},
	}

	// Основной цикл
	for {
		printBoard(board)
		board = updateBoard(board)
		time.Sleep(500 * time.Millisecond) // Задержка между поколениями
	}
}

const (
	width  = 10
	height = 10
)

// Функция для отображения текущего состояния поля
func printBoard(board [][]bool) {
	for _, row := range board {
		for _, cell := range row {
			if cell {
				fmt.Print("o") // Живая клетка
			} else {
				fmt.Print(" ") // Мертвая клетка
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Функция для подсчета живых соседей
func liveNeighbors(board [][]bool, x, y int) int {
	count := 0
	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < height && ny >= 0 && ny < width && board[nx][ny] {
			count++
		}
	}
	return count
}

// Функция обновления состояния поля
func updateBoard(board [][]bool) [][]bool {
	newBoard := make([][]bool, height)
	for i := range newBoard {
		newBoard[i] = make([]bool, width)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			livingNeighbors := liveNeighbors(board, i, j)
			if board[i][j] {
				newBoard[i][j] = livingNeighbors == 2 || livingNeighbors == 3
			} else {
				newBoard[i][j] = livingNeighbors == 3
			}
		}
	}
	return newBoard
}

func task_14() {
	var number, sum, result int
	var slice1 []int

	fmt.Println("Введите число: ")
	fmt.Scanln(&number)

	for i := 0; i == 0; {
		if number/10 > 0 {
			slice1 = append(slice1, number%10)
			number /= 10
		}
		if number/10 == 0 {
			slice1 = append(slice1, number%10)
			i = 1
		}
	}

	for i := 0; i == 0; {
		if len(slice1) > 1 {
			for _, r := range slice1 {
				sum += r
			}
		} else {
			fmt.Println("Цифровой корень числа = ", result)
			break
		}

		slice1 = nil

		for i := 0; i == 0; {
			if sum/10 > 0 {
				slice1 = append(slice1, sum%10)
				sum /= 10
			}
			if sum/10 == 0 {
				slice1 = append(slice1, sum%10)
				i = 1
			}
		}
		result = sum
		sum = 0
	}
}

func task_15() {
	var number int

	fmt.Print("Введите арабское число: ")
	fmt.Scanln(&number)

	if number < 1 || number > 3999 {
		fmt.Println("Пожалуйста, введите число в диапазоне от 1 до 3999.")
		return
	}

	romanNumeral := arabicToRoman(number)
	fmt.Println("Римское число: ", romanNumeral)
}
func arabicToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	i := 0

	for num > 0 {
		for num >= val[i] {
			roman += sym[i]
			num -= val[i]
		}
		i++
	}

	return roman
}
