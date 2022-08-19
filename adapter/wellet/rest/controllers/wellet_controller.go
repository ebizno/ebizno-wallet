package controllers

import (
	"errors"
	"net/http"

	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/rest/controllers/interfaces"
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/rest/presenter"
	"github.com/Paulo-Lopes-Estevao/e-bizno/usecase/dto"
	IwelletUse "github.com/Paulo-Lopes-Estevao/e-bizno/usecase/interfaces"
)

type welletController struct {
	welletUsecaseInterface IwelletUse.WelletUsecaseInterface
}

func NewWelletController(IwelletUsecase IwelletUse.WelletUsecaseInterface) interfaces.IwelletController {
	return &welletController{
		welletUsecaseInterface: IwelletUsecase,
	}
}

func (w welletController) CreateWellet(ctx presenter.WelletPresenterContext) error {
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	output, err := w.welletUsecaseInterface.AddWellet(&input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, output)
}

func (w welletController) GetBalanceWellet(ctx presenter.WelletPresenterContext) error {
	id := ctx.Param("id")
	output, err := w.welletUsecaseInterface.BalanceWellet(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, output.Balance)
}

func (w welletController) DebitWelletCliente(ctx presenter.WelletPresenterContext) error {
	iduser := ctx.Param("iduser")
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	output, err := w.welletUsecaseInterface.DebitWelletClienteOut(input.Balance, iduser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, output)
}

func (w welletController) CreditWelletCliente(ctx presenter.WelletPresenterContext) error {
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	output, err := w.welletUsecaseInterface.CreditWelletClienteInt(&input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, output)
}

func (w welletController) CreditWelletVendedor(ctx presenter.WelletPresenterContext) error {
	iduser := ctx.Param("iduser")
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	output, err := w.welletUsecaseInterface.CreditWelletVendedorInt(input.Balance, iduser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, output)
}

func (w welletController) DebitWelletTarifeVendedor(ctx presenter.WelletPresenterContext) error {
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	output, err := w.welletUsecaseInterface.DebitWelletTarifeVendedorInt(&input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, output)
}

func (w welletController) CreditWelletTarifeEbizno(ctx presenter.WelletPresenterContext) error {
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	output, err := w.welletUsecaseInterface.CreditWelletTarifeEbiznoInt(&input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, output)
}

func (w welletController) DebitWelletIVAVendedor(ctx presenter.WelletPresenterContext) error {
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	output, err := w.welletUsecaseInterface.DebitWelletIVAVendedorOut(&input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, output)
}

func (w welletController) CreditWelletIVAEbizno(ctx presenter.WelletPresenterContext) error {
	var input dto.WelletDtoInput
	if err := ctx.Bind(&input); !errors.Is(err, nil) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	output, err := w.welletUsecaseInterface.CreditWelletIVAEbiznoInt(&input)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, output)
}
