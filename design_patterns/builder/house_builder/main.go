package main

import "fmt"

// https://golangbyexample.com/builder-pattern-golang/

// When to use
// 1. 當要建構一個複雜或多步驟建構的物件時
// 2. 當有不同版本的產品要做出來時，可以實踐同一個界面來產出不同版本的產品
// 3. 當需要確保產品是完整而不是建構到一半時

type director struct {
	b builder
}

func newDirector(b builder) *director {
	return &director{b}
}

func (d *director) setBuilder(b builder) {
	d.b = b
}

func (d *director) buildHouse() house {
	d.b.setDoorType()
	d.b.setFloorNum()
	return d.b.getHouse()
}

type builder interface {
	setDoorType()
	setFloorNum()
	getHouse() house
}

func getBuilder(builderType string) builder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	default:
		return nil
	}
}

type house struct {
	doorType string
	floorNum int
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("normalHouse: %+v\n", normalHouse)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("iglooHouse: %+v\n", iglooHouse)
}

type normalBuilder struct {
	house
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "wooden"
}

func (b *normalBuilder) setFloorNum() {
	b.floorNum = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType: b.doorType,
		floorNum: b.floorNum,
	}
}

type iglooBuilder struct {
	house
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "snow"
}

func (b *iglooBuilder) setFloorNum() {
	b.floorNum = 1
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType: b.doorType,
		floorNum: b.floorNum,
	}
}
