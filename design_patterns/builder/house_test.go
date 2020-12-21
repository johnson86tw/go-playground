package builder

import (
	"testing"
)

func TestBuildHouse(t *testing.T) {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	if normalHouse.doorType != "wooden" || normalHouse.floorNum != 2 {
		t.Error("fail to build normal house")
	}

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	if iglooHouse.doorType != "snow" || iglooHouse.floorNum != 1 {
		t.Error("fail to build igloo house")
	}
}
