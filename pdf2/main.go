package main

import (
	"github.com/signintech/gopdf"
)

func main() {
	//创建pdf

	//pt W: 2250, H: 1302
	//px W: 3000, H: 1726

	//PDF cm单位尺寸 / 2.54 * DPI = PDF像素

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 25.4, H: 14.7}, Unit: gopdf.UnitCM})

	pdf.AddPage()
	//添加图片
	pdf.Image("out2.jpg", 0, 0, &gopdf.Rect{W: 25.4, H: 14.7})

	// pdf.AddPage()
	// //添加图片
	// pdf.Image("01.jpg", 0, 0, &gopdf.Rect{W: 25.4, H: 14.7})

	// pdf.AddPage()
	// //添加图片
	// pdf.Image("01_300.jpg", 0, 0, &gopdf.Rect{W: 25.4, H: 14.7})

	pdf.AddPage()
	//添加图片
	pdf.Image("output.jpg", 0, 0, &gopdf.Rect{W: 25.4, H: 14.7})

	// for i := 1; i < 14; i++ {
	// 	pdf.AddPage()
	// 	//添加图片
	// 	pdf.Image("out2.jpg", 0, 0, &gopdf.Rect{W: 25.4, H: 14.7})
	// }

	//save to file
	pdf.WritePdf("sample.pdf")

}
