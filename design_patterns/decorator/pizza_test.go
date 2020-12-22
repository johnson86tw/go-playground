package decorator

import "testing"

func TestPizzaTopping(t *testing.T) {
	veggiePizza := &veggeMania{}

	veggiePizzaWithCheese := &cheeseTopping{
		pizza: veggiePizza,
	}

	veggiePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	if veggiePizzaWithCheeseAndTomato.getPrice() != 32 {
		t.Fatal("The price of veggie pizza with cheese and tomato is wrong.")
	}

}
