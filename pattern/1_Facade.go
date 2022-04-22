package main

import "fmt"

func main() {
	te := TextEditor{}
	c := Compiler{}
	r := run{}
	ide := IDEFacade{
		editor:  te,
		compile: c,
		run:     r,
	}
	p := Programmer{}
	p.CreateApp(ide)
}

type (
	TextEditor struct{}
	Compiler   struct{}
	run        struct{}
	IDEFacade  struct {
		editor  TextEditor
		compile Compiler
		run     run
	}
	Programmer struct{}
)

func (te TextEditor) CreateCode() {
	fmt.Println("write code")
}

func (te TextEditor) Save() {
	fmt.Println("save code")
}

func (c Compiler) Compile() {
	fmt.Println("compile app")
}

func (r run) Execute() {
	fmt.Println("execute app")
}

func (r run) Finish() {
	fmt.Println("finish app")
}

func (ide IDEFacade) Start() {
	ide.editor.CreateCode()
	ide.editor.Save()
	ide.compile.Compile()
	ide.run.Execute()
}

func (ide IDEFacade) Finish() {
	ide.run.Finish()
}

func (p Programmer) CreateApp(facade IDEFacade) {
	facade.Start()
	facade.Finish()
}
