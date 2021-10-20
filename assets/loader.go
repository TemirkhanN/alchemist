package assets

import (
	"embed"
	"log"

	"github.com/TemirkhanN/alchemist/pkg/gui"
)

//go:embed sprites
var SpritesFs embed.FS

var TESAssets = func() *gui.Assets {
	loadedAssets := new(gui.Assets)
	// todo shall filesystem be passed by reference or not?
	err := loadedAssets.RegisterAssets("sprites", &SpritesFs)
	if err != nil {
		log.Fatal(err)
	}

	return loadedAssets
}()
