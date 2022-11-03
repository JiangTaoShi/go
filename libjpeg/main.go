package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pixiv/go-libjpeg/jpeg"
)

func main() {

	fmt.Println("test")

	// Decoding JPEG into image.Image
	io, err := os.Open("01.jpg")
	if err != nil {
		log.Fatal(err)
	}
	img, err := jpeg.Decode(io, &jpeg.DecoderOptions{})
	if err != nil {
		log.Fatalf("Decode returns error: %v\n", err)
	}

	// Encode JPEG
	f, err := os.Create("out.jpg")
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	if err := jpeg.Encode(w, img, &jpeg.EncoderOptions{Quality: 90}); err != nil {
		log.Printf("Encode returns error: %v\n", err)
		return
	}
	w.Flush()
	f.Close()
}
