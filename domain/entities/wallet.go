package entities

import (
	"errors"
)

type Wellet struct {
	IdWellet     string
	IdUser       string
	Balance      int
	DebitAmount  int
	CreditAmount int
	Debit        int
	Credit       int
	Tarife       int
	Iva          int
	Cancel       bool
}

const (
	Iva    = 0.07
	Tarife = 0.010 / 100
)

// NewWellet Criar uma Wellet "Carteira"
func NewWellet() *Wellet {
	return &Wellet{}
}

func iva(amount int) int {
	balanceIVA := float64(amount) * Iva
	return int(balanceIVA)
}

func tarife(amount int) int {
	balanceTarife := float64(amount) * Tarife
	return int(balanceTarife)
}

func debit(w *Wellet, amount int) (int, error) {
	if w.Balance >= amount {
		return w.Balance - amount, nil
	}
	return amount, errors.New("Não tem saldo suficiente")
}

func credit(w *Wellet, amount int) int {
	creditTarife := w.Balance + amount
	return creditTarife
}

// #1 Transação Saida: Debitar valor na carteira do Cliente
func (w *Wellet) DebitWelletCliente(amount int) (*Wellet, error) {
	debitBalance, err := debit(w, amount)
	if err != nil {
		return nil, err
	}
	return &Wellet{
		Balance: debitBalance,
		Debit:   amount,
	}, nil
}

// #1.1 Transação Entrada: Creditar valor na carteira do Cliente #1 cancelamento ou devolução
func (w *Wellet) CreditWelletCliente() *Wellet {
	return &Wellet{
		Cancel: true,
	}
}

// #2 Transação Entrada: Creditar valor na carteira do Vendedor
func (w *Wellet) CreditWelletVendedor(amount int) *Wellet {
	creditBalance := credit(w, amount)
	return &Wellet{
		Balance: creditBalance,
		Credit:  amount,
	}
}

// #3 Tarifa Saida: Debitar tarifa no valor do Product do Vendedor
func (w *Wellet) DebitWelletTarifeVendedor(amount int) (*Wellet, error) {
	balanceTarife := tarife(amount)
	debitTarife, err := debit(w, balanceTarife)
	if err != nil {
		return nil, err
	}
	creditTarife := credit(w, balanceTarife)
	return &Wellet{
		DebitAmount:  debitTarife,
		CreditAmount: creditTarife,
		Debit:        amount,
		Tarife:       balanceTarife,
	}, nil
}

// #4 Tarifa Entrada: Creditar tarifa na carteira E-bizno
func (w *Wellet) CreditWelletTarifeEbizno(amount int) *Wellet {
	balanceTarife := tarife(amount)
	return &Wellet{
		Credit: amount,
		Tarife: balanceTarife,
	}
}

// #5 IVA Saida: Debitar IVA no valor do Product do Vendedor
func (w *Wellet) DebitWelletIVAVendedor(amount int) *Wellet {
	balanceIVA := iva(amount)
	return &Wellet{
		Debit: amount,
		Iva:   balanceIVA,
	}
}

// #6 IVA Entrada: Creditar IVA na carteira E-bizno
func (w *Wellet) CreditWelletIVAEbizno(amount int) *Wellet {
	balanceIVA := iva(amount)
	return &Wellet{
		Credit: amount,
		Iva:    balanceIVA,
	}
}
