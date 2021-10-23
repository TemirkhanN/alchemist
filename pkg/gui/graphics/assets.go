package graphics

import (
	"embed"
	"errors"
	"fmt"
	"image"
	"strings"

	"github.com/faiface/pixel"
	errors2 "github.com/pkg/errors"
)

type Assets struct {
	filesystem  embed.FS
	sprites     map[string]string
	cache       map[string]*Sprite
	initialized bool
}

func (assets *Assets) RegisterAssets(directory string, fs embed.FS) error {
	if !assets.initialized {
		assets.initialized = true
		assets.filesystem = fs
		assets.sprites = make(map[string]string)
		assets.cache = make(map[string]*Sprite)
	}

	if assets.filesystem != fs {
		return errors.New("multiple attempts to register assets")
	}

	dirEntries, err := fs.ReadDir(strings.TrimRight(directory, "/"))
	if err != nil {
		return errors2.Wrap(err, "Couldn't register assets")
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
