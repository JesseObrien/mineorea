package mineorea 

import (

)

type market struct {
	name string
	selling map[int][]resource
	buying map[int][]resource
}

func (m market) addBuyOrder(float32 price, int32 quantity) {

} 
