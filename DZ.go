// На основе шаблона напишите программу, подсчитывающую сколько раз буква встречается в предложении, а также частоту встречаемости в процентах. То есть в предложении hello world результатом будет:
// h - 1 0.1
// e - 1 0.1
// l - 3 0.33
// o - 2 0.2
// w - 1 0.1
// r - 1 0.1
// d - 1 0.1

package main

import (
	"fmt"
	"strings"
)

func getRuneInString(rune rune, str string) int {
	var coin int = 0
	for _, i := range str {
		if rune == i {
			coin++
		}
	}
	return coin
}

func OnlySymbol(s string) string {
	var str string
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			continue
		}
		str = str + string(r)
	}
	return str
}
func getCountRuneInString(s string) map[rune]int {
	result := make(map[rune]int)
	for len(s) > 0 {
		result[rune(s[0])] = getRuneInString(rune(s[0]), s)
		s = strings.ReplaceAll(s, string(s[0]), "")
	}
	return result
}

func main() {
	var str string = "Hello World     hgch.,./"
	str = strings.ToLower(str)
	fmt.Println(str)
	str = OnlySymbol(str)
	for k, v := range getCountRuneInString(str) {
		fmt.Printf("Char  %c    Count in string  %d      Percent %0.2f \n", k, v, float64(v)/float64(len(str))*100)
	}

}
