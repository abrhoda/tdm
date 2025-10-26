package main

import (
	"github.com/abrhoda/tdm/pkg/storage"
	"github.com/abrhoda/tdm/foundry"
)

// TODO storage.Sense.Name should be normalized. right now its like low-light-vision, tremorsense, darkvision
func ConvertSenses(sMap map[string]foundry.Sense) []storage.Sense {
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

func ConvertProficiencies(pMap map[string]map[string]int) []storage.Proficiency {
	ps := make([]storage.Proficiency, len(pMap))

	for k,v := range pMap {
		ps = append(ps, storage.Proficiency{ Name: k, Rank: v["rank"] })
	}

	return ps
}

func ConvertAncestryFeature(f foundry.Feature) storage.AncestryFeature {
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
		af.Senses = ConvertSenses(f.System.SubFeatures.Senses)
	}

	if f.System.SubFeatures.Proficiencies != nil {
		af.Proficiencies = ConvertProficiencies(f.System.SubFeatures.Proficiencies)
	}

	return af
}

func ConvertAncestry(a foundry.Ancestry) storage.Ancestry {
	return storage.Ancestry{}
}
