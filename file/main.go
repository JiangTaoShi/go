package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {

	//imagePath, _ := filepath.Abs("./images/5.jpg")
	//file, err := os.Open(imagePath)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//fileinfo, err := file.Stat()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//filesize := fileinfo.Size()
	//buffer := make([]byte, filesize)
	//file.Read(buffer)
	//
	//for i := 151; i <= 180; i++ {
	//	newImagePath, _ := filepath.Abs(fmt.Sprintf("./images/%s.jpg", fmt.Sprintf("%03d", i)))
	//	newFile, err := os.Create(newImagePath)
	//	if err != nil {
	//		panic(err)
	//	}
	//	newFile.Write(buffer)
	//	newFile.Close()
	//}

	path1, _ := filepath.Abs("./images")
	path2, _ := filepath.Abs("./images2")
	fmt.Println(path1)
	fmt.Println(path2)
	copyDir(path1, path2)
}

func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		return s
	}
}

func copyDir(src string, dest string) {
	src = FormatPath(src)
	dest = FormatPath(dest)
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("xcopy", src, dest, "/I", "/E")
	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}
	_, e := cmd.Output()
	if e != nil {
		fmt.Println(e.Error())
		return
	}
}
