package routes

import (
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/rest/controllers/interfaces"
	"github.com/labstack/echo/v4"
)

func WelletRouter(e *echo.Echo, Icontroller interfaces.IwelletController) *echo.Echo {
	e.POST("/wellet/create", func(context echo.Context) error {
		return Icontroller.CreateWellet(context)
	})
	e.GET("wallet/balance/:id", func(context echo.Context) error { return Icontroller.GetBalanceWellet(context) })
	e.PUT("wallet/debitcliente/:iduser", func(context echo.Context) error {
		return Icontroller.DebitWelletCliente(context)
	})
	e.POST("wellet/creditcliente", func(context echo.Context) error {
		return Icontroller.CreditWelletCliente(context)
	})
	e.PUT("wallet/creditvendedor/:iduser", func(context echo.Context) error {
		return Icontroller.CreditWelletVendedor(context)
	})
	e.POST("/wellet/debittarifevendedor", func(context echo.Context) error {
		return Icontroller.DebitWelletTarifeVendedor(context)
	})
	e.POST("/wellet/credittarifeebizno", func(context echo.Context) error {
		return Icontroller.CreditWelletTarifeEbizno(context)
	})
	e.POST("/wellet/debitivavendedor", func(context echo.Context) error {
		return Icontroller.DebitWelletIVAVendedor(context)
	})
	e.POST("/wellet/creditivaebizno", func(context echo.Context) error {
		return Icontroller.CreditWelletIVAEbizno(context)
	})

	return e
}
