package gui

import (
	"embed"
	"errors"
	"fmt"
	"image"
	"strings"

	errors2 "github.com/pkg/errors"

	"github.com/TemirkhanN/alchemist/pkg/gui/graphics"
)

type Assets struct {
	filesystem  embed.FS
	sprites     map[string]string
	cache       map[string]*graphics.Sprite
	initialized bool
}

// RegisterAssets Loads assets from given directory and filesystem for further usage
// currently supports only png files which are recognized as sprites
// dirPath is a relative path to directory from fs root. It is used as prefix for further calls.
// Example:
//     fs ~ assetsDir/
//          ├── logo.png
//          ├── character/
//              ├── male.png
//              ├── female.png
//              └── alien.png
//
//     gameAssets := new(Assets)
//     gameAssets.RegisterAssets("assetsDir", fs)
//
//     maleSprite := gameAssets.GetSprite("character/male")
//     femaleSprite := gameAssets.GetSprite("character/female")
//     alienSprite := gameAssets.GetSprite("character/alien")
//     logoSprite := gameAssets.GetSprite("logo")
func (assets *Assets) RegisterAssets(dirPath string, fs embed.FS) error {
	if !assets.initialized {
		assets.initialized = true
		assets.filesystem = fs
		assets.sprites = make(map[string]string)
		assets.cache = make(map[string]*graphics.Sprite)
	}

	if assets.filesystem != fs {
		return errors.New("multiple attempts to register assets")
	}

	dirEntries, err := fs.ReadDir(strings.TrimRight(dirPath, "/"))
	if err != nil {
		return errors2.Wrap(err, "Couldn't register assets")
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			_ = assets.RegisterAssets(fmt.Sprintf("%s/%s", dirPath, dirEntry.Name()), fs)

			continue
		}

		if strings.HasSuffix(dirEntry.Name(), ".png") {
			assets.sprites[dirEntry.Name()] = fmt.Sprintf("%s/%s", dirPath, dirEntry.Name())
		}
	}

	return nil
}

// GetSprite Returns sprite for the given spriteName.
// spriteName shall be passed without extension. See example in RegisterAssets.
func (assets Assets) GetSprite(spriteName string) *graphics.Sprite {
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

		assets.cache[spriteName] = graphics.NewSprite(img)
	}

	return assets.cache[spriteName]
}
