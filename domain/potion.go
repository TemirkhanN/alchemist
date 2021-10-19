package domain

import (
	"fmt"
	"strings"
)

type Potion struct {
	name    string
	effects []PotionEffect
}

type PotionEffect struct {
	magnitude float64
	duration  float64
	Effect
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

func (pe PotionEffect) Description() string {
	if pe.isImmediate() {
		return pe.Name()
	}

	if pe.isDurationOnly() {
		return fmt.Sprintf("%s for %d %s", pe.Name(), int(pe.duration), timeMeasure)
	}

	if pe.isMagnitudeOnly() {
		return fmt.Sprintf("%s %d %s", pe.Name(), int(pe.magnitude), pe.measure)
	}

	return fmt.Sprintf("%s %d %s for %d %s", pe.Name(), int(pe.magnitude), pe.measure, int(pe.duration), timeMeasure)
}
