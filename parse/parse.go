package parse

import (
	"bufio"
	"os"
)

func Files() (files []string, err error) {
	var filesPath string
	fp := bufio.NewReader(os.Stdin)
	filesPath, err = fp.ReadString('\n')
	if err != nil {
		return
	}
	files = splitFiles(filesPath)

	return
}

func splitFiles(filePath string) (files []string) {
	l := len(filePath)
	for i := 0; i < l; i++ {
		if filePath[i] != '"' {
			continue
		}
		var j int
		for j = i + 1; j < l && filePath[j] != '"'; j++ {
		}
		if i < l && j < l {
			files = append(files, filePath[i+1:j])
		}
		i = j + 1
	}

	return
}
