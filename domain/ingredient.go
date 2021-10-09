package domain

type EffectType string

const (
	Positive EffectType = "+"
	Negative EffectType = "-"
)

type Effect struct {
	name      string
	eType     EffectType
	power     int16
	increased bool
}

type Ingredient struct {
	name    string
	effects []Effect
}

func (i *Ingredient) Name() string {
	return i.name
}
