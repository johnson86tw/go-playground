package observer

import "fmt"

type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

type observer interface {
	update(string)
	getID() string
}

type item struct {
	observers []observer
	name      string
	inStock   bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observers = append(i.observers, o)
}

func (i *item) deregister(o observer) {
	i.observers = removeFromSlice(i.observers, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observers {
		observer.update(i.name)
	}
}

// 這個是否應該寫成 receiver function ?
func removeFromSlice(observers []observer, observerToRemove observer) []observer {
	observerLen := len(observers)
	for i := range observers {
		observers[observerLen-1], observers[i] = observers[i], observers[observerLen-1]
		return observers[:observerLen-1]
	}
	return observers
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
