package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type PizzaCake struct {
}

func (p *PizzaCake) getPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {

	pizzaCake := &PizzaCake{}
	fmt.Println("Prince of only pizza cake is ", pizzaCake.getPrice())

	pizzaWithCheese := &CheeseTopping{
		pizza: pizzaCake,
	}

	fmt.Println("Price of Pizza with cheese topping is", pizzaWithCheese.getPrice())

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of Pizza with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
