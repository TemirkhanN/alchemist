package GUI

import (
	"embed"
	"errors"
	"fmt"
	"github.com/faiface/pixel"
	"image"
	"strings"
)

type Assets struct {
	filesystem *embed.FS
	sprites    map[string]string
	cache      map[string]*Sprite
}

func (assets *Assets) RegisterAssets(directory string, fs *embed.FS) error {
	if assets.filesystem == nil {
		assets.filesystem = fs
		assets.sprites = make(map[string]string)
		assets.cache = make(map[string]*Sprite)
	}

	if assets.filesystem != fs {
		return errors.New("multiple attempts to register assets")
	}

	directory = strings.TrimRight(directory, "/")
	dirEntries, err := fs.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			_ = assets.RegisterAssets(fmt.Sprintf("%s/%s", directory, dirEntry.Name()), fs)
			continue
		}

		if strings.HasSuffix(dirEntry.Name(), ".png") {
			assets.sprites[dirEntry.Name()] = fmt.Sprintf("%s/%s", directory, dirEntry.Name())
		}
	}

	return nil
}

func (assets Assets) GetSprite(spriteName string) *Sprite {
	if assets.cache == nil {
		panic("attempt to load sprite while there are not any assets registered")
	}

	cachedSprite := assets.cache[spriteName]
	if cachedSprite == nil {
		key := fmt.Sprintf("%s.%s", spriteName, "png")
		spritePath, spriteExists := assets.sprites[key]
		if !spriteExists {
			panic(fmt.Sprintf("sprite %s does not exist", spriteName))
		}

		file, err := assets.filesystem.Open(spritePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			panic(err)
		}

		pic := pixel.PictureDataFromImage(img)

		assets.cache[spriteName] = &Sprite{src: pixel.NewSprite(pic, pic.Bounds())}
	}

	return assets.cache[spriteName]
}
