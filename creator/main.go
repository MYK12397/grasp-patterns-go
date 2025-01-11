package main

import "fmt"

type Item struct {
	Name  string
	Price float64
}

func NewItem(name string, price float64) *Item {
	return &Item{Name: name, Price: price}
}

type Cart struct {
	items []*Item
}

func (c *Cart) AddItemToCart(name string, price float64) {
	item := &Item{Name: name, Price: price}
	c.items = append(c.items, item)

}

func (c *Cart) GetTotal() float64 {
	total := 0.0

	for _, item := range c.items {
		total += item.Price
	}
	return total
}

func (c *Cart) PrintCartInfo() {
	fmt.Println("Cart Items:")
	for _, item := range c.items {
		fmt.Printf("- %s: $%.2f\n", item.Name, item.Price)
	}
	fmt.Printf("Total: $%.2f\n", c.GetTotal())
}

func main() {
	c := &Cart{}

	c.AddItemToCart("Mobile", 1000)
	c.AddItemToCart("Laptop", 2000)

	c.PrintCartInfo()
}
