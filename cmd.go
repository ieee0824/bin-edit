package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

func read(name string) []byte {
	bin, _ := ioutil.ReadFile(name)
	return bin
}

func main() {
	in := flag.String("i", "none", "input file")
	out := flag.String("o", "out.png", "output file")
	flag.Parse()
	if *in == "none" {
		return
	}
	bin := read(*in)
	binLength := len(bin)
	width := 384
	height := len(bin) / width
	rect := image.Rect(0, 0, width, height)

	img := image.NewRGBA(rect)

	p := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var c color.RGBA
			for addr := p; addr < p+3 && addr < binLength; addr++ {
				if addr%3 == 0 {
					c.R = uint8(bin[addr])
				} else if addr%3 == 1 {
					c.G = uint8(bin[addr])
				} else if addr%3 == 2 {
					c.B = uint8(bin[addr])
				}
				c.A = 0xff
			}
			p += 3
			img.Set(x, y, c)
		}
	}
	f, _ := os.OpenFile(*out, os.O_CREATE|os.O_WRONLY, 0666)
	png.Encode(f, img)
}
