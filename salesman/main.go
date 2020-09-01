package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const (
	popSize     = 100
	citiesNum   = 10
	elitism     = 20
	trnmtSize   = 4 // popsize - elitism must be bigger than this
	mutation    = 0.01
	generations = 500
)

var cities = make([]cityT, citiesNum)

func rankRoutes(population []routeT) []routeT {
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].routeDistance() < population[j].routeDistance()
	})
	return population
}

func bestRouteSlice(population []routeT) routeT {
	population = rankRoutes(population)
	return population[0]
}

func contains(s []int, x int) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}

func selectMatingPool(population []routeT) []routeT {
	var pick int
	var picked, winners []int
	var trnmtRounds int = int((len(population)) / 2)
	var matingSlice []routeT

	for i := 0; i < elitism; i++ {
		matingSlice = append(matingSlice, population[i])
	}

	pick = rand.Intn(len(population)-elitism) + elitism // do inial pick as default 0 is not valid
	for round := 0; round < trnmtRounds; round++ {
		candidates := []routeT{}
		for it := 0; it < trnmtSize; it++ {
			for contains(picked, pick) == true || contains(winners, pick) == true {
				pick = rand.Intn(len(population)-elitism) + elitism
			}
			picked = append(picked, pick)
			candidates = append(candidates, population[pick])
		}
		picked = []int{}
		winners = append(winners, pick)
		matingSlice = append(matingSlice, bestRouteSlice(candidates))
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

func breed(parent1 routeT, parent2 routeT) routeT {
	var child routeT
	geneA := float64(rand.Intn(len(parent1.route)))
	geneB := float64(rand.Intn(len(parent1.route)))
	startGene := int(math.Min(geneA, geneB))
	endGene := int(math.Max(geneA, geneB))
	sequence := parent1.route[startGene:(endGene + 1)]
	for k, city := range parent2.route {
		if k == startGene {
			child.route = append(child.route, sequence...)
		}
		if routeContains(sequence, city) || routeContains(child.route, city) {
			continue
		}
		child.route = append(child.route, city)
	}
	return child
}

func breedPopulation(matingpool []routeT) []routeT {
	var children []routeT
	var poolSize int = len(matingpool)
	for i := 0; i < elitism; i++ {
		children = append(children, matingpool[i])
	}

	idx := rand.Perm(len(matingpool))
	pool := make([]routeT, len(matingpool))
	for i := range idx {
		pool[i] = matingpool[idx[i]]
	}

	for i := 0; i < poolSize; i++ {
		children = append(children, breed(pool[i], pool[len(matingpool)-i-1]))
	}
	return children
}

func nextGeneration(currentGen []routeT, elitism int, mutation float32) []routeT {
	popRanked := rankRoutes(currentGen)
	matingPool := selectMatingPool(popRanked) // select mating candidates from current population
	children := breedPopulation(matingPool)   // breed new childrens and join them with elites
	// nextGeneration = mutatePopulation(children, mutationRate) // mutate childrens
	// return nextGeneration
	return children
}

//? https://towardsdatascience.com/evolution-of-a-salesman-a-complete-genetic-algorithm-tutorial-for-python-6fe5d2b3ca35
func main() {
	rand.Seed(time.Now().UnixNano())
	initCities()
	// fmt.Println(cities)
	population := rankRoutes(initialPopulation())
	// fmt.Println(population)

	fmt.Println("Initial best distance: ", population[0].routeDistance())
	for i := 0; i < generations; i++ {
		population = nextGeneration(population, elitism, mutation)
	}
	fmt.Println("Final distance: ", population[0].routeDistance())
	fmt.Println("Best route: ", population[0].route)
}
