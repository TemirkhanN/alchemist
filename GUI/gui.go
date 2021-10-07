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
	pathMap    map[string]string
}

func (assets *Assets) RegisterAssets(directory string, fs *embed.FS) error {
	if assets.filesystem != nil {
		return errors.New("multiple attempts to register assets")
	}

	assets.filesystem = fs
	assets.pathMap = make(map[string]string)

	directory = strings.TrimRight(directory, "/")
	dirEntries, err := fs.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}
		assets.pathMap[dirEntry.Name()] = fmt.Sprintf("%s/%s", directory, dirEntry.Name())
	}

	return nil
}

func (assets Assets) GetSprite(spriteName string) *pixel.Sprite {
	key := fmt.Sprintf("%s.%s", spriteName, "png")
	spritePath, spriteExists := assets.pathMap[key]
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

	return pixel.NewSprite(pic, pic.Bounds())
}
