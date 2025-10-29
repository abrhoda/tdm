package main

import (
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

func convertAncestryFeature(f foundry.Feature) storage.AncestryFeature {
	prereqs := make([]string, len(f.System.Prerequisites.Value))
	for _, vnode := range f.System.Prerequisites.Value {
		prereqs = append(prereqs, vnode.Value)
	}

	af := storage.AncestryFeature{
		Name: f.Name,
	  Description: f.System.Description.Value,
	  GameMasterDescription: f.System.Description.GameMasterDescription,
		Title: f.System.Publication.Title,
		Remaster:f.System.Publication.Remaster,
		License:f.System.Publication.License,
		Rarity: f.System.Traits.Rarity,
		Traits: f.System.Traits.Value,
		Rules: f.System.Rules,
		ActionType: f.System.ActionType.Value,
		Actions: f.System.Actions.Value,
		Category: f.System.Category,
		Level: f.System.Level.Value,
		Prerequisites: prereqs,
		GrantsLanguages: f.System.SubFeatures.Languages.Granted,
		GrantsLanguageCount: f.System.SubFeatures.Languages.Slots,
		SuppressedFeatures: f.System.SubFeatures.SuppressedFeatures,
		KeyAbilityOptions: f.System.SubFeatures.KeyOptions,
	}

	if f.System.SubFeatures.Senses != nil {
		af.Senses = convertSenses(f.System.SubFeatures.Senses)
	}

	if f.System.SubFeatures.Proficiencies != nil {
		af.Proficiencies = convertProficiencies(f.System.SubFeatures.Proficiencies)
	}

	return af
}

func convertAncestry(fa foundry.Ancestry) storage.Ancestry {
	a := storage.Ancestry{
		Name: fa.Name,
		Description: fa.System.Description.Value,
	  GameMasterDescription: fa.System.Description.GameMasterDescription,
		Title: fa.System.Publication.Title,
		Remaster:fa.System.Publication.Remaster,
		License:fa.System.Publication.License,
		Rarity: fa.System.Traits.Rarity,
		Traits: fa.System.Traits.Value,
		Rules: fa.System.Rules,
		FreeBoost: "any",
		Languages: fa.System.Languages.Value,
		AdditionalLanguageCount: fa.System.AdditionalLanguages.Count,
		AdditionalLanguages: fa.System.AdditionalLanguages.Value,
		HP: fa.System.HP,
		Reach: fa.System.Reach,
		Size: fa.System.Size,
		Speed: fa.System.Speed,
		Vision: fa.System.Vision,
	}

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

	return a
}
