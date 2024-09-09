module github.com/thewerther/pokedexGo

go 1.23.0

require internal/pokeApi v1.2.3
replace internal/pokeApi => ./internal/pokeApi

require internal/pokeCache v1.2.3
replace internal/pokeCache => ./internal/pokeCache
