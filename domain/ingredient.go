package domain

type EffectType string

const (
	Positive EffectType = "+"
	Negative EffectType = "-"
)

type Effect struct {
	name      string
	eType     EffectType
	power     int
	increased bool
}

type Ingredient struct {
	name    string
	effects []Effect
}

func (i *Ingredient) Name() string {
	return i.name
}

func (e *Effect) Name() string {
	return e.name
}
