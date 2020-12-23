package observer

import "testing"

func TestObserver(t *testing.T) {
	shirtItem := newItem("Nike Shirt")
	observer1 := &customer{id: "a"}
	observer2 := &customer{id: "b"}

	shirtItem.register(observer1)
	shirtItem.register(observer2)

	shirtItem.updateAvailability()

	if len(shirtItem.observers) != 2 {
		t.Error("observers number is wrong")
	}

	shirtItem.deregister(observer1)

	if shirtItem.observers[0].getID() != observer2.getID() {
		t.Error("Fail to deregister observer1")
	}

}
