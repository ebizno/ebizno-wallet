package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutDebitValueClient(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 400

	ou_wallet, err := in_wellet.DebitWelletCliente(300)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, ou_wallet.Balance, 100)
}

func TestInCreditiValueVendedor(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 200

	ou_wallet := in_wellet.CreditWelletVendedor(100)
	assert.Equal(t, ou_wallet.Balance, 300)
}

func TestOutDebitValueClientANDInCreditiValueVendedor(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 2000

	ou_wallet_CLI, _ := in_wellet.DebitWelletCliente(100)

	ou_wallet_VEN := in_wellet.CreditWelletVendedor(ou_wallet_CLI.Balance)
	assert.Equal(t, ou_wallet_VEN.Balance, 3900)
}

func TestOutDebitTarifeVendedor(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 1000

	ou_wallet, _ := in_wellet.DebitWelletTarifeVendedor(200000)
	assert.Equal(t, ou_wallet.Tarife, 20)
}

func TestInCreditTarifeEbizno(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 0

	ou_wallet := in_wellet.CreditWelletTarifeEbizno(200000)
	assert.Equal(t, ou_wallet.Tarife, 20)
}

func TestOutDebitIVAVendedor(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 0

	ou_wallet := in_wellet.DebitWelletIVAVendedor(200000)
	assert.Equal(t, ou_wallet.Iva, 14000)
}

func TestInCreditIVAEbizno(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 0

	ou_wallet := in_wellet.CreditWelletIVAEbizno(200000)
	assert.Equal(t, ou_wallet.Iva, 14000)
}

func TestOutDebitIVAVendedorANDInCreditIVAEbizno(t *testing.T) {
	in_wellet := NewWellet()
	in_wellet.IdWellet = "1"
	in_wellet.IdUser = "1"
	in_wellet.Balance = 10

	ou_wallet_VEN := in_wellet.DebitWelletIVAVendedor(200000)
	ou_wellet_EBI := ou_wallet_VEN.CreditWelletTarifeEbizno(ou_wallet_VEN.Debit)
	assert.Equal(t, ou_wellet_EBI.Iva, 0)
}
