package interfaces

import "github.com/Paulo-Lopes-Estevao/e-bizno/usecase/dto"

type WelletUsecaseInterface interface {
	AddWellet(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error)
	BalanceWellet(IdUser string) (*dto.WelletDtoOutput, error)
	DebitWelletClienteOut(balance int, id_user string) (*dto.WelletDtoOutput, error)
	CreditWelletClienteInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error)
	CreditWelletVendedorInt(balance int, id_user string) (*dto.WelletDtoOutput, error)
	DebitWelletTarifeVendedorInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error)
	CreditWelletTarifeEbiznoInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error)
	DebitWelletIVAVendedorOut(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error)
	CreditWelletIVAEbiznoInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error)
}
