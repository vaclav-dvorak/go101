package main

import (
	"math/rand"
)

const names = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456798"

var letters = []rune(names)

func initCities() {
	cities = []cityT{{"A", 11, 15}, {"B", 9, 7}, {"C", 1, 11}, {"D", 13, 18}, {"E", 3, 18}, {"F", 1, 11}, {"G", 19, 2}, {"H", 2, 13}, {"I", 2, 19}, {"J", 10, 8} }
	//* This can be used when we want random cities
	// for i := range cities {
	// 	cities[i] = cityT{string(letters[i]), rand.Intn(mapSize), rand.Intn(mapSize)}
	// }
}

func initialPopulation() []routeT {
	population := make([]routeT, popSize)
	for popi := range population {
		idx := rand.Perm(citiesNum)
		sample := make([]cityT, citiesNum)
		for i := range idx {
			sample[i] = cities[idx[i]]
		}
		population[popi] = routeT{route: sample}
	}
	return population
}
