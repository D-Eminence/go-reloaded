package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . sample.txt result.txt")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	extract := Extract(inputFile)
	parsed := ToArray(extract)
	modified := ModLoop(parsed)
	Quote := MergeQuote(modified)
	compiled := strings.Join(Quote, " ")

	Write(outputFile, compiled)
	fmt.Println("Compiled Successful")
}
func Extract(input string) string {
	content, _ := os.ReadFile(input)
	return string(content)
}
func ToArray(input string) []string {
	data := strings.Fields(input)
	return data
}
func Write(input string, content string) {
	os.WriteFile(input, []byte(content), 0664)
}
func ModLoop(input []string) []string {
	var result []string

	for i := 0; i < len(input); i++ {
		if input[i] == "(hex)" {
			lastIndex := len(result) - 1
			if lastIndex >= 0 {
				result[lastIndex] = JoinDec(result[lastIndex], 16, 64)
			}
		} else if input[i] == "(bin)" {
			lastIndex := len(result) - 1
			if lastIndex >= 0 {
				result[lastIndex] = JoinDec(result[lastIndex], 2, 64)
			}
		} else if input[i] == "(up)" {
			lastIndex := len(result) - 1
			if lastIndex >= 0 {
				result[lastIndex] = ToUpperOrToLower(result[lastIndex], true)
			}
		} else if input[i] == "(low)" {
			lastIndex := len(result) - 1
			if lastIndex >= 0 {
				result[lastIndex] = ToUpperOrToLower(result[lastIndex], false)
			}
		} else if input[i] == "(cap)" {
			lastIndex := len(result) - 1
			if lastIndex >= 0 {
				result[lastIndex] = ToCap(result[lastIndex])
			}
		} else if input[i] == "(cap," {
			num := Parser(input[i+1])
			count := 0
			for j := len(result) - 1; j >= 0 && count < num; j-- {
				if result[j] != " " {
					result[j] = ToCap(result[j])
					count++
				}
			}
			i++
		} else if input[i] == "(up," {
			num := Parser(input[i+1])
			count := 0
			for j := len(result) - 1; j >= 0 && count < num; j-- {
				if result[j] != " " {
					result[j] = ToUpperOrToLower(result[j], true)
					count++
				}
			}
			i++
		} else if input[i] == "(low," {
			num := Parser(input[i+1])
			count := 0
			for j := len(result) - 1; j >= 0 && count < num; j-- {
				if result[j] != " " {
					result[j] = ToUpperOrToLower(result[j], false)
					count++
				}
			}
			i++
		} else if IsPunc(input[i][0]) {
			last := len(result) - 1
			puncend := 0
			for puncend < len(input[i]) && IsPunc(input[i][puncend]) {
				puncend++
			}
			if last >= 0 {
				result[last] += input[i][:puncend]
			}
			if puncend < len(input[i]) {
				result = append(result, input[i][puncend:])
			}
		} else if input[i] == "a" {
			if IsVowel(input[i+1][0]) {
				result = append(result, "an")
			} else {
				result = append(result, "a")
			}
		} else if input[i] == "A" {
			if IsVowel(input[i+1][0]) {
				result = append(result, "An")
			} else {
				result = append(result, "A")
			}
		} else {
			result = append(result, input[i])
		}
	}

	return result
}

func JoinDec(input string, base int, size int) string {
	content, _ := strconv.ParseInt(input, base, size)
	return fmt.Sprintf("%d", content)
}

func ToCap(input string) string {
	return strings.ToUpper(input[:1]) + strings.ToLower(input[1:])
}

func ToUpperOrToLower(input string, name bool) string {
	if name {
		return strings.ToUpper(input)
	} else {
		return strings.ToLower(input)
	}
}
func Parser(input string) int {
	data := strings.TrimSuffix(input, ")")
	n, _ := strconv.Atoi(data)
	return n
}
func IsPunc(input byte) bool {
	data := strings.Contains(",.:;?!", string(input))
	return data
}
func IsVowel(input byte) bool {
	data := strings.Contains("aeiouhAEIOUH", string(input))
	return data
}

func MergeQuote(input []string) []string {
	start := 0

	for i := 0; i < len(input); i++ {
		if input[i] == "'" {
			if start == 0 {
				start = i
			} else {
				end := i
				input[start] = "'" + strings.Join(input[start+1:end], " ") + "'"
				input = append(input[:start+1], input[end+1:]...)
				i = start
				start = 0
			}
		}
	}
	return input
}
