package interfaces

import "github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/rest/presenter"

type IwelletController interface {
	CreateWellet(ctx presenter.WelletPresenterContext) error
	GetBalanceWellet(ctx presenter.WelletPresenterContext) error
	DebitWelletCliente(ctx presenter.WelletPresenterContext) error
	CreditWelletCliente(ctx presenter.WelletPresenterContext) error
	CreditWelletVendedor(ctx presenter.WelletPresenterContext) error
	DebitWelletTarifeVendedor(ctx presenter.WelletPresenterContext) error
	CreditWelletTarifeEbizno(ctx presenter.WelletPresenterContext) error
	DebitWelletIVAVendedor(ctx presenter.WelletPresenterContext) error
	CreditWelletIVAEbizno(ctx presenter.WelletPresenterContext) error
}
