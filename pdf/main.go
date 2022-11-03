package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	//设置页面参数
	pdf := gofpdf.New("P", "cm", "A4", "")

	//添加一页
	pdf.AddPage()

	//将图片放入到 pdf 文档中
	//ImageOptions(src, x, y, width, height, flow, options, link, linkStr)

	var infoPtr *gofpdf.ImageInfoType
	var imageFileStr = "C:/Work/Jiangtao/go/pdf/02.jpg"
	var imgWd, imgHt, lf, tp float64

	infoPtr = pdf.RegisterImage(imageFileStr, "")
	infoPtr.SetDpi(96)
	imgWd, imgHt = infoPtr.Extent()
	fmt.Println(imgWd)
	fmt.Println(imgHt)
	fmt.Println(infoPtr.Width())
	fmt.Println(infoPtr.Height())
	lf = 0
	tp = 0
	imgHt = 0
	imgWd = 0
	pdf.Image(imageFileStr, lf, tp, imgWd, imgHt, false, "", 0, "")

	if err := pdf.OutputFileAndClose("hello.pdf"); err != nil {
		panic(err.Error())
	}

}
