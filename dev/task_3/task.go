package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	unsuitable []string

	fileName string
	k        string
	n        bool
	r        bool
	u        bool
)

//TODO: убрать возвращаемый арр

func main() {
	flag.StringVar(&fileName, "name", "", "имя файла")
	flag.StringVar(&k, "k", " ", "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	file := OpenFile(fileName)

	arr := strings.Split(file, k)
	LowerCase(arr)

	if u {
		arr = RemoveDuplicate(arr)
	}
	arr = Sort(arr, 0, len(arr)-1)

	fmt.Println(arr, cap(arr))
}

//RemoveDuplicate удаляет повторяющиеся строки
func RemoveDuplicate(arr []string) []string {
	checkedChar := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		_, ok := checkedChar[arr[i]]
		if ok {
			if i != len(arr) {
				arr = append(arr[:i], arr[i+1:]...)
				i--
			}
		}

		checkedChar[arr[i]] = i
	}
	return arr
}

//LowerCase приводит строки к нижнему регистру
func LowerCase(arr []string) {
	for i := range arr {
		arr[i] = strings.ToLower(arr[i])
	}
}

func OpenFile(FileName string) string {
	file, errOpen := os.Open(FileName)
	if errOpen != nil {
		code, _ := fmt.Fprintln(os.Stderr, errOpen)
		os.Exit(code)
	}
	defer file.Close()

	wr := bytes.Buffer{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wr.WriteString(scanner.Text())
	}

	return wr.String()
}

func Sort(arr []string, start, end int) []string {
	arr = quickSort(arr, start, end)

	for _, v := range unsuitable {
		arr = append(arr, v)
	}

	return arr
}

func quickSort(arr []string, start, end int) []string {
	arr, left, right := parti(arr, start, end)

	if start < right {
		arr = quickSort(arr, start, right)
	}
	if left < end {
		arr = quickSort(arr, left, end)
	}

	return arr
}

func parti(arr []string, start, end int) ([]string, int, int) {
	var mid int
	var midIsDigit bool
	if n {
		mid, midIsDigit = getNum(arr[(start+end)/2])
		for !midIsDigit {
			arr = RemoveIfIsNotDigit(arr, (start+end)/2)
			mid, midIsDigit = getNum(arr[(start+end)/2])
		}
	} else {
		mid = (start + end) / 2
	}

	markerLeft := start
	markerRight := end

	for i := start; markerLeft < markerRight; i++ {
		// получаем маркер в зависимости от аргумента -r
		// если передать в функцию getMarker параметры compareForLeft = -1 и compareForRight = 1
		// сортировка произойдет в алфавитном порядке
		// если присвоить 1 и -1 соответственно, сортировка осуществится в обратном порядке
		// TODO: вынести reverse в отдельную функцию
		if n {
			arr, markerLeft, markerRight = getMarkerForNum(arr, markerLeft, markerRight, mid)
		} else {
			var compareLeft, compareRight = 1, -1
			if !r {
				compareLeft, compareRight = -1, 1
			}

			markerLeft, markerRight = getMarkerForString(arr, markerLeft, markerRight, mid, compareLeft, compareRight)
		}

		fmt.Println("marker", markerLeft, markerRight)
		if markerLeft <= markerRight {
			temp := arr[markerLeft]
			arr[markerLeft] = arr[markerRight]
			arr[markerRight] = temp
			markerLeft++
			markerRight--
		}
	}

	return arr, markerLeft, markerRight
}

//getMarkerForString выдает индекс для элементов
func getMarkerForString(arr []string, left, right, mid, compareForLeft, compareForRight int) (int, int) {
	for strings.Compare(arr[left], arr[mid]) == compareForLeft {
		left++
	}
	for strings.Compare(arr[right], arr[mid]) == compareForRight {
		right--
	}

	return left, right
}

//getMarkerForNum Выдает индекс для элементов, перед которыми стоят числа
func getMarkerForNum(arr []string, left, right, mid int) ([]string, int, int) {
	fmt.Println(mid, arr)
	//if !haveNum {
	//	fmt.Println("в списке нет цифр")
	//	os.Exit(1)
	//}

	leftNum, leftIsDigit := getNum(arr[left])
	rightNum, rightIsDigit := getNum(arr[right])

	for leftNum < mid {
		//удаляем элемент если он не начинается с числа
		if !leftIsDigit {
			arr = RemoveIfIsNotDigit(arr, left)
			leftNum, leftIsDigit = getNum(arr[left])
			continue
		}
		left++
		leftNum, leftIsDigit = getNum(arr[left])
	}
	for rightNum > mid || !rightIsDigit {
		if !rightIsDigit {
			if right == len(arr)-1 && right > 1 {
				right--
			}
			arr = RemoveIfIsNotDigit(arr, right)
			rightNum, rightIsDigit = getNum(arr[right])
			continue
		}
		right--
		rightNum, rightIsDigit = getNum(arr[right])
	}

	return arr, left, right
}

//Возвращает числа стоящие перед строкой
func getNum(str string) (int, bool) {
	var num string
	for i := 0; i < len(str); i++ {
		if 48 <= str[i] && str[i] <= 57 {
			num += string(str[i])
		} else {
			if i == 0 {
				return 0, false
			}
			break
		}
	}

	nums, _ := strconv.Atoi(num)
	return nums, true
}

// RemoveIfIsNotDigit обнуляет значение и добавляет его в конец
// TODO: значения по итогу оказываются по середине
func RemoveIfIsNotDigit(arr []string, index int) []string {
	unsuitable = append(unsuitable, arr[index])
	copy(arr[index:], arr[index+1:])
	arr[len(arr)-1] = ""
	return arr
}
