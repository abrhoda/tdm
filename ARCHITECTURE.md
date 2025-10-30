# Architecture

## Project File Structure
```
.
|-- ARCHITECTURE.md
|-- build/
|   |-- Dockerfile
|   |-- docker-compose.yaml
|-- cmd/
|   |-- pf2e.go
|-- go.mod
|-- internal/
|   |-- foundry/
|   |    |-- ancestry.go
|   |    |-- background.go
|   |    |-- class.go
|   |    |-- common.go
|   |    |-- dataset.go
|   |    |-- effect.go
|   |    |-- equipment.go
|   |    |-- feat.go
|   |    |-- heritage.go
|   |    |-- journal.go
|   |    |-- rule.go
|   |    |-- rules.md
|   |    |-- spell.go
|   |-- convert.go
|   |-- sanitize.go
|   |-- string.go
|-- LICENSE
|-- Makefile
|-- main.go
|-- storage/
|   |-- model.go
|   |-- store.go

```
Description of dirs and files of note:
 - `build/` contains build specific files. These are only docker related files for now.
 - `cmd/` is where entrypoint commands are located. Add a new file here to have a new entrypoint. Currently only `pf2e` commands exist.
 - `internal/foundry/` contains the slop of unmarshalling types to handle the wild data choices being made in `foundryvtt/pf2e/packs/` (where the dataset for pathfinder is sourced).
 - `internal/convert.go` is where all foundry model -> database model conversion code is located.
 - `internal/sanitize.go` is where all data sanitization functions for cleansing fields in the foundry models.
 - `storage/model.go` holds all database models that line up 1 to 1 in the output dataset.
 - `storage/store.go` datastore interface and implementations.

## Building the Dataset (pf2e)
The path that should be followed for building the file dataset:
1. cmd/pf2e.go reads in args and build options
2. build func is called with build opts and reads json files under `packs/` dir
3. each file is converted to it's corresponding `internal/foundry/` data type
4. foundry data is not added to the dataset if filter options were passed in and the data does not meet the criteria (passed in cli)
5. sanitize the foundry data model by removing compendium tags and html tags as well as removing sanitize options (passed in cli)
6. convert from the foundry model to the storage model
7. insert in dataset selected in the datastore options (passed in cli)
