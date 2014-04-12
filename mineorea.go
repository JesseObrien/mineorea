package mineorea

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
	"github.com/manveru/faker"
)

var wg sync.WaitGroup
var gridX int = 10
var gridY int = 10

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type mineorea struct {
	miners int
}


func Server() *mineorea {
	return &mineorea{}
}

func newFaker() *faker.Faker {
	fake, _ := faker.New("en")
	fake.Rand = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return fake
}

func (m mineorea) Run() {

	nm := rand.Intn(5)

	fmt.Println("Resource Sim 0.0.1\n")
	
	miners := make([]*miner, nm)
	
	for i := range miners {
		miners[i] = Miner() 
		wg.Add(1)
		fmt.Printf("Miner %v away! Miner has %v kg Capacity.\n", miners[i].name, miners[i].currentStorageCapacity())
		// @TODO: Pass a grid and a market into the miners so they can mine the grid clean
		go miners[i].mineUntilFull(coal{})
	}

	wg.Wait()

	fmt.Println("\nMining complete!")
	
}

func seedGrid() [][]resource {
	grid := make([][]resource, gridX)

	fmt.Println("Seeding Grid")
	for x := 0; x < gridX; x++ {
		grid[x] = make([]resource, gridY)
		for y := 0; y < gridY; y++ {
			resource := rand.Intn(4)
			if resource > 1 {
				grid[x][y] = coal{} 
			} else {
				grid[x][y] = nil 
			}
		}
		fmt.Println(grid[x])
	}

	return grid
}
