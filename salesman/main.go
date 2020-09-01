package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	popSize     = 5
	citiesNum   = 10
	elitism     = 20
	mutation    = 0.1
	generations = 500
)

var cities = make([]cityT, citiesNum)

func rankRoutes(population []routeT) []routeT {
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].routeDistance() < population[j].routeDistance()
	})
	return population
}

func nextGeneration(p int, elitism int, mutation float32) (r int) {
	// popRanked = rankRoutes(currentGen)
	// selectionResults = selection(popRanked, eliteSize) // select mating candidates from current population
	// matingpool = matingPool(currentGen, selectionResults) // put candidates into mating pool
	// children = breedPopulation(matingpool, eliteSize) // breed new childrens and join them with elites
	// nextGeneration = mutatePopulation(children, mutationRate) // mutate childrens
	// return nextGeneration
	r = 5
	return
}

//? https://towardsdatascience.com/evolution-of-a-salesman-a-complete-genetic-algorithm-tutorial-for-python-6fe5d2b3ca35
func main() {
	rand.Seed(time.Now().UnixNano())
	initCities()
	// fmt.Println(cities)
	population := initialPopulation()
	// fmt.Println(population)
	population = rankRoutes(population)
	fmt.Println("Initial best distance: ", population[0].routeDistance())
	// for i := 0; i < generations; i++ {
	// pop = nextGeneration(pop, elitism, mutation)
	// }
	fmt.Println("Final distance: ", 5)
}
