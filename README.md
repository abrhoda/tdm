# tdm
Tabletop data manager (tdm) tool.

## Steps
### Pathfinder 2nd Edition
#### Build command basics
- [x] read all files under packs dir.
- [x] read ancestries into ancestry struct
- [x] read ancestryfeatures into ancestryFeature struct
- [x] read backgounds into background struct
- [x] read classes into class struct
- [x] read classfeatures into classFeature struct
- [x] read equipment into proper struct based on type field. provide common json attrs in equipment struct
- [x] read equipment-effects into equipmentEffects struct
- [x] read feats into feats struct
- [x] read feat-effects into featEffects struct
- [x] read heritages into heritage struct
- [ ] read spells into spell struct
- [ ] read spell-effects into spellEffect struct
- [ ] create database models to sanitize the foundry model structure, fields, and text.
- [ ] connect entities such as class to class features.
- [ ] read other-effects into otherEffect (maybe just called effect?) struct
- [ ] read read deities into deity struct
- [ ] read read conditions into condition struct
- [ ] read read hazards into hazard struct
- [ ] read read actions into action struct
- [ ] read bestiaries (all bestiary dirs) into creatures struct
- [ ] apply license filter (this could be done before reading into struct)
- [ ] apply legacy content filter (this could be done before reading into struct)
- [ ] apply text filters
- [ ] save in db

#### CRUD actions
- [ ] GET resource\_type resource\_name (standard + homebrew)
- [ ] POST resource\_type (create a homebrew)
- [ ] UPDATE resource\_type resource\_name (update a homebrew)
- [ ] DELETE resource\_type resource\_name (delete a homebrew)

#### Party Actions
- [ ] CREATE party entity (with name?)
- [ ] CREATE characters using a basic character sheet. (allow pathbuilder imports too)
- [ ] ADD character to party
- [ ] UPDATE + DELETE party
- [ ] UPDATE + DELETE characters

### Dungeons & Dragons 5th Edition (2024)
- [ ] Read SRD PDF file.


## Potential Improvements
1. In `build.go`, initialize slices with a reasonable capacity to avoid many unneeded allocations as it resizes the underlying memory.
2. In `build.go`, disallow unknown fields in the json fields for the expected type. Foundryvtt/pf2e uses a VERY loose json structure for thigns of the same type so this would break when a new key is introduced. This would be useful to allow for knowing and updating when some new json key is introduced into the files that isn't expected. 
```
var data T
decoder := json.NewDecoder(strings.NewReader(string(content)))
decoder.DisallowUnknownFields()
decoder.Decode(&data)
```
3. `rule.go` and `rules.md` contains the starts of unwinding the insane rules structure used in foundryvtt/pf2e. These rules heavily rely on javascript's loose rules and are messy to unwind. However, this can be done later to apply some more advanced features. For example, picking the "Fangs" ancestryfeature for an Anadi could apply that directly to the character sheet.
