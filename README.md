# tdm
Tabletop data manager (tdm) tool.

## Steps
### Pathfinder 2nd Edition
#### Build command basics
- [ ] read all files under packs dir.
- [ ] read ancestries into ancestry struct
- [ ] read ancestryfeatures into ancestryFeature struct
- [ ] read backgounds into background struct
- [ ] read classes into class struct
- [ ] read classfeatures into classFeature struct
- [ ] read equipment into proper struct based on type field. provide common json attrs in equipment struct
- [ ] read equipment-effects into equipmentEffects struct
- [ ] read feats into feats struct
- [ ] read feat-effects into featEffects struct
- [ ] read heritages into heritage struct
- [ ] read spells into spell struct
- [ ] read spell-effects into spellEffect struct
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
