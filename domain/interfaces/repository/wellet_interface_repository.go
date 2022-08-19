package repository

type WelletInterfaceRepository interface {
	CreateWellet(id_user string, balance int) error
	DebitWelletCliente(balance, debit int, id_user string) error
	CreditWelletCliente(balance int) error
	CreditWelletVendedor(balance, credit int, id_user string) error
	DebitWelletTarifeVendedor(debitamount, creditamount, debit, tarife int) error
	CreditWelletTarifeEbizno(credit, tarife int) error
	DebitWelletIVAVendedor(debit, iva int) error
	CreditWelletIVAEbizno(credit, iva int) error
	FindWelletBYIdUser(id_user string) (interface{}, error)
}
