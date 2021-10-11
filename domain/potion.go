package domain

type Potion struct {
	name    string
	effects []PotionEffect
}

type PotionEffect struct {
	magnitude float64
	duration  float64
	Effect
}
