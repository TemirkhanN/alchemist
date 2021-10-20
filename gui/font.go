package gui

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	"github.com/faiface/pixel/text"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

type Font struct {
	atlas *text.Atlas
}

func LoadFont(fontName string, ttfPath string, fontSize int) Font {
	ttfPath, err := filepath.Abs(ttfPath)
	if err != nil {
		panic(err)
	}

	fontFile := flag.String(fontName, ttfPath, "Lorem ipsum dolor")

	fontBytes, err := ioutil.ReadFile(*fontFile)
	if err != nil {
		panic(err)
	}

	fontOpts := &truetype.Options{
		Size:              float64(fontSize),
		DPI:               72,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	return Font{atlas: text.NewAtlas(truetype.NewFace(font, fontOpts), text.ASCII)}
}
