package mineorea

import (
	"testing"
)

func TestAddBuyOrder(t *testing.T) {
	m := NewMarket()
	m.submitMarketSellOrder(20, "coal")

	l := len(m.orders["coal"])

	if l != 1 {
		t.Errorf("Number of coal orders: %v, want %v", l, 1)
	}

}
