module github.com/timokae/pokedex

replace github.com/timokae/pokedex/internal/pokeapi => ./internal/pokeapi

replace github.com/timokae/pokedex/internal/pokecache => ./internal/pokecache

go 1.22.2

require (
	github.com/timokae/pokedex/internal/pokeapi v0.0.0
	github.com/timokae/pokedex/internal/pokecache v0.0.0 // indirect
)
