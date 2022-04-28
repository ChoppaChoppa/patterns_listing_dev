package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	A int
	B int
	C int
	c bool
	i bool
	v bool
	n bool
	F bool
)

func main() {
	flag.IntVar(&A, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "количество строк")
	flag.BoolVar(&i, "i", false, "игнорировать регистр")
	flag.BoolVar(&v, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&F, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, "напечатать номер строки")
	flag.Parse()

	// /Users/moratherest/Projects/wb_l2/dev/task_5/Text.txt
	arr := OpenFile(flag.Arg(1))
	regex, errReg := regexp.Compile(flag.Arg(0))
	if errReg != nil {
		fmt.Printf("regex wrong: %v", errReg)
	}

	if i {
		ConvertToLowerCase(arr)
	}

	var val map[int]string
	if F {
		val = ExactMatch(arr, os.Args[len(os.Args)-2])
	} else if A > 0 {
		val = After(arr, regex, A)
	} else if B > 0 {
		val = Before(arr, regex, B)
	} else if C > 0 {
		val = Context(arr, regex, C, C)
	} else {
		val = grep(arr, regex, v)
	}

	var keys = make([]int, 0, len(val))
	for k := range val {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	if c {
		fmt.Println(len(val))
	} else if n {
		for _, key := range keys {
			fmt.Println(key, val[key])
		}
	} else {
		for _, key := range keys {
			fmt.Println(val[key])
		}
	}
}

func OpenFile(path string) []string {
	var arr []string

	file, errOpen := os.Open(path)
	if errOpen != nil {
		code, _ := fmt.Fprintln(os.Stderr, errOpen)
		os.Exit(code)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	return arr
}

// Выводим совместимые с паттерном
// Сразу выполняется проверка на получение совпадений или не совпадений
func grep(arr []string, reg *regexp.Regexp, invert bool) map[int]string {
	var mapa = make(map[int]string)
	if invert {
		for index, val := range arr {
			if !reg.MatchString(val) {
				mapa[index] = val
			}
		}
	} else {
		for index, val := range arr {
			if reg.MatchString(val) {
				mapa[index] = val
			}
		}
	}

	return mapa
}

//After Находит совпадение с паттерном и запускает цикл, который выводит следующие num строк
func After(arr []string, reg *regexp.Regexp, num int) map[int]string {
	var mapa = make(map[int]string)

	for index, val := range arr {
		if reg.MatchString(val) {
			for j := 0; j <= num && index != len(arr); j++ {
				mapa[index] = arr[index]
				index++
			}

			break
		}
	}

	return mapa
}

//Before тоже самое, что и After, но в цикл в обратном порядке
func Before(arr []string, reg *regexp.Regexp, num int) map[int]string {
	var mapa = make(map[int]string)

	for index, val := range arr {
		if reg.MatchString(val) {
			for j := num; j >= 0 && index > 0; j-- {
				mapa[index] = arr[index]
				index--
			}
			break
		}
	}

	return mapa
}

//Context выполняет After и Before, после объединяет их
func Context(arr []string, reg *regexp.Regexp, after, before int) map[int]string {
	mapaAfter := After(arr, reg, 3)
	mapaBefore := Before(arr, reg, 3)

	//если есть одинаковые ключи строка пропускается, чтобы избежать дубликатов
	for key, value := range mapaAfter {
		if _, ok := mapaBefore[key]; ok {
			continue
		}

		mapaBefore[key] = value
	}

	return mapaBefore
}

//GetLineCount Получить кол-во строк
func GetLineCount(mapa map[int]string) int {
	return len(mapa)
}

func ConvertToLowerCase(arr []string) {
	for index, val := range arr {
		arr[index] = strings.ToLower(val)
	}
}

//ExactMatch поиск по полному совпадению со строкой
func ExactMatch(arr []string, substr string) map[int]string {
	mapa := make(map[int]string)
	var matchCount int
	length := len(substr)
	substrIndex := 0

	for index, val := range arr {
		for k := 0; k < len(val); k++ {
			if val[k] == substr[substrIndex] {
				substrIndex++
				matchCount++
			} else {
				substrIndex = 0
				matchCount = 0
			}
			if matchCount == length {
				mapa[index] = val
				matchCount = 0
				substrIndex = 0
				break
			}
		}
	}

	return mapa
}
