package parse

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Files() (files []string, err error) {
	var root string
	if _, err = fmt.Scanln(&root); err != nil {
		return
	}

	err = filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	return
}
