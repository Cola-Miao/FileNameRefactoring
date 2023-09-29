package main

import (
	"FileNameRefactoring/parse"
	"FileNameRefactoring/selector"
	"FileNameRefactoring/text"
	"fmt"
	"log"
	"regexp"
)

func main() {
	log.SetFlags(log.Llongfile)
	fmt.Println(text.WaitREGEXP)
	re, err := getREGEXP()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(text.WaitFiles)
	files, err := parse.Files()
	if err != nil {
		log.Println(err)
	}

	if err = selector.Selector(re, files); err != nil {
		log.Println(err)
		return
	}
}

func getREGEXP() (re *regexp.Regexp, err error) {
	var expr string
	if _, err = fmt.Scanln(&expr); err != nil {
		return
	}
	re, err = regexp.Compile(expr)

	return
}
