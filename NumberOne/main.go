package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Задание номер 1")
	fmt.Println(WordCount("go     is     fun     go"))
	fmt.Println(WordCount("Hello hello    World"))
	fmt.Println("---------------------------")

	fmt.Println("Задание номер 2")
	fmt.Println(AreAnagrams("listen", "silent"))
	fmt.Println(AreAnagrams("Eleven plus two", "Twelve plus one"))
	fmt.Println(AreAnagrams("Hello", "World"))
	fmt.Println("---------------------------")

	fmt.Println("Задание номер 3")
	fmt.Printf("%c\n", FirstUnique("abacabad"))
	fmt.Printf("%c\n", FirstUnique("abcabcd"))
	fmt.Println(FirstUnique("aabbcc"))
	fmt.Println("---------------------------")

	fmt.Println("Задание номер 4")
	fmt.Println(RemoveDuplicates([]int{1, 2, 2, 3, 1}))
	fmt.Println(RemoveDuplicates([]int{4, 4, 4, 4}))
	fmt.Println("---------------------------")

	fmt.Println("Задание номер 5")
	fmt.Println(RemoveElement([]int{1, 2, 2, 3, 1}, 3))
	fmt.Println(RemoveElement([]int{1, 2, 3}, 5))
	fmt.Println("---------------------------")

	fmt.Println("Задание номер 6")
	fmt.Println(IsPalindrome("шалаш"))
	fmt.Println(IsPalindrome("А роза упала на лапу Азора"))
	fmt.Println(IsPalindrome("hello"))
	fmt.Println("---------------------------")

	fmt.Println("Задание номер 7")
	PrintChessboard()
}

func WordCount(s string) map[string]int {
	arrayStr := strings.Fields(s)
	result := make(map[string]int)

	for _, str := range arrayStr {
		if value, ok := result[str]; ok {
			result[str] = value + 1
		} else {
			result[str] = 1
		}
	}

	return result
}

func AreAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	str1 := strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	str2 := strings.ToLower(strings.ReplaceAll(s2, " ", ""))

	str1ToMap := stringToMap(str1)
	str2ToMap := stringToMap(str2)

	for char := range str1ToMap {
		if _, ok := str2ToMap[char]; ok {
			if str1ToMap[char] != str2ToMap[char] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

type countAndPosition struct {
	count int
	index int
}

func FirstUnique(s string) rune {
	strToMap := make(map[rune]countAndPosition)

	for index, char := range s {
		if value, ok := strToMap[char]; ok {
			value.count++
			value.index = -1
			strToMap[char] = value
		} else {
			value.count++
			value.index = index
			strToMap[char] = value
		}
	}
	var res rune = 0
	minIndex := math.MaxInt

	for char := range strToMap {
		if strToMap[char].count == 1 && strToMap[char].index < minIndex {
			minIndex = strToMap[char].index
			res = char
		}
	}
	return res
}

func RemoveDuplicates(nums []int) []int {
	var resultSlice []int
	tempMap := make(map[int]int)

	for _, num := range nums {
		if _, isExists := tempMap[num]; !isExists {
			resultSlice = append(resultSlice, num)
			tempMap[num] = num
		}
	}

	return resultSlice
}

func RemoveElement(nums []int, index int) ([]int, error) {
	if len(nums)-1 < index {
		return nil, errors.New("index out of range")
	}

	leftSlice := nums[:index]
	rightSlice := nums[index+1:]
	resultSlice := append(leftSlice, rightSlice...)

	return resultSlice, nil
}

func IsPalindrome(s string) bool {
	var strBuilder strings.Builder

	for _, char := range s {
		if unicode.IsLetter(char) {
			strBuilder.WriteRune(unicode.ToLower(char))
		}
	}

	str := strBuilder.String()
	strRune := []rune(str)
	lenStr := len(strRune)

	for i := 0; i < (lenStr-1)/2; i++ {
		if strRune[i] != strRune[lenStr-1-i] {
			return false
		}
	}

	return true
}

func PrintChessboard() {
	const boardSize = 8

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func stringToMap(s string) map[rune]int {
	result := make(map[rune]int)

	for _, char := range s {
		if value, ok := result[char]; ok {
			result[char] = value + 1
		} else {
			result[char] = 1
		}
	}

	return result
}
