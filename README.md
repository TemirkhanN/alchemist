# alchemist
![Code Quality Check](https://github.com/TemirkhanN/alchemist/workflows/code-quality-check/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/TemirkhanN/alchemist)](https://goreportcard.com/report/github.com/TemirkhanN/alchemist)

This is an implementation of an alchemy system from The Elder Scrolls IV: Oblivion.
It uses https://github.com/faiface/pixel as OpenGL adapter to create windows and draw elements.

## Requirements
https://github.com/faiface/pixel#requirements

## Launch

1. Meet [requirements](#requirements)
2. `git clone`
3. `go run cmd/alchemy/main.go`


## Domain

Currently, only mortar is implemented.
Retort, Calcinator and Alembic are in backlog.
Some ingredients sprites are missing. They are commented in ingredient_repository.


## GUI

Alchemist level, luck and mortar level are currently hardcoded in main.go.
Ingredient replacement needs to be implemented yey.

![ingredients](.docs/main.jpeg)
![ingredients](.docs/ingredients.jpeg)
![ingredients](.docs/main-ingredients-filled.jpeg)
