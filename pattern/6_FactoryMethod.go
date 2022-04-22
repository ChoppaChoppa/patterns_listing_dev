package main

import "fmt"

func main() {

	dev := PanelDeveloper{}
	house2 := dev.Create()

	dev.IDeveloper = WoodDeveloper{}
	house := dev.Create()

	fmt.Println(house)
	fmt.Println(house2)
}

type IDeveloper interface {
	Create() IHouse
}

// Panel

type PanelDeveloper struct {
	IDeveloper
}

func (p *PanelDeveloper) Create() IHouse {
	return PanelHouse{}
}

//

//Wood

type WoodDeveloper struct {
	IDeveloper
}

func (w WoodDeveloper) Create() IHouse {
	return WoodHouse{}
}

//

// houses

type IHouse interface{}

type PanelHouse struct {
	Name string
	IHouse
}

func (p PanelHouse) PanelHouse() {
	fmt.Println("панельный дом")
}

type WoodHouse struct {
	IHouse
}

func (w WoodHouse) WoodHouse() {
	fmt.Println("деревянный дом")
}
