package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	popSize     = 100
	elitism     = 20
	mutation    = 0.1
	generations = 500
)

type city struct {
	name  string
	x, y  int
	dists map[string]float64
}

func (c *city) distance(t city) float64 {
	if c.dists == nil {
		c.dists = make(map[string]float64)
	}
	if c.dists[t.name] != 0 {
		return c.dists[t.name]
	}
	xDis := (c.x - t.x)
	yDis := (c.y - t.y)
	c.dists[t.name] = math.Sqrt(float64((xDis * xDis) + (yDis * yDis)))
	return c.dists[t.name]
}

const cities = 10

func initialPopulation(n int, c int) (r int) {
	r = 10
	return
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
	pop := initialPopulation(popSize, cities)
	fmt.Println("Initial best distance: ", 10)
	for i := 0; i < generations; i++ {
		pop = nextGeneration(pop, elitism, mutation)
	}
	fmt.Println("Final distance: ", 5)
}
