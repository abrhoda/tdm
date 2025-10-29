package main

import (
	"fmt"
	"github.com/abrhoda/tdm/pkg/storage"
	"github.com/abrhoda/tdm/foundry"
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
			s = append(s, storage.Sense { Name: k, Acuity: v.Acuity, Range: v.Range, ElevateIfHasLowLightVision: elevate} )
	}
	return s
}

func convertProficiencies(pMap map[string]map[string]int) []storage.Proficiency {
	ps := make([]storage.Proficiency, len(pMap))

	for k,v := range pMap {
		ps = append(ps, storage.Proficiency{ Name: k, Rank: v["rank"] })
	}

	return ps
}

func validateFoundryAncestryFeature(a foundry.Feature) error {
		if (a.System.ActionType.Value != "passive") {
			return fmt.Errorf("Expected ActionType to be passive. Was %s", a.System.ActionType.Value)
		}

		if (a.System.Actions.Value != 0) {
			return fmt.Errorf("Expected Actions to be null/0.")
		}

		if (a.System.Category != "ancestryfeature") {
			return fmt.Errorf("Expected Category to be 'ancestryfeature. Was %s'", a.System.Category)
		}

		if (a.System.Level.Value > 1) {
			return fmt.Errorf("Expected Level to be 0 or 1.")
		}

		if (len(a.System.Prerequisites.Value) != 0) {
			return fmt.Errorf("Expected prerequisites to be empty.")
		}

		if (len(a.System.SubFeatures.KeyOptions) != 0) {
			return fmt.Errorf("Expected subFeature.KeyOptions to be empty.")
		}
		
		if (len(a.System.SubFeatures.Proficiencies) != 0) {
			return fmt.Errorf("Expected subFeature.Proficiencies to be empty.")

		}
		
		if (len(a.System.SubFeatures.SuppressedFeatures) != 0) {
			return fmt.Errorf("Expected subFeature.SuppressedFeatures to be empty.")
		
		}

	return nil
}

func convertAncestryFeature(f foundry.Feature) (storage.AncestryFeature, error) {
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
	af.Remaster =f.System.Publication.Remaster
	af.License =f.System.Publication.License
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

func convertAncestry(fa foundry.Ancestry) (storage.Ancestry, error) {
	a := storage.Ancestry{}
	err := validateFoundryAncestry(fa)
	if err != nil {
		return a, err
	}

	a.Name = fa.Name
	a.Description = fa.System.Description.Value
	a.GameMasterDescription = fa.System.Description.GameMasterDescription
	a.Title = fa.System.Publication.Title
	a.Remaster =fa.System.Publication.Remaster
	a.License =fa.System.Publication.License
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

	if len(fa.System.Boosts.First.Value) == 6 {
		a.FirstBoost = "any"
	} else {
		a.FirstBoost = fa.System.Boosts.First.Value[0]
	}

	// UNWIND THE INSANE CHOICE IN FOUNDRYVTT/PF2E:
	// 1. Any ancestry will use boosts.First to set the first boost
	// 2. for any non human, boost.Second is empty UNLESS the ancestry also has a flaw.
	// 		if the ancestry has a flaw, boost.Second is a list of 1 to boost and Flaws.First is the flaw
	// 3. if boosts.Second is empty because the ancestry has no flaw, boosts.Third is all attributes (free boost)
	// 3. for humans, both boost.First and boost.Second are all attributes??? because they just didnt want to follow the pattern of using boosts 1 and 3??? i hate these people and their insane choices.
	if len(fa.System.Boosts.Second.Value) != 0 && len(fa.System.Flaws.First.Value) != 0 {
		a.Flaw = fa.System.Flaws.First.Value[0]
		a.SecondBoost = fa.System.Boosts.Second.Value[0]
	}

	return a, nil
}

func validateFoundryAncestry(a foundry.Ancestry) error {
	if (len(a.System.Boosts.Second.Value) == 1 && len(a.System.Flaws.First.Value) != 1) {
		return fmt.Errorf("Additional boost found without a flaw.")
	}

	if (len(a.System.Boosts.First.Value) == 6 && len(a.System.Boosts.Second.Value) == 6 && len(a.System.Boosts.Third.Value) != 0) {
		return fmt.Errorf("Cannot have more than 1 free boost.")
	}

	if (a.System.Languages.Custom != "") {
		return fmt.Errorf("Languages.Custom was not null but expected null.")
	}
	
	if (a.System.AdditionalLanguages.Custom != "") {
		return fmt.Errorf("additionalLanguages.Custom was not null but expected null.")
	}

	return nil
}
