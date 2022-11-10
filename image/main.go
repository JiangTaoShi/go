package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	exif2 "github.com/dsoprea/go-exif/v2"
	exifcommon "github.com/dsoprea/go-exif/v2/common"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
)

func main() {

	absPath, _ := filepath.Abs("01.jpg")
	SetExifData(absPath)

}

func SetExifData(filepath string) error {
	jmp := jpegstructure.NewJpegMediaParser()

	intfc, err := jmp.ParseFile(filepath)
	log.Println(err)

	os.Remove(filepath)

	sl := intfc.(*jpegstructure.SegmentList)

	// Make sure we don't start out with EXIF data.
	wasDropped, err := sl.DropExif()
	log.Println(err)

	if wasDropped != true {
		fmt.Printf("Expected the EXIF segment to be dropped, but it wasn't.")
	}

	im := exif2.NewIfdMapping()

	err = exif2.LoadStandardIfds(im)
	log.Println(err)

	ti := exif2.NewTagIndex()
	rootIb := exif2.NewIfdBuilder(im, ti, exifcommon.IfdPathStandard, exifcommon.EncodeDefaultByteOrder)

	err = rootIb.AddStandardWithName("XResolution", []exifcommon.Rational{{Numerator: 300, Denominator: 1}})
	log.Println(err)

	err = rootIb.AddStandardWithName("YResolution", []exifcommon.Rational{{Numerator: 300, Denominator: 1}})
	log.Println(err)

	err = sl.SetExif(rootIb)
	log.Println(err)

	b := new(bytes.Buffer)

	err = sl.Write(b)
	log.Println(err)

	if err := ioutil.WriteFile(filepath, b.Bytes(), 0644); err != nil {
		fmt.Printf("write file err: %v", err)
	}
	return nil
}
