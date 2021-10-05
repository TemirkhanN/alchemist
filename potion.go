package main

import (
	"strconv"
	"strings"
)

type Potion struct {
	name    string
	effects []Effect
}

func (p Potion) Description() string {
	var description strings.Builder

	description.WriteString("Name: " + p.name + "\n")
	description.WriteString("Effects: \n")
	for _, effect := range p.effects {
		description.WriteString("    " + effect.name + " (" + effect.description + ")")
		description.WriteString("power(" + strconv.Itoa(int(effect.power)) + ")\n")
	}

	return description.String()
}
