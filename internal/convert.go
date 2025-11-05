package internal

import (
	"fmt"
	"github.com/abrhoda/tdm/internal/foundry"
	"github.com/abrhoda/tdm/storage"
)

var rankIntToString = map[int]string{
	1: "untrained",
	2: "trained",
	3: "expert",
	4: "master",
	5: "legendary",
}

var proficiencyStringToType = map[string]string {
	"fortitude": "Saving Throw",
	"reflex": "Saving Throw",
	"will": "Saving Throw",
	"unarmed": "Attack",
	"simple": "Attack",
	"martial": "Attack",
	"advanced": "Attack",
	"unarmored": "Defense",
	"light": "Defense",
	"medium": "Defense",
	"heavy": "Defense",

	"spellcasting": "Spellcasting DC",

	// some of these aren't used but better to have for future.
	"alchemist": "Class DC",
	"animist": "Class DC",
	"barbarian": "Class DC",
	"bard": "Class DC",
	"champion": "Class DC",
	"cleric": "Class DC",
	"commander": "Class DC",
	"druid": "Class DC",
	"exempler": "Class DC",
	"fighter": "Class DC",
	"guardian": "Class DC",
	"gunslinger": "Class DC",
	"inventor": "Class DC",
	"investigator": "Class DC",
	"kineticist": "Class DC",
	"magus": "Class DC",
	"monk": "Class DC",
	"oracle": "Class DC",
	"psychic": "Class DC",
	"ranger": "Class DC",
	"rogue": "Class DC",
	"sorcerer": "Class DC",
	"summoner": "Class DC",
	"swashbuckler": "Class DC",
	"thaumaturge": "Class DC",
	"witch": "Class DC",
	"wizard": "Class DC",
}

var abilityShortNameToLongName = map[string]string{
	"str": "strength",
	"dex": "dexterity",
	"con": "constitution",
	"int": "intelligence",
	"wis": "wisdom",
	"cha": "charisma",
}

var senseMangledNameToLongName = map[string]string{
	"low-light-vision":   "Low Light Vision",
	"lowLightVision":     "Low Light Vision",
	"darkvision":         "Dark Vision",
	"tremorsense":        "Tremor Sense",
	"thoughtsense":       "Thought Sense",
	"scent":              "Scent",
	"spiritsense":        "Spirit Sense",
	"see-invisibility":   "See Invisibility",
	"greater-darkvision": "Greater Invisibility",
	"truesight":          "True Sight",
	"lifesense":          "Life Sense",
	"motion-sense":       "Motion Sense",
	"echolocation":       "Echo Location",
}

var allTraits = make(map[string]storage.Trait)

// first checks if trait exists in allTraits and uses that if found.
// if not, adds to the allTraits map and then uses it.
func convertTraits(traits []string) []storage.Trait {
	out := make([]storage.Trait, len(traits))
	for i, t := range traits {
		st, found := allTraits[t]
		if found {
			out[i] = st
		} else {
			st = storage.Trait{Name: t}
			allTraits[t] = st
			out[i] = st
		}
	}
	return out
}

func convertSenses(sMap map[string]foundry.Sense) []storage.Sense {
	s := make([]storage.Sense, len(sMap))
	for k, v := range sMap {
		elevate := false
		if v.Special != nil {
			val, ok := v.Special["ancestry"]
			if ok && !elevate {
				elevate = val
			}

			val, ok = v.Special["llv"]
			if ok && !elevate {
				elevate = val
			}

			val, ok = v.Special["second"]
			if ok && !elevate {
				elevate = val
			}
		}
		s = append(s, storage.Sense{Name: senseMangledNameToLongName[k], Acuity: v.Acuity, Range: v.Range, ElevateIfHasLowLightVision: elevate})
	}
	return s
}

func convertProficiencies(pMap map[string]map[string]int) []storage.Proficiency {
	ps := make([]storage.Proficiency, len(pMap))

	for k, v := range pMap {
		ps = append(ps, storage.Proficiency{Name: k, Rank: rankIntToString[v["rank"]]})
	}

	return ps
}

func validateFoundryAncestryFeature(af foundry.Feature) error {
	if af.System.ActionType.Value != "passive" {
		return fmt.Errorf("Expected ActionType to be passive. Was %s", af.System.ActionType.Value)
	}

	if af.System.Actions.Value != 0 {
		return fmt.Errorf("Expected Actions to be null/0.")
	}

	if af.System.Category != "ancestryfeature" {
		return fmt.Errorf("Expected Category to be 'ancestryfeature. Was %s'", af.System.Category)
	}

	if af.System.Level.Value > 1 {
		return fmt.Errorf("Expected Level to be 0 or 1.")
	}

	if len(af.System.Prerequisites.Value) != 0 {
		return fmt.Errorf("Expected prerequisites to be empty.")
	}

	if len(af.System.SubFeatures.KeyOptions) != 0 {
		return fmt.Errorf("Expected subFeature.KeyOptions to be empty.")
	}

	if len(af.System.SubFeatures.Proficiencies) != 0 {
		return fmt.Errorf("Expected subFeature.Proficiencies to be empty.")

	}

	if len(af.System.SubFeatures.SuppressedFeatures) != 0 {
		return fmt.Errorf("Expected subFeature.SuppressedFeatures to be empty.")

	}

	if len(af.System.Traits.OtherTags) != 0 {
		return fmt.Errorf("Expected len of `traits.otherTags` to be 0.")
	}

	return nil
}

func ConvertAncestryProperty(f foundry.Feature) (storage.AncestryProperty, error) {
	af := storage.AncestryProperty{}
	err := validateFoundryAncestryFeature(f)
	if err != nil {
		return af, err
	}

	prereqs := make([]string, len(f.System.Prerequisites.Value))
	for _, vnode := range f.System.Prerequisites.Value {
		prereqs = append(prereqs, vnode.Value)
	}

	af.Name = f.Name
	af.Description = f.System.Description.Value
	af.GameMasterDescription = f.System.Description.GameMasterDescription
	af.Title = f.System.Publication.Title
	af.Remaster = f.System.Publication.Remaster
	af.License = f.System.Publication.License
	af.Rarity = f.System.Traits.Rarity
	af.Traits = convertTraits(f.System.Traits.Value)
	af.Rules = string(f.System.Rules)
	af.ActionType = f.System.ActionType.Value
	af.Actions = f.System.Actions.Value
	af.Category = f.System.Category
	af.Level = f.System.Level.Value
	af.Prerequisites = prereqs
	af.GrantsLanguages = f.System.SubFeatures.Languages.Granted
	af.GrantsLanguageCount = f.System.SubFeatures.Languages.Slots
	af.SuppressedFeatures = f.System.SubFeatures.SuppressedFeatures
	af.KeyAbilityOptions = f.System.SubFeatures.KeyOptions

	if f.System.SubFeatures.Senses != nil {
		af.Senses = convertSenses(f.System.SubFeatures.Senses)
	}

	if f.System.SubFeatures.Proficiencies != nil {
		af.Proficiencies = convertProficiencies(f.System.SubFeatures.Proficiencies)
	}

	return af, nil
}

func ConvertAncestry(fa foundry.Ancestry) (storage.Ancestry, error) {
	a := storage.Ancestry{}
	err := validateFoundryAncestry(fa)
	if err != nil {
		return a, err
	}

	a.Name = fa.Name
	a.Description = fa.System.Description.Value
	a.GameMasterDescription = fa.System.Description.GameMasterDescription
	a.Title = fa.System.Publication.Title
	a.Remaster = fa.System.Publication.Remaster
	a.License = fa.System.Publication.License
	a.Rarity = fa.System.Traits.Rarity
	a.Traits = convertTraits(fa.System.Traits.Value)
	a.Rules = string(fa.System.Rules)
	a.Languages = fa.System.Languages.Value
	a.AdditionalLanguageCount = fa.System.AdditionalLanguages.Count
	a.AdditionalLanguages = fa.System.AdditionalLanguages.Value
	a.HP = fa.System.HP
	a.Reach = fa.System.Reach
	a.Size = fa.System.Size
	a.Speed = fa.System.Speed
	a.Vision = fa.System.Vision

	boosts_map := make(map[string][]storage.Boost, 3)
	boosts_map["first"] = make([]storage.Boost, len(fa.System.Boosts.First.Value))
	boosts_map["second"] = make([]storage.Boost, len(fa.System.Boosts.Second.Value))
	boosts_map["third"] = make([]storage.Boost, len(fa.System.Boosts.Third.Value))
	flaw := make([]storage.Boost, len(fa.System.Flaws.First.Value))
	for i, v := range fa.System.Boosts.First.Value {
		boosts_map["first"][i] = storage.Boost{Name: abilityShortNameToLongName[v]}
	}
	for i, v := range fa.System.Boosts.Second.Value {
		boosts_map["second"][i] = storage.Boost{Name: abilityShortNameToLongName[v]}
	}
	for i, v := range fa.System.Boosts.Third.Value {
		boosts_map["third"][i] = storage.Boost{Name: abilityShortNameToLongName[v]}
	}
	for i, v := range fa.System.Flaws.First.Value {
		flaw[i] = storage.Boost{Name: abilityShortNameToLongName[v]}
	}
	a.Boosts = boosts_map
	a.Flaw = flaw

	return a, nil
}

func validateFoundryAncestry(a foundry.Ancestry) error {
	if len(a.System.Boosts.Second.Value) == 1 && len(a.System.Flaws.First.Value) != 1 {
		return fmt.Errorf("Additional boost found without a flaw.")
	}

	if len(a.System.Boosts.First.Value) == 6 && len(a.System.Boosts.Second.Value) == 6 && len(a.System.Boosts.Third.Value) != 0 {
		return fmt.Errorf("Cannot have more than 1 free boost.")
	}

	if a.System.Languages.Custom != "" {
		return fmt.Errorf("Languages.Custom was not null but expected null.")
	}

	if a.System.AdditionalLanguages.Custom != "" {
		return fmt.Errorf("additionalLanguages.Custom was not null but expected null.")
	}

	if len(a.System.Traits.OtherTags) != 0 {
		return fmt.Errorf("Expected len of `traits.otherTags` to be 0.")
	}

	return nil
}

func ConvertBackground(fb foundry.Background) (storage.Background, error) {
	b := storage.Background{}
	err := validateFoundryBackground(fb)
	if err != nil {
		return b, err
	}

	b.Name = fb.Name
	b.Description = fb.System.Description.Value
	b.GameMasterDescription = fb.System.Description.GameMasterDescription
	b.Title = fb.System.Publication.Title
	b.Remaster = fb.System.Publication.Remaster
	b.License = fb.System.Publication.License
	b.Rarity = fb.System.Traits.Rarity
	b.Traits = convertTraits(fb.System.Traits.Value)
	b.Rules = string(fb.System.Rules)

	boosts_map := make(map[string][]storage.Boost, 3)
	boosts_map["first"] = make([]storage.Boost, len(fb.System.Boosts.First.Value))
	boosts_map["second"] = make([]storage.Boost, len(fb.System.Boosts.Second.Value))
	boosts_map["third"] = make([]storage.Boost, len(fb.System.Boosts.Third.Value))
	for i, v := range fb.System.Boosts.First.Value {
		boosts_map["first"][i] = storage.Boost{Name: abilityShortNameToLongName[v]}
	}
	for i, v := range fb.System.Boosts.Second.Value {
		boosts_map["second"][i] = storage.Boost{Name: abilityShortNameToLongName[v]}
	}
	for i, v := range fb.System.Boosts.Third.Value {
		boosts_map["third"][i] = storage.Boost{Name: abilityShortNameToLongName[v]}
	}

	// merge the 2 skill lists
	loreCount := len(fb.System.TrainedSkills.Lore)
	nonLoreCount := len(fb.System.TrainedSkills.Value)
	trainedSkills := make([]storage.Proficiency, loreCount+nonLoreCount)

	// TODO should we care if its a lore skill or not?
	for i, s := range fb.System.TrainedSkills.Lore {
		trainedSkills[i] = storage.Proficiency{Name: s, Rank: rankIntToString[1], Type: "skill"}
	}

	for i, s := range fb.System.TrainedSkills.Value {
		trainedSkills[i+loreCount] = storage.Proficiency{Name: s, Rank: rankIntToString[1], Type: "skill"}
	}

	return b, nil
}

func validateFoundryBackground(fb foundry.Background) error {
	if fb.System.Boosts.Third.Value != nil {
		return fmt.Errorf("Background boosts has unexpected third key. Backgrounds should offer only 2 boosts.")
	}

	if len(fb.System.Items) > 1 {
		return fmt.Errorf("Background items is more than 1. Backgrounds should offer only 1 general feat.")
	}

	if fb.System.TrainedSkills.Custom != "" {
		return fmt.Errorf("Background has a custom trained skill. Expected this to be empty.")
	}

	if len(fb.System.Traits.OtherTags) != 0 {
		return fmt.Errorf("Expected len of `traits.otherTags` to be 0.")
	}

	return nil
}

func validateGeneralFeat(f foundry.Feature) error {
	if f.System.Category != "general" {
		return fmt.Errorf("Expected Category to be 'general'. Was %s", f.System.Category)
	}

	if len(f.System.SubFeatures.SuppressedFeatures) != 0 {
		return fmt.Errorf("Expected subFeature.SuppressedFeatures to be empty.")
	}

	if len(f.System.SubFeatures.Senses) != 0 {
		return fmt.Errorf("Expected subFeature.Senses to be empty.")
	}

	if len(f.System.Traits.OtherTags) != 0 {
		return fmt.Errorf("Expected len of `traits.otherTags` to be 0.")
	}

	if f.System.Frequency.Max != 0 && f.System.Frequency.Per == "" {
		return fmt.Errorf("Expected frequency.Max to be 0 if frequency.Per is blank/empty.")
	}

	if f.System.Frequency.Max == 0 && f.System.Frequency.Per != "" {
		return fmt.Errorf("Expected frequency.Per to be blank/empty if frequency.Max is 0.")
	}

	return nil
}

func ConvertGeneralFeat(f foundry.Feature) (storage.GeneralFeat, error) {
	gf := storage.GeneralFeat{}
	err := validateGeneralFeat(f)
	if err != nil {
		return gf, err
	}

	prereqs := make([]storage.Prerequisite, len(f.System.Prerequisites.Value))
	for _, vnode := range f.System.Prerequisites.Value {
		prereqs = append(prereqs, storage.Prerequisite{Value: vnode.Value})
	}

	gf.Name = f.Name
	gf.Description = f.System.Description.Value
	gf.GameMasterDescription = f.System.Description.GameMasterDescription
	gf.Title = f.System.Publication.Title
	gf.Remaster = f.System.Publication.Remaster
	gf.License = f.System.Publication.License
	gf.Rarity = f.System.Traits.Rarity
	gf.Traits = convertTraits(f.System.Traits.Value)
	gf.Rules = string(f.System.Rules)
	gf.ActionType = f.System.ActionType.Value
	gf.Actions = f.System.Actions.Value
	gf.Category = f.System.Category
	gf.Level = f.System.Level.Value
	gf.Prerequisites = prereqs
	
	if f.System.MaxTakable == 0 {
		gf.MaxTakable = 1
	} else {
		gf.MaxTakable = f.System.MaxTakable
	}

	// require a max and a per both be present in frequency to set either.
	if f.System.Frequency.Max != 0 && f.System.Frequency.Per != "" {
		gf.FrequencyMax = f.System.Frequency.Max
		gf.FrequencyPeriod = f.System.Frequency.Per
	}

	if f.System.SubFeatures.Proficiencies != nil {
		gf.Proficiencies = convertProficiencies(f.System.SubFeatures.Proficiencies)
	}

	return gf, nil
}

func validateClassProperty(f foundry.Feature) error {
	if f.System.Category != "classfeature" {
		return fmt.Errorf("Expected Category to be 'classfeature'. Was %s", f.System.Category)
	}

	// NOTE this could be a false assumption?
	if f.System.ActionType.Value != "passive" {
		return fmt.Errorf("Expected ActionType to be passive. Was %s", f.System.ActionType.Value)
	}

	// NOTE this could be a false assumption?
	if f.System.Actions.Value != 0 {
		return fmt.Errorf("Expected Actions to be null/0.")
	}
	
	if f.System.MaxTakable != 0 {
		return fmt.Errorf("Expected MaxTakable for a classfeature (property) to be 0.")
	}

	if len(f.System.SubFeatures.SuppressedFeatures) != 0 {
		return fmt.Errorf("Expected subFeature.SuppressedFeatures to be empty.")
	}

	if len(f.System.SubFeatures.Senses) != 0 {
		return fmt.Errorf("Expected subFeature.Senses to be empty.")
	}

	if f.System.Frequency.Max != 0 {
		return fmt.Errorf("Expected frequency.Max to be 0.")
	}

	if f.System.Frequency.Per != "" {
		return fmt.Errorf("Expected frequency.Per to be blank/empty.")
	}
	return nil
}

func ConvertClassProperty(f foundry.Feature) (storage.ClassProperty, error) {
	cp := storage.ClassProperty{}
	err := validateClassProperty(f)
	if err != nil {
		return cp, err
	}
	cp.Name = f.Name
	cp.Description = f.System.Description.Value
	cp.GameMasterDescription = f.System.Description.GameMasterDescription
	cp.Title = f.System.Publication.Title
	cp.Remaster = f.System.Publication.Remaster
	cp.License = f.System.Publication.License
	cp.Rarity = f.System.Traits.Rarity
	cp.Traits = convertTraits(f.System.Traits.Value)
	cp.Rules = string(f.System.Rules)
	cp.ActionType = f.System.ActionType.Value
	cp.Actions = f.System.Actions.Value
	cp.Category = f.System.Category
	cp.Level = f.System.Level.Value

	cp.Tags = make([]storage.Tag, len(f.System.Traits.OtherTags))
	for i, tag := range f.System.Traits.OtherTags {
		cp.Tags[i] = storage.Tag{Name: tag}
	}

	cp.KeyAbilities = make([]storage.KeyAbility, len(f.System.SubFeatures.KeyOptions))
	for i, k := range f.System.SubFeatures.KeyOptions {
		cp.KeyAbilities[i] = storage.KeyAbility{Value: k}
	}

	cp.Proficiencies = make([]storage.Proficiency, len(f.System.SubFeatures.Proficiencies))

	count := 0
	for k, v := range f.System.SubFeatures.Proficiencies {
		cp.Proficiencies[count] = storage.Proficiency{Name: k, Rank: rankIntToString[v["rank"]], Type: proficiencyStringToType[k]}
		count++
	}

	return cp, nil
}

func validateClass(c foundry.Class) error {
	if len(c.System.Rules) != 0 {
		return fmt.Errorf("Expected class to have no rules. Found system.Rules len != 0")
	}

	if c.System.Description.GameMasterDescription != "" {
		return fmt.Errorf("Expected class to have no gm description.")
	}

	return nil
}

func ConvertClass(fc foundry.Class) (storage.Class, error) {
	c := storage.Class{}
	err := validateClass(fc)
	if err != nil {
		return c, err
	}

	c.Name = fc.Name
	c.Description = fc.System.Description.Value
	c.Title = fc.System.Publication.Title
	c.Remaster = fc.System.Publication.Remaster
	c.License = fc.System.Publication.License
	c.Rarity = fc.System.Traits.Rarity
	c.HP = fc.System.HP
	c.Perception = rankIntToString[fc.System.Perception]
	c.Spellcasting = rankIntToString[fc.System.Spellcasting]
	c.AdditionalTrainedSkills = fc.System.ClassTrainedSkills.Additional

	featLevelCount := len(fc.System.AncestryFeatLevels.Value) +
		len(fc.System.ClassFeatLevels.Value) +
		len(fc.System.GeneralFeatLevels.Value) +
		len(fc.System.SkillFeatLevels.Value) 
	
	c.FeatLevels = make([]storage.FeatLevel, featLevelCount)
	
	current := 0
	for i, f := range fc.System.AncestryFeatLevels.Value {
		c.FeatLevels[current+i] = storage.FeatLevel{Level: f, Type: "Ancestry"}
		current++
	}

	for i, f := range fc.System.ClassFeatLevels.Value {
		c.FeatLevels[current+i] = storage.FeatLevel{Level: f, Type: "Class"}
		current++
	}

	for i, f := range fc.System.GeneralFeatLevels.Value {
		c.FeatLevels[current+i] = storage.FeatLevel{Level: f, Type: "General"}
		current++
	}

	for i, f := range fc.System.SkillFeatLevels.Value {
		c.FeatLevels[current+i] = storage.FeatLevel{Level: f, Type: "Skill"}
		current++
	}

	c.SkillIncreaseLevels = make([]storage.SkillIncreaseLevel, len(fc.System.SkillIncreaseLevels.Value))
	for i, l := range fc.System.SkillIncreaseLevels.Value {
		c.SkillIncreaseLevels[i] = storage.SkillIncreaseLevel{Level: l}
	}

	c.AttackProficiencies = make([]storage.Proficiency, 5)
	c.AttackProficiencies[0] = storage.Proficiency{Name: "unarmed", Rank: rankIntToString[fc.System.Attacks.Unarmed], Type: "Attack"}
	c.AttackProficiencies[1] = storage.Proficiency{Name: "simple", Rank: rankIntToString[fc.System.Attacks.Simple], Type: "Attack"}
	c.AttackProficiencies[2] = storage.Proficiency{Name: "martial", Rank: rankIntToString[fc.System.Attacks.Martial], Type: "Attack"}
	c.AttackProficiencies[3] = storage.Proficiency{Name: "advanced", Rank: rankIntToString[fc.System.Attacks.Advanced], Type: "Attack"}
	if fc.System.Attacks.Other.Name != "" {
		c.AttackProficiencies[4] = storage.Proficiency{Name: fc.System.Attacks.Other.Name, Rank: rankIntToString[fc.System.Attacks.Other.Rank], Type: "Attack"}
	}


	c.DefenseProficiencies = make([]storage.Proficiency, 4)
	c.DefenseProficiencies[0] = storage.Proficiency{Name: "unarmored", Rank: rankIntToString[fc.System.Defenses.Unarmored], Type: "Defense"}
	c.DefenseProficiencies[1] = storage.Proficiency{Name: "light", Rank: rankIntToString[fc.System.Defenses.Light], Type: "Defense"}
	c.DefenseProficiencies[2] = storage.Proficiency{Name: "medium", Rank: rankIntToString[fc.System.Defenses.Medium], Type: "Defense"}
	c.DefenseProficiencies[3] = storage.Proficiency{Name: "heavy", Rank: rankIntToString[fc.System.Defenses.Heavy], Type: "Defense"}

	c.KeyAbilities = make([]storage.KeyAbility, len(fc.System.KeyAbility.Value))
	for i, k := range fc.System.KeyAbility.Value {
		c.KeyAbilities[i] = storage.KeyAbility{Value: k}
	}

	c.SavingThrowProficiencies = make([]storage.Proficiency, 3)
	c.SavingThrowProficiencies[0] = storage.Proficiency{Name:"fortitude", Rank: rankIntToString[fc.System.SavingThrows.Fortitude], Type: "Saving Throw"}
	c.SavingThrowProficiencies[1] = storage.Proficiency{Name:"reflex", Rank: rankIntToString[fc.System.SavingThrows.Reflex], Type: "Saving Throw"}
	c.SavingThrowProficiencies[2] = storage.Proficiency{Name:"will", Rank: rankIntToString[fc.System.SavingThrows.Will], Type: "Saving Throw"}

	c.TrainedSkills = make([]storage.Proficiency, len(fc.System.ClassTrainedSkills.Value))
	for i, k := range fc.System.ClassTrainedSkills.Value {
		c.TrainedSkills[i] = storage.Proficiency{Name: k, Rank: rankIntToString[1], Type: "Skill"}
	}

	return c, nil
}
