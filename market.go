package mineorea 

import (
	"fmt"
)

type market struct {
	name string
	orders map[string][]order
	orderId int64
}

func NewMarket() *market{
	fmt.Println("New market spawned.")
	return &market{orderId: 0, orders: make(map[string][]order)}
}

func (m *market) NextOrderId() int64 {
	m.orderId += 1
	return m.orderId 
}

type order interface {
	execute() bool	
}

type MarketSellOrder struct {
	quantity int32
	orderId int64
}

func NewMarketSellOrder(q int32, o int64) order {
	return &MarketSellOrder{quantity: q, orderId: o}
}

func (m MarketSellOrder) execute() bool {
	return true
}

func (m *market) submitMarketBuyOrder(quantity int32, resourceName string) {


	fmt.Println("Added market buy order.")
} 

func (m *market) submitMarketSellOrder(quantity int32, resourceName string) {

	m.orders[resourceName] = append(m.orders[resourceName], NewMarketSellOrder(quantity, m.NextOrderId()))
	fmt.Println("Added market sell order.")

}


