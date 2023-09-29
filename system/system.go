package system

import (
	"log"
	"os"
	"sync"
)

func ModifyFilePath(compare map[string]string) {
	var wg sync.WaitGroup

	for oldF, newF := range compare {
		wg.Add(1)
		go func(oldF, newF string) {
			defer wg.Done()
			if err := os.Rename(oldF, newF); err != nil {
				log.Println(err, oldF, newF)
			}
		}(oldF, newF)
	}
	wg.Wait()

	return
}
