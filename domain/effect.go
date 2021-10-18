package domain

type EffectMeasure struct {
	value string
}

func (em EffectMeasure) String() string {
	return em.value
}

type EffectType struct {
	value int
}

type Effect struct {
	name     string
	positive bool
	eType    EffectType
	baseCost float64
	measure  EffectMeasure
}

func (e Effect) Name() string {
	return e.name
}

func (e Effect) HideEffectDetails() Effect {
	return Effect{
		name:     "Unknown Effect",
		positive: e.positive,
		eType:    e.eType,
		baseCost: e.baseCost,
		measure:  e.measure,
	}
}

func (e Effect) IsUnknown() bool {
	return e.name == "Unknown Effect"
}

func (e Effect) IsCommon() bool {
	return e.eType == typeCommon
}

func (e Effect) IsDurationOnly() bool {
	return e.eType == typeDurationOnly
}

func (e Effect) IsMagnitudeOnly() bool {
	return e.eType == typeMagnitudeOnly
}

func (e Effect) IsImmediate() bool {
	return e.eType == typeImmediate
}

func createEffect(name string, eType EffectType, positive bool, baseCost float64, measure ...EffectMeasure) Effect {
	effectMeasure := plainMeasure
	if len(measure) == 1 {
		effectMeasure = measure[0]
	}

	return Effect{
		name:     name,
		positive: positive,
		eType:    eType,
		baseCost: baseCost,
		measure:  effectMeasure,
	}
}

var (
	typeCommon        = EffectType{value: 0}
	typeMagnitudeOnly = EffectType{value: 1}
	typeDurationOnly  = EffectType{value: 2}
	typeImmediate     = EffectType{value: 3}
	percentMeasure    = EffectMeasure{value: "%"}
	timeMeasure       = EffectMeasure{value: "secs"}
	plainMeasure      = EffectMeasure{value: "pts"}
	distanceMeasure   = EffectMeasure{value: "feet"}

	restoreIntelligenceEffect = createEffect("Restore Intelligence", typeCommon, true, 38)
	resistPoisonEffect        = createEffect("Resist Poison", typeCommon, true, 0.5, percentMeasure)
	lightEffect               = createEffect("Light", typeCommon, true, 0.051, distanceMeasure)
	damageFatigueEffect       = createEffect("Damage Fatigue", typeCommon, false, 4.4)
	restoreFatigueEffect      = createEffect("Restore Fatigue", typeCommon, true, 2)
	restoreHealthEffect       = createEffect("Restore Health", typeCommon, true, 10)
	damageMagickaEffect       = createEffect("Damage Magicka", typeCommon, false, 2.5)
	invisibilityEffect        = createEffect("Invisibility", typeDurationOnly, true, 40)
	damageLuckEffect          = createEffect("Damage Luck", typeCommon, false, 100)
	fortifyWillpowerEffect    = createEffect("Fortify Willpower", typeCommon, true, 0.6)
	damageHealthEffect        = createEffect("Damage Health", typeCommon, false, 12)
	restoreAgilityEffect      = createEffect("Restore Agility", typeCommon, true, 38)
	fortifyStrengthEffect     = createEffect("Fortify Strength", typeCommon, true, 0.6)
	burdenEffect              = createEffect("Burden", typeCommon, false, 0.21)
	shieldEffect              = createEffect("Shield", typeCommon, true, 0.45, percentMeasure)
	fortifyAgilityEffect      = createEffect("Fortify Agility", typeCommon, true, 0.6)
	dispelEffect              = createEffect("Dispel", typeMagnitudeOnly, true, 3.6)
	resistDiseaseEffect       = createEffect("Resist Disease", typeCommon, true, 0.5, percentMeasure)
	silenceEffect             = createEffect("Silence", typeDurationOnly, false, 60)
	resistShockEffect         = createEffect("Resist Shock", typeCommon, true, 0.5, percentMeasure)
	fortifyEnduranceEffect    = createEffect("Fortify Endurance", typeCommon, true, 0.6)
	restoreMagickaEffect      = createEffect("Restore Magicka", typeCommon, true, 2.5)
	chameleonEffect           = createEffect("Chameleon", typeCommon, true, 0.63, percentMeasure)
	resistParalysisEffect     = createEffect("Resist Paralysis", typeCommon, true, 0.75, percentMeasure)
	fortifyHealthEffect       = createEffect("Fortify Health", typeCommon, true, 0.14)
	damageSpeedEffect         = createEffect("Damage Speed", typeCommon, false, 100)
	damagePersonalityEffect   = createEffect("Damage Personality", typeCommon, false, 100)
	damageEnduranceEffect     = createEffect("Damage Endurance", typeCommon, false, 100)
	detectLifeEffect          = createEffect("Detect Life", typeCommon, true, 0.08, distanceMeasure)
	damageAgilityEffect       = createEffect("Damage Agility", typeCommon, false, 100)
	damageStrengthEffect      = createEffect("Damage Strength", typeCommon, false, 100)
	damageIntelligenceEffect  = createEffect("Damage Intelligence", typeCommon, false, 100)
	shockDamageEffect         = createEffect("Shock Damage", typeCommon, false, 7.8)
	resistFireEffect          = createEffect("Resist Fire", typeCommon, true, 0.5)
	fireShieldEffect          = createEffect("Fire Shield", typeCommon, true, 0.95)
	restoreEnduranceEffect    = createEffect("Restore Endurance", typeCommon, true, 38)
	reflectSpellEffect        = createEffect("Reflect Spell", typeCommon, true, 0, percentMeasure)
	cureDiseaseEffect         = createEffect("Cure Disease", typeImmediate, true, 1400)
	paralyzeEffect            = createEffect("Paralyze", typeDurationOnly, false, 475)
	fortifyIntelligenceEffect = createEffect("Fortify Intelligence", typeCommon, true, 0.6)
	restorePersonalityEffect  = createEffect("Restore Personality", typeCommon, true, 38)
	resistFrostEffect         = createEffect("Resist Frost", typeCommon, true, 0.5, percentMeasure)
	fortifyMagickaEffect      = createEffect("Fortify Magicka", typeCommon, true, 0.15)
	shockShieldEffect         = createEffect("Shock Shield", typeCommon, true, 0.95, percentMeasure)
	reflectDamageEffect       = createEffect("Reflect Damage", typeCommon, true, 2.5, percentMeasure)
	waterBreathingEffect      = createEffect("Water Breathing", typeDurationOnly, true, 14.5)
	restoreLuckEffect         = createEffect("Restore Luck", typeCommon, true, 38)
	frostDamageEffect         = createEffect("Frost Damage", typeCommon, false, 7.4)
	damageWillpowerEffect     = createEffect("Damage Willpower", typeCommon, false, 100)
	fireDamageEffect          = createEffect("Fire Damage", typeCommon, false, 7.5)
	featherEffect             = createEffect("Feather", typeCommon, true, 0.01)
	fortifyFatigueEffect      = createEffect("Fortify Fatigue", typeCommon, true, 0.04)
	frostShieldEffect         = createEffect("Frost Shield", typeCommon, true, 0.95, percentMeasure)
	restoreSpeedEffect        = createEffect("Restore Speed", typeCommon, true, 38)
	curePoisonEffect          = createEffect("Cure Poison", typeImmediate, true, 600)
	waterWalkingEffect        = createEffect("Water Walking", typeCommon, true, 13)
	fortifyLuckEffect         = createEffect("Fortify Luck", typeCommon, true, 0.6)
	fortifyPersonalityEffect  = createEffect("Fortify Personality", typeCommon, true, 0.6)
	cureParalysisEffect       = createEffect("Cure Paralysis", typeImmediate, true, 500)
	restoreWillpowerEffect    = createEffect("Restore Willpower", typeCommon, true, 38)
	restoreStrengthEffect     = createEffect("Restore Strength", typeCommon, true, 38)
	fortifySpeedEffect        = createEffect("Fortify Speed", typeCommon, true, 0.6)
	nightEyeEffect            = createEffect("Night-Eye", typeDurationOnly, true, 22)
)
