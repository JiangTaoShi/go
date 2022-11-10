package main

import (
	"fmt"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	pdf "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	// The default import conf uses the special pos:full argument
	// which overrides all other import conf parms.
	imp := pdf.DefaultImportConfig()
	// if _, err := api.Import("f:A4, dpi:300, p:c", pdf.POINTS); err != nil {
	// 	fmt.Println(err)
	// }
	outFile, _ := filepath.Abs("test.pdf")
	var imgFiles []string

	imgPath, _ := filepath.Abs("out2.jpg")
	imgFiles = append(imgFiles, imgPath)
	if err := api.ImportImagesFile(imgFiles, outFile, imp, nil); err != nil {
		fmt.Println(outFile, err)
	}
	if err := api.ValidateFile(outFile, nil); err != nil {
		fmt.Println(err)
	}
}
