package mineorea 

import (
	"math/rand"
)

type resource interface {
	getName() string
	extract() resource 
	// Density is defined in kg/m^3
	density() float32 
	unitMass() float32
}

type coal struct {
	mass float32
	name string
}

func (c coal) extract() resource {
	return &coal{name: "Coal", mass: rand.Float32()}	
}

func (c coal) unitMass() float32 {
	return c.mass
}

func (c coal) density() float32 {
	return 774.20 
}

func (c coal) getName() string {
	return c.name
}

type copper struct {}

func (c copper) density() float32 {
	return 896.00
}

