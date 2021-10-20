package alchemist

import (
	"fmt"
	"strings"

	"github.com/TemirkhanN/alchemist/pkg/alchemy/ingredient"
)

type Potion struct {
	name    string
	effects []PotionEffect
}

type PotionEffect struct {
	magnitude float64
	duration  float64
	effect    ingredient.Effect
}

func (p Potion) Effects() []PotionEffect {
	return p.effects
}

func (p Potion) Description() string {
	descriptionBuilder := strings.Builder{}

	for _, effect := range p.effects {
		descriptionBuilder.WriteString(effect.Description())
		descriptionBuilder.WriteString("\n")
	}

	return descriptionBuilder.String()
}

func (pe PotionEffect) Name() string {
	return pe.effect.Name()
}

func (pe PotionEffect) Description() string {
	effect := pe.effect
	if effect.IsImmediate() {
		return effect.Name()
	}

	if pe.effect.IsDurationOnly() {
		return fmt.Sprintf("%s for %d secs", effect.Name(), int(pe.duration))
	}

	if pe.effect.IsMagnitudeOnly() {
		return fmt.Sprintf("%s %d %s", effect.Name(), int(pe.magnitude), effect.Measure())
	}

	return fmt.Sprintf("%s %d %s for %d secs", effect.Name(), int(pe.magnitude), effect.Measure(), int(pe.duration))
}
