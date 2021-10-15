package gui

import (
	"flag"
	"io/ioutil"
	"path/filepath"

	"github.com/faiface/pixel/text"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func createAtlas(fontName string, ttfPath string) *text.Atlas {
	ttfPath, err := filepath.Abs(ttfPath)
	if err != nil {
		panic(err)
	}

	fontFile := flag.String(fontName, ttfPath, "Lorem ipsum dolor")

	fontBytes, err := ioutil.ReadFile(*fontFile)
	if err != nil {
		panic(err)
	}

	oblivionFontOpts := &truetype.Options{
		Size:              28,
		DPI:               72,
		Hinting:           0,
		GlyphCacheEntries: 0,
		SubPixelsX:        0,
		SubPixelsY:        0,
	}

	oblivionFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	return text.NewAtlas(truetype.NewFace(oblivionFont, oblivionFontOpts), text.ASCII)
}
