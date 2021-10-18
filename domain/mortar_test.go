package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TemirkhanN/alchemist/domain"
)

var equipmentLevels = []struct {
	name        string
	level       domain.EquipmentLevel
	rawStrength float64
	strength    float64
}{
	{name: "Novice", level: domain.EquipmentNovice, rawStrength: 0.1, strength: 2.5},
	{name: "Apprentice", level: domain.EquipmentApprentice, rawStrength: 0.25, strength: 6.25},
	{name: "Journeyman", level: domain.EquipmentJourneyman, rawStrength: 0.5, strength: 12.5},
	{name: "Expert", level: domain.EquipmentExpert, rawStrength: 0.75, strength: 18.75},
	{name: "Master", level: domain.EquipmentMaster, rawStrength: 1, strength: 25},
}

func TestMortar_Strength(t *testing.T) {
	for i := 0; i < len(equipmentLevels); i++ {
		equipment := equipmentLevels[i]
		t.Run(equipment.name, func(sub *testing.T) {
			sub.Parallel()
			mortar := domain.NewMortar(equipment.level)
			assert.Equal(sub, equipment.strength, mortar.Strength())
		})
	}
}
