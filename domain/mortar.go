package domain

type EquipmentLevel int

const (
	EquipmentNovice EquipmentLevel = iota
	EquipmentApprentice
	EquipmentJourneyman
	EquipmentExpert
	EquipmentMaster
)

type Mortar struct {
	level EquipmentLevel
}

func NewNoviceMortar() *Mortar {
	return &Mortar{level: EquipmentNovice}
}

func (m *Mortar) Strength() float64 {
	return m.getEquipmentStrength() * 25
}

func (m *Mortar) getEquipmentStrength() float64 {
	switch m.level {
	case EquipmentNovice:
		return 0.1
	case EquipmentApprentice:
		return 0.25
	case EquipmentJourneyman:
		return 0.5
	case EquipmentExpert:
		return 0.75
	case EquipmentMaster:
		return 1
	default:
		panic("Unknown equipment level")
	}
}
