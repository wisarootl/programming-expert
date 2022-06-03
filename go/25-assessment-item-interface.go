package main

import "fmt"

// Write your code here.

type Item interface {
	GetPrice() float64
}

type Gravel struct {
	lbs        float64
	pricePerLb float64
	grade      string
}

type Shovel struct {
	size  string
	price float64
}

func (g *Gravel) GetPrice() float64 {
	return g.lbs * g.pricePerLb
}

func (s *Shovel) GetPrice() float64 {
	return s.price
}

func CalculateOrderPrice(order []Item) float64 {
	totalPrice := 0.0

	for _, item := range order {
		totalPrice += item.GetPrice()
	}

	return totalPrice
}

func main() {
	items := []Item{
		&Gravel{4.0, 10.5, "fine"},
		&Shovel{"S", 25.0},
		&Gravel{3.0, 12.0, "coarse"},
		&Shovel{"M", 35.0},
	}
	expected := 138.0
	actual := CalculateOrderPrice(items)
	fmt.Println(expected, actual)
}