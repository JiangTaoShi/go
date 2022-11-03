package main

import (
	"fmt"

	"github.com/signintech/gopdf"
)

func main() {
	//创建pdf
	pdf := gopdf.GoPdf{}
	//pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 25.6, H: 15.1}, Unit: gopdf.Unit_CM})
	//25.6/15.1

	for i := 1; i < 14; i++ {
		pdf.AddPage()
		//添加图片
		pdf.Image(fmt.Sprintf("%s.jpg", fmt.Sprintf("%02d", i)), 0, 0, &gopdf.Rect{W: 25.6, H: 15.1})
	}

	//save to file
	pdf.WritePdf("sample.pdf")

}
