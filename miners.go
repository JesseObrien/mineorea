package mineorea 

import (
	"fmt"
	"time"
)

type modifier struct {
		
}

type miner struct {
	name string
	moveSpeed int16
	miningSpeed int16
	// Capacity is defined in meters cubed (m^3) 
	storageCapacity float32
	inventory []resource
}

func (m *miner) mineUntilFull(r resource) {
	
	for m.currentStorageCapacity() > 0.3 {
		m.mine(r)
	}

	defer wg.Done()
}

func (m *miner) mine(r resource){
	
	// Extract some resource to see how much we get
	er := r.extract()

	// @TODO: If resource is exhausted, move on, otherwise extract moar

	// Time to wait for "mining" depends on mass of unit mined 
	// This simulates some kind of real time taken
	time.Sleep(time.Duration(er.unitMass() * 1000) * time.Millisecond)

	fmt.Printf("\n%v mined %.2fkg of %v", m.name, er.unitMass(), er.getName())

	if m.currentStorageCapacity() - er.unitMass()> 0 {
		m.inventory = append(m.inventory, er)
	}
} 

func (m *miner) totalInventoryMass() float32 {
	var mass float32
	mass = 0.0
	
	for _, r := range m.inventory {
		mass += r.unitMass() 
	}

	return mass
}

func (m *miner) currentStorageCapacity() float32 {
	return m.storageCapacity - m.totalInventoryMass()
}


