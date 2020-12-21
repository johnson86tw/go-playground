package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printGunDetails(ak47)
	printGunDetails(maverick)
}

func printGunDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
