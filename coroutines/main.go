package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10 * 12)

	for i := 1; i < 13; i++ {
		imagePath, _ := filepath.Abs(fmt.Sprintf("./images/%d.jpg", i))
		file, err := os.Open(imagePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fileinfo, err := file.Stat()
		if err != nil {
			fmt.Println(err)
			return
		}
		filesize := fileinfo.Size()
		buffer := make([]byte, filesize)
		file.Read(buffer)

		for ii := i * 10; ii < (i*10 + 10); ii++ {
			go func(val int) {
				defer wg.Done()
				newImagePath, _ := filepath.Abs(fmt.Sprintf("./images2/%s.jpg", fmt.Sprintf("%05d", val)))
				newFile, err := os.Create(newImagePath)
				if err != nil {
					panic(err)
				}
				newFile.Write(buffer)
				newFile.Close()
			}(ii)
		}
	}

	wg.Wait()

}

func fmtp(str int) {
	defer wg.Done()
	fmt.Println(str)
}
