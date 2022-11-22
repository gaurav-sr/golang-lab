package structs

import "testing"

func TestProduct(t *testing.T) {
	customer := Customer{
		product: Product{
			Name:     "CricketBall",
			Price:    float64(9.88),
			Currency: "Euro",
		},
		stock: int32(12),
	}

	totalPrice := customer.calculateTotalCost()
	if totalPrice != float64(118.56) {
		t.Errorf("Error in calculating price")
	}
}

func TestSell(t *testing.T) {
	customer := Customer{
		product: Product{Name: "Bat", Price: 15.99, Currency: "Euro"},
		stock:   int32(32),
	}
	customer.sell(10)
	if customer.stock != 22 {
		t.Errorf("Error selling the product")
	}
}
