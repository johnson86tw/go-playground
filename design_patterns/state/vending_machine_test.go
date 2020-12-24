package state

import (
	"testing"
)

func TestVendingMachine(t *testing.T) {
	vm := newVendingMachine(1, 10)

	err := vm.insertMoney(12)
	if err == nil {
		t.Error("Vending machine should not be inserted money, but it can.")
	}

	err = vm.requestItem()
	if err != nil {
		t.Error(err.Error())
	}

	err = vm.insertMoney(10)
	if err != nil {
		t.Error(err.Error())
	}

	err = vm.dispenseItem()
	if err != nil {
		t.Error(err.Error())
	}

	err = vm.requestItem()
	if err.Error() == "item out of stock" {
		t.Error("Vending machine has no item, and it cannot be requested but it does.")
	}
}
