package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const (
	popSize     = 5
	elitism     = 20
	mutation    = 0.1
	generations = 500
)

type cityT struct {
	name string
	x, y int
}

func (c *cityT) distance(t cityT) float64 {
	xDis := (c.x - t.x)
	yDis := (c.y - t.y)
	return math.Sqrt(float64((xDis * xDis) + (yDis * yDis)))
}

type routeT struct {
	route    []cityT
	distance float64
}

func (f *routeT) routeDistance() float64 {
	if f.distance == 0 {
		dist := 0.0
		for k, city := range f.route {
			from := city
			to := cityT{}
			if k+1 < len(f.route) {
				to = f.route[k+1]
			} else {
				to = f.route[0]
			}
			dist += from.distance(to)
		}
		f.distance = dist
	}
	return f.distance
}

func initCities() {
	for i := range cities {
		cities[i] = cityT{fmt.Sprintf("city%d", i), rand.Intn(100), rand.Intn(100)}
	}
}

const citiesNum = 10

var cities = make([]cityT, citiesNum)

func initialPopulation() []routeT {
	population := make([]routeT, popSize)
	for popi := 0; popi < popSize; popi++ {
		idx := rand.Perm(citiesNum)
		sample := make([]cityT, citiesNum)
		for i := range idx {
			sample[i] = cities[idx[i]]
		}
		population[popi] = routeT{route: sample}
	}
	return population
}

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
