module github.com/timokae/pokedex/internal/pokeapi

replace github.com/timokae/pokedex/internal/pokecache => ../pokecache

go 1.22.2

require "github.com/timokae/pokedex/internal/pokecache" v0.0.0
