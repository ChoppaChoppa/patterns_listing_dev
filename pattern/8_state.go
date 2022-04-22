package main

import "fmt"

func main() {
	water := &Water{State: &LiquidWaterState{}}
	water.Heat()
	water.Frost()
	water.Frost()
}

type IWaterState interface {
	Heat(*Water)
	Frost(*Water)
}

type Water struct {
	State IWaterState
}

func (w *Water) Heat() {
	w.State.Heat(w)
}
func (w *Water) Frost() {
	w.State.Frost(w)
}

type SolidWaterState struct {
	IWaterState
}

func (s *SolidWaterState) Heat(water *Water) {
	fmt.Println("превращаем лед в жидкость")
	water.State = &LiquidWaterState{}
}
func (s *SolidWaterState) Frost(water *Water) {
	fmt.Println("продолжаем заморозку льда")
}

type LiquidWaterState struct {
	IWaterState
}

func (l *LiquidWaterState) Heat(water *Water) {
	fmt.Println("Превращаем жидкость в пар")
	water.State = &GasWaterState{}
}
func (l *LiquidWaterState) Frost(water *Water) {
	fmt.Println("Превращаем жидкость в лед")
	water.State = &SolidWaterState{}
}

type GasWaterState struct {
	IWaterState
}

func (g *GasWaterState) Heat(water *Water) {
	fmt.Println("Повышаем температуру водяного пара")
}

func (g *GasWaterState) Frost(water *Water) {
	fmt.Println("Превращаем водяной пар в жидкость")
	water.State = &LiquidWaterState{}
}
