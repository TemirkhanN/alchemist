package ingredient

type Ingredient struct {
	name    string
	effects []Effect
}

func (i Ingredient) Name() string {
	return i.name
}

func (i Ingredient) Effects() []Effect {
	return i.effects
}
