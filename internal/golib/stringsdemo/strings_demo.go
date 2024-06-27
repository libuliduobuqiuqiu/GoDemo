package stringsdemo

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckStringIndex() {
	text := `"hello,zhangsan"`
	text = strings.TrimSpace(text)

	start := strings.Index(text, `"`)
	fmt.Println(text[start+1:])

	end := strings.Index(text[start+1:], `"`)
	fmt.Println(start, end)
	fmt.Println(text[start+1 : end+1])

	if end == -1 {
		fmt.Println("text error.")
	} else {
		fmt.Println(text[start+1 : end+1])
	}
}

func RegexUrl(row string) {
	r3, _ := regexp.Compile(`\[\s*HTTP::[a-z]+\s*]\s+starts_with`)
	if r3.MatchString(row) {
		str := strings.Split(row, "starts_with ")
		if len(str) > 1 {
			url := strings.TrimSpace(str[1])
			start := strings.Index(url, `"`)
			end := strings.Index(url[start+1:], `"`)
			if start == -1 || end == -1 {
				fmt.Println("error")
				return
			}
			fmt.Println(start+1, start+end+1)
			fmt.Println(url[start+1 : start+end+1])
		}
	}
}

func CheckUrlIndex() {
	text := `elseif ( [matchclass [HTTP::uri] starts_with ebank_off_group_801 ]} (HTTP::redirect "https://www.qsbank.cc/")}`
	text2 := `if { [HTTP::uri] starts_with "/xxx/" } { pool DC01_801_ebank_POL_443 }`
	RegexUrl(text)
	RegexUrl(text2)
	fmt.Println(text[0:3])
}
