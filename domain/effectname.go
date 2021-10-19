package domain

//go:generate stringer -linecomment -type=EffectName
type EffectName int

const (
	_                   = EffectName(iota)
	RestoreIntelligence // Restore Intelligence
	ResistPoison        // Resist Poison
	Light               // Light
	DamageFatigue       // Damage Fatigue
	RestoreFatigue      // Restore Fatigue
	RestoreHealth       // Restore Health
	DamageMagicka       // Damage Magicka
	Invisibility        // Invisibility
	DamageLuck          // Damage Luck
	FortifyWillpower    // Fortify Willpower
	DamageHealth        // Damage Health
	RestoreAgility      // Restore Agility
	FortifyStrength     // Fortify Strength
	Burden              // Burden
	Shield              // Shield
	FortifyAgility      // Fortify Agility
	Dispel              // Dispel
	ResistDisease       // Resist Disease
	Silence             // Silence
	ResistShock         // Resist Shock
	FortifyEndurance    // Fortify Endurance
	RestoreMagicka      // Restore Magicka
	Chameleon           // Chameleon
	ResistParalysis     // Resist Paralysis
	FortifyHealth       // Fortify Health
	DamageSpeed         // Damage Speed
	DamagePersonality   // Damage Personality
	DamageEndurance     // Damage Endurance
	DetectLife          // Detect Life
	DamageAgility       // Damage Agility
	DamageStrength      // Damage Strength
	DamageIntelligence  // Damage Intelligence
	ShockDamage         // Shock Damage
	ResistFire          // Resist Fire
	FireShield          // Fire Shield
	RestoreEndurance    // Restore Endurance
	ReflectSpell        // Reflect Spell
	CureDisease         // Cure Disease
	Paralyze            // Paralyze
	FortifyIntelligence // Fortify Intelligence
	RestorePersonality  // Restore Personality
	ResistFrost         // Resist Frost
	FortifyMagicka      // Fortify Magicka
	ShockShield         // Shock Shield
	ReflectDamage       // Reflect Damage
	WaterBreathing      // Water Breathing
	RestoreLuck         // Restore Luck
	FrostDamage         // Frost Damage
	DamageWillpower     // Damage Willpower
	FireDamage          // Fire Damage
	Feather             // Feather
	FortifyFatigue      // Fortify Fatigue
	FrostShield         // Frost Shield
	RestoreSpeed        // Restore Speed
	CurePoison          // Cure Poison
	WaterWalking        // Water Walking
	FortifyLuck         // Fortify Luck
	FortifyPersonality  // Fortify Personality
	CureParalysis       // Cure Paralysis
	RestoreWillpower    // Restore Willpower
	RestoreStrength     // Restore Strength
	FortifySpeed        // Fortify Speed
	NightEye            // Night-Eye
)
