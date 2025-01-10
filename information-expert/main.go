package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type DiscountCalculator struct{}

func (d *DiscountCalculator) Calculate(product Product, discountPercentage float64) float64 {
	return product.Price * (discountPercentage / 100)
}

// Checkout is responsible for processing the checkout
type Checkout struct {
	products []Product
	total    float64
}

func (c *Checkout) AddProduct(product Product) {
	c.products = append(c.products, product)
	c.total += product.Price
}

func (c *Checkout) GetTotal() float64 {
	return c.total
}

func (c *Checkout) PrintSummary(discountPercentage float64) {
	fmt.Println("Checckout Summary: ")

	for _, product := range c.products {
		fmt.Printf("- %s:$%.2f\n", product.Name, product.Price)
	}

	if discountPercentage > 0 {
		discountCalculator := &DiscountCalculator{}
		discountAmount := discountCalculator.Calculate((c.products[0]), discountPercentage) //say discount applies on first product.
		netPrice := c.total - discountAmount
		fmt.Printf("Discount: $%.2f\n", discountAmount)
		fmt.Printf("Total After Discount: $%.2f\n", netPrice)
	} else {
		fmt.Printf("Total: $%.2f\n", c.total)
	}
}

func main() {
	product1 := Product{Name: "Mobile", Price: 1000.00}
	product2 := Product{Name: "Laptop", Price: 2000.00}

	checkout := &Checkout{}
	checkout.AddProduct(product1)
	checkout.AddProduct(product2)

	//Print Checkout Summary
	checkout.PrintSummary(10) // 10% discount

}
