package alchemist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/alchemist"
)

var equipmentLevels = []struct {
	name        string
	level       alchemist.EquipmentLevel
	rawStrength float64
	strength    float64
}{
	{name: "Novice", level: alchemist.EquipmentNovice, strength: 2.5},
	{name: "Apprentice", level: alchemist.EquipmentApprentice, strength: 6.25},
	{name: "Journeyman", level: alchemist.EquipmentJourneyman, strength: 12.5},
	{name: "Expert", level: alchemist.EquipmentExpert, strength: 18.75},
	{name: "Master", level: alchemist.EquipmentMaster, strength: 25},
}

func TestMortar_Strength(t *testing.T) {
	for i := 0; i < len(equipmentLevels); i++ {
		equipment := equipmentLevels[i]
		t.Run(equipment.name, func(sub *testing.T) {
			sub.Parallel()
			mortar := alchemist.NewMortar(equipment.level)
			assert.Equal(sub, equipment.strength, mortar.Strength())
		})
	}
}
