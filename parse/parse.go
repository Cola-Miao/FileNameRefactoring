package parse

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const DefaultBufSize = 5 * 1024

func Files() (files []string, err error) {
	var buf = make([]byte, DefaultBufSize)
	if _, err = io.ReadFull(os.Stdin, buf); err != nil {
		if errors.Is(err, io.ErrUnexpectedEOF) {
			err = nil
		} else {
			return
		}
	}
	fmt.Scanln()
	filesPath := string(buf)
	if filesPath[0] == '"' {
		files = splitFilesExistSpace(filesPath)
	} else {
		files = strings.Split(filesPath, " ")
	}

	return
}

func splitFilesExistSpace(filePath string) (files []string) {
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
