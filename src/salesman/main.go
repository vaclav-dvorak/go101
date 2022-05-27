package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	popSize     = 100
	citiesNum   = 10
	mapSize     = 20
	elitism     = 2
	trnmtRounds = 5 // popsize - elitism must be bigger than this
	mutation    = 0.015
	generations = 500
)

var cities = make([]cityT, citiesNum)

func rankRoutes(population []routeT) []routeT {
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].routeDistance() < population[j].routeDistance()
	})
	return population
}

func killClones(pop []routeT) (ret []routeT) {
	hashMap := make(map[string]struct{}, len(pop))
	for i := 0; i < len(pop); i++ {
		if _, ok := hashMap[pop[i].toString()]; ok {
			continue
		} else {
			hashMap[pop[i].toString()] = struct{}{}
			ret = append(ret, pop[i])
		}
	}
	fmt.Printf("clones killed = %02d, survived = %02d\n", len(pop)-len(ret), len(ret))
	return
}

func routeContains(s []cityT, x cityT) bool {
	for i := 0; i < len(s); i++ {
		if x.name == s[i].name {
			return true
		}
	}
	return false
}

func breed(parent1 routeT, parent2 routeT) (child routeT) {
	startGene := rand.Intn(len(parent1.route))
	endGene := rand.Intn(len(parent1.route))
	if startGene > endGene {
		startGene, endGene = endGene, startGene
	}
	sequence := parent1.route[startGene:endGene]
	for k, city := range parent2.route {
		if k == startGene {
			child.route = append(child.route, sequence...)
		}
		if routeContains(sequence, city) || routeContains(child.route, city) {
			continue
		}
		child.route = append(child.route, city)
	}

	return
}

func mutatePopulation(pop []routeT) (ret []routeT) {
	ret = make([]routeT, len(pop))
	mut, mutants := 0, 0
	mutated := false
	copy(ret, pop)
	for i := 0; i < len(ret); i++ {
		mutated = false
		for c := 0; c < len(ret[i].route); c++ {
			if rand.Float64() < mutation {
				mutated = true
				swap := rand.Intn(len(ret[i].route))
				ret[i].route[c], ret[i].route[swap] = ret[i].route[swap], ret[i].route[c]
				mut++
			}
		}
		if mutated {
			mutants++
		}
	}
	fmt.Printf("mutations = %02d,  mutants = %02d, ", mut, mutants)
	return
}

func tournament(pop []routeT) (winner routeT) {
	winner = routeT{distance: 1000000}
	for it := 0; it < trnmtRounds; it++ {
		pick := rand.Intn(len(pop))
		if winner.routeDistance() > pop[pick].routeDistance() {
			winner = pop[pick]
		}
	}
	return
}

func nextGeneration(currentGen []routeT) []routeT {
	popRanked := rankRoutes(currentGen)
	children := make([]routeT, popSize-elitism)
	for i := elitism; i < popSize; i++ {
		parent1 := tournament(popRanked)
		parent2 := tournament(popRanked)
		children[i-elitism] = breed(parent1, parent2)
	}
	children = mutatePopulation(children)
	newPop := popRanked[:elitism]
	newPop = append(newPop, children...)
	nextGeneration := killClones(newPop)
	return nextGeneration
}

// https://gist.github.com/turbofart/3428880
func main() {
	if citiesNum > 62 {
		fmt.Println("Number of cities cannot be bigger then 62")
		return
	}
	rand.Seed(time.Now().UnixNano())
	initCities()
	population := rankRoutes(initialPopulation())

	fmt.Println("Initial best distance: ", population[0].routeDistance())
	for i := 0; i <= generations; i++ {
		fmt.Printf("%03d: ", i)
		population = nextGeneration(population)
	}
	population = rankRoutes(population)
	for i := 0; i < citiesNum; i++ {
		fmt.Printf("Final distance: %.3f , Best route: %s\n", population[i].routeDistance(), population[i].toString())
	}
}
