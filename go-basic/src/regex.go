package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "banana burger soup"
	regex, _ := regexp.Compile(`[a-z]+`)
	fmt.Println(regex.MatchString(text))
	search := regex.FindAllString(text, -1)
	fmt.Println(search)

	var regex2 *regexp.Regexp = regexp.MustCompile("b([a-z])i")
	fmt.Println(regex2.MatchString("eko"))
	fmt.Println(regex2.MatchString("bai"))
	fmt.Println(regex2.MatchString("buba"))
}
