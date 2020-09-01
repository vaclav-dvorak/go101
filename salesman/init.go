package main

import (
	"fmt"
	"math/rand"
)

func initCities() {
	for i := range cities {
		cities[i] = cityT{fmt.Sprintf("city%d", i), rand.Intn(100), rand.Intn(100)}
	}
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
