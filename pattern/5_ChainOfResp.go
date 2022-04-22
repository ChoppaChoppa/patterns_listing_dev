package main

import "fmt"

func main() {
	receiver := Receiver{
		BankTransfer:   false,
		BTCTransfer:    true,
		PayPalTransfer: true,
	}

	bank := BankPaymentHandler{}
	btc := BTCTransferHandler{}
	paypal := PayPalTransferHandler{}
	bank.Payment.Successor = paypal
	paypal.Payment.Successor = btc

	bank.Handle(receiver)
}

// receiver

type Receiver struct {
	BankTransfer   bool
	BTCTransfer    bool
	PayPalTransfer bool
}

//

// Handler
type IPaymentHandler interface {
	Handle(receiver Receiver)
}

type PaymentHandler struct {
	Successor IPaymentHandler
}

type BankPaymentHandler struct {
	Payment PaymentHandler
}

func (e *BankPaymentHandler) Handle(receiver Receiver) {
	if receiver.BankTransfer {
		fmt.Println("выполняем банковский перевод")
	} else if e.Payment.Successor != nil {
		e.Payment.Successor.Handle(receiver)
	}
}

type BTCTransferHandler struct {
	Payment PaymentHandler
}

func (e BTCTransferHandler) Handle(receiver Receiver) {
	if receiver.BTCTransfer {
		fmt.Println("выполняем перевод через биткоины")
	} else if e.Payment.Successor != nil {
		e.Payment.Successor.Handle(receiver)
	}
}

type PayPalTransferHandler struct {
	Payment PaymentHandler
}

func (e PayPalTransferHandler) Handle(receiver Receiver) {
	if receiver.PayPalTransfer {
		fmt.Println("перевод через paypal")
	} else if e.Payment.Successor != nil {
		e.Payment.Successor.Handle(receiver)
	}
}
