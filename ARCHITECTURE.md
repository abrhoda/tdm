# Architecture

## Project File Structure
```
.
|-- ARCHITECTURE.md
|-- podman/
|   |-- .env
|   |-- local-compose.yaml
|-- cmd/
|   |-- main.go
|-- go.mod
|-- build.go
|-- config.go
|-- internal/
|   |-- convert.go
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
|   |-- sanitize.go
|   |-- string.go
|-- LICENSE
|-- Makefile
|-- sql/
|-- |-- pf2e.sql
|-- storage/
|   |-- model.go
|   |-- store.go

```
Description of dirs and files of note:
 - `podman/` contains podman/docker specific files. `local-compose.yaml` allows for a built postgres and pgadmin interface.
 - `cmd/` is where entrypoint commands are located. Add a new file here to have a new entrypoint. Currently only `main.go` exists as an entrypoint.
 - `internal/foundry/` contains the slop of unmarshalling types to handle the wild data choices being made in `foundryvtt/pf2e/packs/` (where the dataset for pathfinder is sourced).
 - `internal/handlers/` handler funcs for actions that can be reused regardless of entrypoint. 
 - `internal/convert.go` is where all foundry model -> database model conversion code is located.
 - `internal/sanitize.go` is where all data sanitization functions for cleansing fields in the foundry models.
 - `storage/model.go` holds all database models that line up 1 to 1 in the output dataset.
 - `storage/store.go` datastore interface and implementations.
 - `build.go` entrypoint for the `Build` func which takes a configuration struct, handles reading the `foundryvtt/pf2e/packs` files, formatting + sanitizing, and then outputs the final dataset to the desired format.
 - `config.go` handles the configuration settings of how the final dataset should be built.

## Building the Dataset (pf2e)
The path that should be followed for building the file dataset:
1. cmd/pf2e.go reads in args and build options
2. build func is called with build opts and reads json files under `packs/` dir
3. each file is converted to it's corresponding `internal/foundry/` data type
4. foundry data is not added to the dataset if filter options were passed in and the data does not meet the criteria (passed in cli)
5. sanitize the foundry data model by removing compendium tags and html tags as well as removing sanitize options (passed in cli)
6. convert from the foundry model to the storage model
7. insert in dataset selected in the datastore options (passed in cli)
