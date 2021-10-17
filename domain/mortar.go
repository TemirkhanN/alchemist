package domain

type EquipmentLevel struct {
	value uint8
}

type Mortar struct {
	level EquipmentLevel
}

func NewMortar(level EquipmentLevel) *Mortar {
	return &Mortar{level: level}
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

var (
	EquipmentNovice     = EquipmentLevel{value: 0}
	EquipmentApprentice = EquipmentLevel{value: 1}
	EquipmentJourneyman = EquipmentLevel{value: 2}
	EquipmentExpert     = EquipmentLevel{value: 3}
	EquipmentMaster     = EquipmentLevel{value: 4}
)
