package selector

import (
	"FileNameRefactoring/method"
	"FileNameRefactoring/text"
	"errors"
	"fmt"
	"regexp"
)

func Selector(re *regexp.Regexp, files []string) (err error) {
	var mtd, part int
	fmt.Println(text.SelectorMethod)
	if _, err = fmt.Scanln(&mtd); err != nil {
		return
	}
	fmt.Println(text.SelectorPart)
	if _, err = fmt.Scanln(&part); err != nil {
		return
	}

	switch mtd {
	case 0:
		return
	case 1:
	case 2:
		switch part {
		case 0:
			return Selector(re, files)
		case 1:
			if err = method.DeleteAll(re, files); err != nil {
				return
			}
		case 2:
		case 3:
		}
	case 3:
	default:
		err = errors.New(text.ErrUnknownInput)
	}

	return
}
