package structs

import "fmt"

type Customer struct {
	product Product
	stock   int32
}

func (c Customer) calculateTotalCost() float64 {
	fmt.Println("Calculating cost...")
	totalCost := c.product.Price * float64(c.stock)
	return totalCost
}

func (c *Customer) sell(sellQuantity int) {
	if c.stock >= int32(sellQuantity) {
		c.stock = c.stock - int32(sellQuantity)
	}
}
