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
	elitism     = 1
	trnmtSize   = 5 // popsize - elitism must be bigger than this
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
	fmt.Printf("clones killed = %d, survived = %d\n", len(pop)-len(ret), len(ret))
	return
}

func selectMatingPool(population []routeT) []routeT {
	var (
		trnmtRounds = int((len(population)) / 2)
		matingSlice []routeT
	)

	if elitism >= len(population) {
		fmt.Print("no room for tournament. prevent population extinction.\n")
		return population
	}

	matingSlice = population[:elitism]

	for round := 0; round < trnmtRounds; round++ {
		winner := routeT{distance: 1000000}
		for it := 0; it < trnmtSize; it++ {
			pick := rand.Intn(len(population)-elitism) + elitism
			if winner.routeDistance() > population[pick].routeDistance() {
				winner = population[pick]
			}
		}
		// winners can appear multiple times in mating pool
		matingSlice = append(matingSlice, winner)
	}
	return matingSlice
}

func routeContains(s []cityT, x cityT) bool {
	for _, n := range s {
		if x.name == n.name {
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

func breedPopulation(pop []routeT) (children []routeT) {
	idx := rand.Perm(len(pop))
	pool := make([]routeT, len(pop))
	for i := range idx {
		pool[i] = pop[idx[i]]
	}

	for i := 0; i < len(pop); i++ {
		children = append(children, breed(pool[i], pool[len(pop)-i-1]))
	}
	return
}

func mutatePopulation(pop []routeT) (mutants []routeT) {
	ret := make([]routeT, len(pop))
	mut := 0
	mutated := false
	mutants = make([]routeT, 0)
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
			mutants = append(mutants, ret[i])
		}
	}
	if mut != 0 {
		fmt.Printf("did %d mutations in this generation, le(mutants) = %d\n", mut, len(mutants))
	}
	return
}

func nextGeneration(currentGen []routeT, elitism int, mutation float32) []routeT {
	popRanked := rankRoutes(currentGen)
	matingPool := selectMatingPool(popRanked)                  // select mating candidates from current population
	children := breedPopulation(matingPool)                    // breed new childrens
	children = append(children, mutatePopulation(children)...) // mutate childrens

	// this whole thing should be different
	// -take one elite route at beginning
	// - from 0 to popsize lets play tournament to get two winning parents
	// - breed child out of them
	// - add mutants if needed
	// - sort it by distance and let next generation play tournament with whole before generation

	nextGeneration := killClones(append(matingPool, children...))
	if len(nextGeneration) > 2*popSize {
		nextGeneration = nextGeneration[:2*popSize]
	}
	return nextGeneration
}

// https://towardsdatascience.com/evolution-of-a-salesman-a-complete-genetic-algorithm-tutorial-for-python-6fe5d2b3ca35
//? https://gist.github.com/turbofart/3428880
func main() {
	if citiesNum > 62 {
		fmt.Println("Number of cities cannot be bigger then 62")
		return
	}
	rand.Seed(time.Now().UnixNano())
	initCities()
	population := rankRoutes(initialPopulation())
	// fmt.Println(population)

	fmt.Println("Initial best distance: ", population[0].routeDistance())
	for i := 0; i <= generations; i++ {
		population = nextGeneration(population, elitism, mutation)
		if i%50 == 0 {
			fmt.Printf("Gen %d complete. Current population: %d\n", i, len(population))
		}
	}
	for i := 0; i < citiesNum; i++ {
		fmt.Printf("Final distance: %.3f , Best route: %s\n", population[i].routeDistance(), population[i].toString())
	}
}
