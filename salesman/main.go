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
	generations = 1
)

var cities = make([]cityT, citiesNum)

func rankRoutes(population []routeT) []routeT {
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].routeDistance() < population[j].routeDistance()
	})
	return population
}

func nextGeneration(currentGen []routeT, elitism int, mutation float32) []routeT {
	popRanked := rankRoutes(currentGen)
	// selectionResults = selection(popRanked, eliteSize) // select mating candidates from current population
	// matingpool = matingPool(currentGen, selectionResults) // put candidates into mating pool
	// children = breedPopulation(matingpool, eliteSize) // breed new childrens and join them with elites
	// nextGeneration = mutatePopulation(children, mutationRate) // mutate childrens
	// return nextGeneration
	return popRanked
}

//? https://towardsdatascience.com/evolution-of-a-salesman-a-complete-genetic-algorithm-tutorial-for-python-6fe5d2b3ca35
func main() {
	rand.Seed(time.Now().UnixNano())
	initCities()
	// fmt.Println(cities)
	population := initialPopulation()
	// fmt.Println(population)

	fmt.Println("Initial best distance: ", population[0].routeDistance())
	for i := 0; i < generations; i++ {
		population = nextGeneration(population, elitism, mutation)
	}
	fmt.Println("Final distance: ", population[0].routeDistance())
}
