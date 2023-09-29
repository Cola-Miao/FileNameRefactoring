package method

import (
	"FileNameRefactoring/system"
	"FileNameRefactoring/text"
	"errors"
	"fmt"
	"regexp"
)

func DeleteAll(re *regexp.Regexp, files []string) (err error) {
	compare := make(map[string]string)
	for _, file := range files {
		if err = checkFileConflict(compare, file); err != nil {
			return
		}
		compare[file] = re.ReplaceAllString(file, "")
	}

	showCompare(compare)
	if err = confirmModify(); err != nil {
		return
	}

	system.ModifyFilePath(compare)

	return
}

func confirmModify() (err error) {
	fmt.Println(text.WaitConfirm)
	var yes string
	if _, err = fmt.Scanln(&yes); err != nil {
		return
	}
	if yes != "YES" {
		err = errors.New(text.ErrUnknownInput)
	}

	return
}

func checkFileConflict(compare map[string]string, file string) (err error) {
	if _, ok := compare[file]; ok {
		err = errors.New(text.ErrFileNameConflict + ":" + file)
	}

	return
}

func showCompare(compare map[string]string) {
	for oldF, newF := range compare {
		fmt.Println(oldF, "-->", newF)
	}
}
