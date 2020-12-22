package decorator

type pizza interface {
	getPrice() int
}

type peppyPaneer struct{}

func (p *peppyPaneer) getPrice() int {
	return 20
}

type veggeMania struct{}

func (v *veggeMania) getPrice() int {
	return 15
}

// topping
type cheeseTopping struct {
	pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

type tomatoTopping struct {
	pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
