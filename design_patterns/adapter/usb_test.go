package adapter

import "testing"

func TestUsbAdapter(t *testing.T) {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
