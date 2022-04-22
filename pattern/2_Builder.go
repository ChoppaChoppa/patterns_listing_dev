package main

import (
	"fmt"
	"strings"
)

func main() {
	baker := Baker{}
	var builder BreadBuilder
	builder = RyeBreadBuilder{}

	ryeBread := baker.Bake(builder)
	fmt.Println(ryeBread.ToString())

	builder = WheatBreadBuilder{}
	wheatBread := baker.Bake(builder)
	fmt.Println(wheatBread.ToString())
}

// Builder

type BreadBuilder interface {
	SetFlour(Bread) Bread
	SetSalt(Bread) Bread
	SetAdditives(Bread) Bread
}

//

// BREAD
type (
	Flour struct {
		Sort string
	}
	Additives struct {
		Name string
	}
	Salt struct {
		Add bool
	}
)

type Bread struct {
	Flour     Flour
	Salt      Salt
	Additives Additives
}

func (b Bread) ToString() string {
	sb := strings.Builder{}

	if b.Flour != (Flour{}) {
		sb.WriteString(b.Flour.Sort + "\n")
	}
	if b.Salt != (Salt{}) {
		sb.WriteString("Соль \n")
	}
	if b.Additives != (Additives{}) {
		sb.WriteString("Добавки: " + b.Additives.Name + "\n")
	}

	return sb.String()
}

//

// Baker

type Baker struct{}

func (b Baker) Bake(builder BreadBuilder) Bread {
	bread := Bread{}
	bread = builder.SetFlour(bread)
	bread = builder.SetSalt(bread)
	bread = builder.SetAdditives(bread)
	return bread
}

//

// rye bread

type RyeBreadBuilder struct {
	BreadBuilder
}

func (rye RyeBreadBuilder) SetFlour(b Bread) Bread {
	b.Flour = Flour{Sort: "ржаная мука"}
	return b
}
func (rye RyeBreadBuilder) SetSalt(b Bread) Bread {
	b.Salt = Salt{Add: true}
	return b
}
func (rye RyeBreadBuilder) SetAdditives(b Bread) Bread {
	return b
}

//

// wheat bread

type WheatBreadBuilder struct {
	BreadBuilder
}

func (w WheatBreadBuilder) SetFlour(b Bread) Bread {
	b.Flour = Flour{Sort: "пшеничная мука высший сорт"}
	return b
}
func (w WheatBreadBuilder) SetSalt(b Bread) Bread {
	b.Salt = Salt{Add: true}
	return b
}
func (w WheatBreadBuilder) SetAdditives(b Bread) Bread {
	b.Additives = Additives{Name: "улучшитель хлебопекарный"}
	return b
}
