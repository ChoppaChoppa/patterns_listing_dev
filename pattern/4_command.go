package main

import "fmt"

func main() {
	tv := TV{}
	tvComm := TvOnCommand{tv}
	pult := Pult{}
	pult.SetCommand(&tvComm)
	pult.PressButton()
	pult.PressUndo()
}

type ICommand interface {
	Execute()
	Undo()
}

//TV

type TV struct{}

func (tv TV) On() {
	fmt.Println("tv on")
}

func (tv TV) Off() {
	fmt.Println("tv off")
}

//

//TvOnComm

type TvOnCommand struct {
	TV
}

func (tv *TvOnCommand) Execute() {
	tv.On()
}
func (tv *TvOnCommand) Undo() {
	tv.Off()
}

//

// Pult

type Pult struct {
	comm ICommand
}

func (p *Pult) SetCommand(command ICommand) {
	p.comm = command
}

func (p *Pult) PressButton() {
	p.comm.Execute()
}

func (p *Pult) PressUndo() {
	p.comm.Undo()
}

//
