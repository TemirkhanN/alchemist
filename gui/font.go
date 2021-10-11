package gui

import (
	"flag"
	"image"
	"io/ioutil"

	"github.com/golang/freetype"
	"golang.org/x/image/colornames"
)

// todo figure out how to use it instead of basicfont.Face7x13.
func RegisterFont(name string, ttfPath string, dpi float64, fontSize float64) *freetype.Context {
	fontFile := flag.String(name, ttfPath, "Lorem ipsum dolor")
	fontBytes, _ := ioutil.ReadFile(*fontFile)
	convertFontSize := flag.Float64("size", fontSize, "font size in points")
	convertedDpi := flag.Float64("dpi", dpi, "screen resolution in Dots Per Inch")

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	c := freetype.NewContext()
	c.SetDPI(*convertedDpi)
	c.SetFont(f)
	c.SetFontSize(*convertFontSize)
	c.SetSrc(image.NewUniform(colornames.Sienna))

	return c
}
