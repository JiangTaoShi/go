package main

import (
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	//设置页面参数
	pdf := gofpdf.New("P", "cm", "A4", "")
	//添加一页
	pdf.AddPage()

	var infoPtr *gofpdf.ImageInfoType

	var imageFileStr, _ = filepath.Abs("04.jpg")
	//var imageFileStr, _ = filepath.Abs("out2.jpg")
	var imgWd, imgHt, lf, tp float64

	infoPtr = pdf.RegisterImage(imageFileStr, "")
	imgWd, imgHt = infoPtr.Extent()
	//infoPtr.SetDpi(300)
	// fmt.Println(imgWd)
	// fmt.Println(imgHt)
	// fmt.Println(infoPtr.Width())
	// fmt.Println(infoPtr.Height())
	// lf = 0
	// tp = 0
	imgHt = 25.4
	imgWd = 14.7
	pdf.Image(imageFileStr, lf, tp, imgWd, imgHt, false, "", 0, "")

	if err := pdf.OutputFileAndClose("hello.pdf"); err != nil {
		panic(err.Error())
	}

}
