package internal

import (
	"fmt"
	"github.com/abrhoda/tdm/internal/foundry"
	"github.com/abrhoda/tdm/storage"
)

// TODO storage.Sense.Name should be normalized. right now its like low-light-vision, tremorsense, darkvision
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
		s = append(s, storage.Sense{Name: k, Acuity: v.Acuity, Range: v.Range, ElevateIfHasLowLightVision: elevate})
	}
	return s
}

func convertProficiencies(pMap map[string]map[string]int) []storage.Proficiency {
	ps := make([]storage.Proficiency, len(pMap))

	for k, v := range pMap {
		ps = append(ps, storage.Proficiency{Name: k, Rank: v["rank"]})
	}

	return ps
}

func validateFoundryAncestryFeature(a foundry.Feature) error {
	if a.System.ActionType.Value != "passive" {
		return fmt.Errorf("Expected ActionType to be passive. Was %s", a.System.ActionType.Value)
	}

	if a.System.Actions.Value != 0 {
		return fmt.Errorf("Expected Actions to be null/0.")
	}

	if a.System.Category != "ancestryfeature" {
		return fmt.Errorf("Expected Category to be 'ancestryfeature. Was %s'", a.System.Category)
	}

	if a.System.Level.Value > 1 {
		return fmt.Errorf("Expected Level to be 0 or 1.")
	}

	if len(a.System.Prerequisites.Value) != 0 {
		return fmt.Errorf("Expected prerequisites to be empty.")
	}

	if len(a.System.SubFeatures.KeyOptions) != 0 {
		return fmt.Errorf("Expected subFeature.KeyOptions to be empty.")
	}

	if len(a.System.SubFeatures.Proficiencies) != 0 {
		return fmt.Errorf("Expected subFeature.Proficiencies to be empty.")

	}

	if len(a.System.SubFeatures.SuppressedFeatures) != 0 {
		return fmt.Errorf("Expected subFeature.SuppressedFeatures to be empty.")

	}

	return nil
}

func ConvertAncestryFeature(f foundry.Feature) (storage.AncestryFeature, error) {
	af := storage.AncestryFeature{}
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
	af.Traits = f.System.Traits.Value
	af.Rules = f.System.Rules
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
	a.Traits = fa.System.Traits.Value
	a.Rules = fa.System.Rules
	a.FreeBoost = "any"
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
		boosts_map["first"][i] = storage.Boost{Name: v}
	}
	for i, v := range fa.System.Boosts.Second.Value {
		boosts_map["second"][i] = storage.Boost{Name: v}
	}
	for i, v := range fa.System.Boosts.Third.Value {
		boosts_map["third"][i] = storage.Boost{Name: v}
	}
	for i, v := range fa.System.Flaws.First.Value {
		flaw[i] = storage.Boost{Name: v}
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
	b.Traits = fb.System.Traits.Value
	b.Rules = fb.System.Rules

	boosts_map := make(map[string][]storage.Boost, 3)
	boosts_map["first"] = make([]storage.Boost, len(fb.System.Boosts.First.Value))
	boosts_map["second"] = make([]storage.Boost, len(fb.System.Boosts.Second.Value))
	boosts_map["third"] = make([]storage.Boost, len(fb.System.Boosts.Third.Value))
	for i, v := range fb.System.Boosts.First.Value {
		boosts_map["first"][i] = storage.Boost{Name: v}
	}
	for i, v := range fb.System.Boosts.Second.Value {
		boosts_map["second"][i] = storage.Boost{Name: v}
	}
	for i, v := range fb.System.Boosts.Third.Value {
		boosts_map["third"][i] = storage.Boost{Name: v}
	}

	// merge the 2 skill lists
	loreCount := len(fb.System.TrainedSkills.Lore)
	nonLoreCount := len(fb.System.TrainedSkills.Value)
	skills := make([]storage.Skill, loreCount+nonLoreCount)

	// TODO should we care if its a lore skill or not?
	for i, s := range fb.System.TrainedSkills.Lore {
		skills[i] = storage.Skill{Name: s}
	}

	for i, s := range fb.System.TrainedSkills.Value {
		skills[i+loreCount] = storage.Skill{Name: s}
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

	return nil
}
