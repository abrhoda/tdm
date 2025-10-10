# tdm
Tabletop data manager (tdm) tool.

## Steps
### Pathfinder 2nd Edition
#### Build command basics
- [ ] read all files under packs dir.
- [ ] create structs for each subdir under packs dir. Equipment should be multiple structs based on the json's "type" key.
- [ ] apply filters
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
