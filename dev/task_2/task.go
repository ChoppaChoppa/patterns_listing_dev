package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := UnPackStr("a4bc2d5e")
	fmt.Println(str)
}

func UnPackStr(str string) string {
	//проверка на пустую строку
	if str == "" {
		return ""
	}

	//проверка, является ли первый символ числом
	if _, err := strconv.Atoi(string(str[0])); err == nil {
		fmt.Println(string(str[0]))
		return "error"
	}

	var outStr string
	for i := 0; i < len(str); i++ {
		var nums string
		var char = string(str[i])
		_, IsDigit := strconv.Atoi(char)

		if IsDigit == nil {
			charRepeat := string(str[i-1])

			// получение символов, которые являются числом,
			// пока не встретится символ не являющийся числом
			for {
				nums += string(str[i])
				if i+1 >= len(str) {
					break
				}

				_, IsDigit = strconv.Atoi(string(str[i+1]))
				if IsDigit != nil {
					break
				}
				i++
			}
			// перевод чисел из string в int
			num, _ := strconv.Atoi(nums)
			// умножения символа на число
			outStr += strings.Repeat(charRepeat, num-1)

		} else if char == "\\" {
			// если символ является \, то в строку добавляется идущий за ней символ
			// i увеличивается здесь и еще раз в условии цикла,
			// из-за чего добавленный символ пропускается
			outStr += string(str[i+1])
			i++
			continue
		} else {
			outStr += char
		}
	}

	return outStr
}

// TODO: является ли возможным и разумным сделать алгоритм через связанный список?

type Char struct {
	C       string
	IsDigit bool
}
type Word struct {
	This Char
	Next Char
}
