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
	mapSize     = 20
	elitism     = 10
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

func bestRoutePop(population []routeT) routeT {
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
	var pick int
	var picked []int
	var trnmtRounds = int((len(population)) / 2)
	var matingSlice []routeT

	if elitism+trnmtRounds+trnmtSize-1 > len(population) {
		fmt.Print("no room for tournament. prevent population extinction.\n")
		return population
	}

	matingSlice = population[:elitism]

	pick = rand.Intn(len(population)-elitism) + elitism // do inial pick as default 0 is not valid
	for round := 0; round < trnmtRounds; round++ {
		candidates := []routeT{}
		for it := 0; it < trnmtSize; it++ {
			for contains(picked, pick) {
				pick = rand.Intn(len(population)-elitism) + elitism
			}
			picked = append(picked, pick)
			candidates = append(candidates, population[pick])
		}
		// winners can appear multiple times in mating pool
		picked = []int{}
		matingSlice = append(matingSlice, bestRoutePop(candidates))
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
	geneA := float64(rand.Intn(len(parent1.route)-2) + 1)
	geneB := float64(rand.Intn(len(parent1.route)-2) + 1)
	startGene := int(math.Min(geneA, geneB) - 1)
	endGene := int(math.Max(geneA, geneB) + 1)
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

func breedPopulation(matingpool []routeT) []routeT {
	var children []routeT
	var poolSize = len(matingpool)
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
	children := breedPopulation(matingPool)   // breed new childrens
	// nextGeneration = mutatePopulation(children, mutationRate) // mutate childrens
	// return nextGeneration
	nextGeneration := rankRoutes(killClones(append(matingPool, children...)))
	if len(nextGeneration) > popSize {
		nextGeneration = nextGeneration[:popSize]
	}
	return nextGeneration
}

//? https://towardsdatascience.com/evolution-of-a-salesman-a-complete-genetic-algorithm-tutorial-for-python-6fe5d2b3ca35
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
	for i := 0; i < 10; i++ {
		fmt.Printf("Final distance: %.3f , Best route: %s\n", population[i].routeDistance(), population[i].toString())
	}
}
