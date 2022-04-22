package main

import "fmt"

func main() {
	structure := Bank{}
	structure.Add(&Person{
		FIO:       "Salim KHabitov",
		AccNumber: "1",
	})
	structure.Add(&Company{
		Name:      "wildberries",
		RegNumber: "test123",
		Number:    "123123",
	})
	structure.Accept(&HtmlVisitor{})
	structure.Accept(&XmlVisitor{})
}

type IAccount interface {
	Accept(visitor IVisitor)
}

type IVisitor interface {
	VisitPersonAcc(acc Person)
	VisitCompanyAcc(acc Company)
}

//html serialize

type HtmlVisitor struct {
	IVisitor
}

func (html *HtmlVisitor) VisitPersonAcc(acc Person) {
	result := "serialize to html" + acc.FIO + acc.AccNumber
	fmt.Println(result)
}

func (html *HtmlVisitor) VisitCompanyAcc(acc Company) {
	result := "serialize to html" + acc.Number + acc.Name
	fmt.Println(result)
}

//

//xml serialize

type XmlVisitor struct {
	IVisitor
}

func (xml *XmlVisitor) VisitPersonAcc(acc Person) {
	result := "serialize to xml" + acc.FIO + acc.AccNumber
	fmt.Println(result)
}

func (xml *XmlVisitor) VisitCompanyAcc(acc Company) {
	result := "serialize to xml" + acc.Number + acc.Name
	fmt.Println(result)
}

//

//Bank

type Bank struct {
	Accounts []IAccount
}

func (b *Bank) Add(acc IAccount) {
	b.Accounts = append(b.Accounts, acc)
}

func (b *Bank) Accept(visitor IVisitor) {
	for _, acc := range b.Accounts {
		acc.Accept(visitor)
	}
}

//

// Person

type Person struct {
	FIO       string
	AccNumber string
}

func (p *Person) Accept(visitor IVisitor) {
	visitor.VisitPersonAcc(*p)
}

//

// Company

type Company struct {
	Name      string
	RegNumber string
	Number    string
}

func (c *Company) Accept(visitor IVisitor) {
	visitor.VisitCompanyAcc(*c)
}

//
