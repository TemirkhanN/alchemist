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

func (e Effect) hideEffectDetails() Effect {
	return Effect{
		name:     "Unknown Effect",
		positive: e.positive,
		eType:    e.eType,
		baseCost: e.baseCost,
		measure:  e.measure,
	}
}

func (e Effect) isUnknown() bool {
	return e.name == "Unknown Effect"
}

func (e Effect) isDurationOnly() bool {
	return e.eType == typeDurationOnly
}

func (e Effect) isMagnitudeOnly() bool {
	return e.eType == typeMagnitudeOnly
}

func (e Effect) isImmediate() bool {
	return e.eType == typeImmediate
}

func createEffect(name EffectName, eType EffectType, positive bool, baseCost float64, measure ...EffectMeasure) Effect {
	effectMeasure := plainMeasure
	if len(measure) == 1 {
		effectMeasure = measure[0]
	}

	return Effect{
		name:     name.String(),
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

	restoreIntelligenceEffect = createEffect(RestoreIntelligence, typeCommon, true, 38)
	resistPoisonEffect        = createEffect(ResistPoison, typeCommon, true, 0.5, percentMeasure)
	lightEffect               = createEffect(Light, typeCommon, true, 0.051, distanceMeasure)
	damageFatigueEffect       = createEffect(DamageFatigue, typeCommon, false, 4.4)
	restoreFatigueEffect      = createEffect(RestoreFatigue, typeCommon, true, 2)
	restoreHealthEffect       = createEffect(RestoreHealth, typeCommon, true, 10)
	damageMagickaEffect       = createEffect(DamageMagicka, typeCommon, false, 2.5)
	invisibilityEffect        = createEffect(Invisibility, typeDurationOnly, true, 40)
	damageLuckEffect          = createEffect(DamageLuck, typeCommon, false, 100)
	fortifyWillpowerEffect    = createEffect(FortifyWillpower, typeCommon, true, 0.6)
	damageHealthEffect        = createEffect(DamageHealth, typeCommon, false, 12)
	restoreAgilityEffect      = createEffect(RestoreAgility, typeCommon, true, 38)
	fortifyStrengthEffect     = createEffect(FortifyStrength, typeCommon, true, 0.6)
	burdenEffect              = createEffect(Burden, typeCommon, false, 0.21)
	shieldEffect              = createEffect(Shield, typeCommon, true, 0.45, percentMeasure)
	fortifyAgilityEffect      = createEffect(FortifyAgility, typeCommon, true, 0.6)
	dispelEffect              = createEffect(Dispel, typeMagnitudeOnly, true, 3.6)
	resistDiseaseEffect       = createEffect(ResistDisease, typeCommon, true, 0.5, percentMeasure)
	silenceEffect             = createEffect(Silence, typeDurationOnly, false, 60)
	resistShockEffect         = createEffect(ResistShock, typeCommon, true, 0.5, percentMeasure)
	fortifyEnduranceEffect    = createEffect(FortifyEndurance, typeCommon, true, 0.6)
	restoreMagickaEffect      = createEffect(RestoreMagicka, typeCommon, true, 2.5)
	chameleonEffect           = createEffect(Chameleon, typeCommon, true, 0.63, percentMeasure)
	resistParalysisEffect     = createEffect(ResistParalysis, typeCommon, true, 0.75, percentMeasure)
	fortifyHealthEffect       = createEffect(FortifyHealth, typeCommon, true, 0.14)
	damageSpeedEffect         = createEffect(DamageSpeed, typeCommon, false, 100)
	damagePersonalityEffect   = createEffect(DamagePersonality, typeCommon, false, 100)
	damageEnduranceEffect     = createEffect(DamageEndurance, typeCommon, false, 100)
	detectLifeEffect          = createEffect(DetectLife, typeCommon, true, 0.08, distanceMeasure)
	damageAgilityEffect       = createEffect(DamageAgility, typeCommon, false, 100)
	damageStrengthEffect      = createEffect(DamageStrength, typeCommon, false, 100)
	damageIntelligenceEffect  = createEffect(DamageIntelligence, typeCommon, false, 100)
	shockDamageEffect         = createEffect(ShockDamage, typeCommon, false, 7.8)
	resistFireEffect          = createEffect(ResistFire, typeCommon, true, 0.5)
	fireShieldEffect          = createEffect(FireShield, typeCommon, true, 0.95)
	restoreEnduranceEffect    = createEffect(RestoreEndurance, typeCommon, true, 38)
	reflectSpellEffect        = createEffect(ReflectSpell, typeCommon, true, 0, percentMeasure)
	cureDiseaseEffect         = createEffect(CureDisease, typeImmediate, true, 1400)
	paralyzeEffect            = createEffect(Paralyze, typeDurationOnly, false, 475)
	fortifyIntelligenceEffect = createEffect(FortifyIntelligence, typeCommon, true, 0.6)
	restorePersonalityEffect  = createEffect(RestorePersonality, typeCommon, true, 38)
	resistFrostEffect         = createEffect(ResistFrost, typeCommon, true, 0.5, percentMeasure)
	fortifyMagickaEffect      = createEffect(FortifyMagicka, typeCommon, true, 0.15)
	shockShieldEffect         = createEffect(ShockShield, typeCommon, true, 0.95, percentMeasure)
	reflectDamageEffect       = createEffect(ReflectDamage, typeCommon, true, 2.5, percentMeasure)
	waterBreathingEffect      = createEffect(WaterBreathing, typeDurationOnly, true, 14.5)
	restoreLuckEffect         = createEffect(RestoreLuck, typeCommon, true, 38)
	frostDamageEffect         = createEffect(FrostDamage, typeCommon, false, 7.4)
	damageWillpowerEffect     = createEffect(DamageWillpower, typeCommon, false, 100)
	fireDamageEffect          = createEffect(FireDamage, typeCommon, false, 7.5)
	featherEffect             = createEffect(Feather, typeCommon, true, 0.01)
	fortifyFatigueEffect      = createEffect(FortifyFatigue, typeCommon, true, 0.04)
	frostShieldEffect         = createEffect(FrostShield, typeCommon, true, 0.95, percentMeasure)
	restoreSpeedEffect        = createEffect(RestoreSpeed, typeCommon, true, 38)
	curePoisonEffect          = createEffect(CurePoison, typeImmediate, true, 600)
	waterWalkingEffect        = createEffect(WaterWalking, typeCommon, true, 13)
	fortifyLuckEffect         = createEffect(FortifyLuck, typeCommon, true, 0.6)
	fortifyPersonalityEffect  = createEffect(FortifyPersonality, typeCommon, true, 0.6)
	cureParalysisEffect       = createEffect(CureParalysis, typeImmediate, true, 500)
	restoreWillpowerEffect    = createEffect(RestoreWillpower, typeCommon, true, 38)
	restoreStrengthEffect     = createEffect(RestoreStrength, typeCommon, true, 38)
	fortifySpeedEffect        = createEffect(FortifySpeed, typeCommon, true, 0.6)
	nightEyeEffect            = createEffect(NightEye, typeDurationOnly, true, 22)
)
