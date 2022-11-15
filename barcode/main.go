package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func main() {
	// 创建一个code128编码的 BarcodeIntCS
	cs, _ := code128.Encode("tdx123456")
	// 创建一个要输出数据的文件
	file, _ := os.Create("tiaoxingma2.jpg")
	defer file.Close()

	// 设置图片像素大小
	qrCode, _ := barcode.Scale(cs, 400, 45)
	// 将code128的条形码编码为png图片
	png.Encode(file, qrCode)
}
