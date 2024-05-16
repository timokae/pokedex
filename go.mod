module github.com/timokae/pokedex

replace github.com/timokae/pokedex/internal/pokeapi => ./internal/pokeapi

go 1.22.2

require (
  github.com/timokae/pokedex/internal/pokeapi v0.0.0
)
