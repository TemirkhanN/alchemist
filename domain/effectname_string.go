// Code generated by "stringer -linecomment -type=EffectName"; DO NOT EDIT.

package domain

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RestoreIntelligence-1]
	_ = x[ResistPoison-2]
	_ = x[Light-3]
	_ = x[DamageFatigue-4]
	_ = x[RestoreFatigue-5]
	_ = x[RestoreHealth-6]
	_ = x[DamageMagicka-7]
	_ = x[Invisibility-8]
	_ = x[DamageLuck-9]
	_ = x[FortifyWillpower-10]
	_ = x[DamageHealth-11]
	_ = x[RestoreAgility-12]
	_ = x[FortifyStrength-13]
	_ = x[Burden-14]
	_ = x[Shield-15]
	_ = x[FortifyAgility-16]
	_ = x[Dispel-17]
	_ = x[ResistDisease-18]
	_ = x[Silence-19]
	_ = x[ResistShock-20]
	_ = x[FortifyEndurance-21]
	_ = x[RestoreMagicka-22]
	_ = x[Chameleon-23]
	_ = x[ResistParalysis-24]
	_ = x[FortifyHealth-25]
	_ = x[DamageSpeed-26]
	_ = x[DamagePersonality-27]
	_ = x[DamageEndurance-28]
	_ = x[DetectLife-29]
	_ = x[DamageAgility-30]
	_ = x[DamageStrength-31]
	_ = x[DamageIntelligence-32]
	_ = x[ShockDamage-33]
	_ = x[ResistFire-34]
	_ = x[FireShield-35]
	_ = x[RestoreEndurance-36]
	_ = x[ReflectSpell-37]
	_ = x[CureDisease-38]
	_ = x[Paralyze-39]
	_ = x[FortifyIntelligence-40]
	_ = x[RestorePersonality-41]
	_ = x[ResistFrost-42]
	_ = x[FortifyMagicka-43]
	_ = x[ShockShield-44]
	_ = x[ReflectDamage-45]
	_ = x[WaterBreathing-46]
	_ = x[RestoreLuck-47]
	_ = x[FrostDamage-48]
	_ = x[DamageWillpower-49]
	_ = x[FireDamage-50]
	_ = x[Feather-51]
	_ = x[FortifyFatigue-52]
	_ = x[FrostShield-53]
	_ = x[RestoreSpeed-54]
	_ = x[CurePoison-55]
	_ = x[WaterWalking-56]
	_ = x[FortifyLuck-57]
	_ = x[FortifyPersonality-58]
	_ = x[CureParalysis-59]
	_ = x[RestoreWillpower-60]
	_ = x[RestoreStrength-61]
	_ = x[FortifySpeed-62]
	_ = x[NightEye-63]
}

const _EffectName_name = "Restore IntelligenceResist PoisonLightDamage FatigueRestore FatigueRestore HealthDamage MagickaInvisibilityDamage LuckFortify WillpowerDamage HealthRestore AgilityFortify StrengthBurdenShieldFortify AgilityDispelResist DiseaseSilenceResist ShockFortify EnduranceRestore MagickaChameleonResist ParalysisFortify HealthDamage SpeedDamage PersonalityDamage EnduranceDetect LifeDamage AgilityDamage StrengthDamage IntelligenceShock DamageResist FireFire ShieldRestore EnduranceReflect SpellCure DiseaseParalyzeFortify IntelligenceRestore PersonalityResist FrostFortify MagickaShock ShieldReflect DamageWater BreathingRestore LuckFrost DamageDamage WillpowerFire DamageFeatherFortify FatigueFrost ShieldRestore SpeedCure PoisonWater WalkingFortify LuckFortify PersonalityCure ParalysisRestore WillpowerRestore StrengthFortify SpeedNight-Eye"

var _EffectName_index = [...]uint16{0, 20, 33, 38, 52, 67, 81, 95, 107, 118, 135, 148, 163, 179, 185, 191, 206, 212, 226, 233, 245, 262, 277, 286, 302, 316, 328, 346, 362, 373, 387, 402, 421, 433, 444, 455, 472, 485, 497, 505, 525, 544, 556, 571, 583, 597, 612, 624, 636, 652, 663, 670, 685, 697, 710, 721, 734, 746, 765, 779, 796, 812, 825, 834}

func (i EffectName) String() string {
	i -= 1
	if i < 0 || i >= EffectName(len(_EffectName_index)-1) {
		return "EffectName(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _EffectName_name[_EffectName_index[i]:_EffectName_index[i+1]]
}
