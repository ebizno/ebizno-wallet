package system

import (
	"encoding/json"

	"github.com/Paulo-Lopes-Estevao/e-bizno/domain/entities"
	"github.com/Paulo-Lopes-Estevao/e-bizno/domain/interfaces/repository"
	"github.com/Paulo-Lopes-Estevao/e-bizno/usecase/dto"
	"github.com/Paulo-Lopes-Estevao/e-bizno/usecase/interfaces"
)

type welletUsecaseSystem struct {
	welletRepository repository.WelletInterfaceRepository
}

func NewWelletUsecaseSystem(welletRepo repository.WelletInterfaceRepository) interfaces.WelletUsecaseInterface {
	return &welletUsecaseSystem{
		welletRepository: welletRepo,
	}
}

func (w *welletUsecaseSystem) AddWellet(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error) {
	in_wellet := entities.NewWellet()
	in_wellet.IdUser = welletInput.UserID
	in_wellet.Balance = welletInput.Balance

	err := w.welletRepository.CreateWellet(in_wellet.IdUser, in_wellet.Balance)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		WelletID: in_wellet.IdWellet,
		UserID:   in_wellet.IdUser,
		Balance:  in_wellet.Balance,
	}, nil
}

func (w *welletUsecaseSystem) BalanceWellet(IdUser string) (*dto.WelletDtoOutput, error) {
	atual_balance, err := w.welletRepository.FindWelletBYIdUser(IdUser)
	if err != nil {
		return nil, err
	}
	dt_wellet := &dto.WelletDtoOutput{}
	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, dt_wellet)
	return dt_wellet, nil
}

func (w *welletUsecaseSystem) DebitWelletClienteOut(balance int, id_user string) (*dto.WelletDtoOutput, error) {
	dt_wellet := dto.WelletDtoOutput{}
	atual_balance, err := w.welletRepository.FindWelletBYIdUser(id_user)
	if err != nil {
		return nil, err
	}
	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, &dt_wellet)

	wellet := entities.Wellet{
		IdUser:  dt_wellet.UserID,
		Balance: dt_wellet.Balance,
	}
	in_wellet, err := wellet.DebitWelletCliente(balance)
	if err != nil {
		return nil, err
	}

	err = w.welletRepository.DebitWelletCliente(in_wellet.Balance, in_wellet.Debit, wellet.IdUser)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		UserID:  wellet.IdUser,
		Balance: in_wellet.Balance,
		Debit:   in_wellet.Debit,
	}, nil

}
func (w *welletUsecaseSystem) CreditWelletClienteInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error) {
	dt_wellet := entities.Wellet{}

	atual_balance, err := w.welletRepository.FindWelletBYIdUser(welletInput.UserID)
	if err != nil {
		return nil, err
	}

	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, &dt_wellet)

	in_wellet := dt_wellet.CreditWelletCliente()

	err = w.welletRepository.CreditWelletCliente(dt_wellet.Balance)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		Cancel: in_wellet.Cancel,
	}, nil
}
func (w *welletUsecaseSystem) CreditWelletVendedorInt(balance int, id_user string) (*dto.WelletDtoOutput, error) {
	dt_wellet := dto.WelletDtoOutput{}

	atual_balance, err := w.welletRepository.FindWelletBYIdUser(id_user)
	if err != nil {
		return nil, err
	}

	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, &dt_wellet)

	wellet := entities.Wellet{
		IdUser:  dt_wellet.UserID,
		Balance: dt_wellet.Balance,
	}

	in_wellet := wellet.CreditWelletVendedor(balance)

	err = w.welletRepository.CreditWelletVendedor(in_wellet.Balance, in_wellet.Credit, wellet.IdUser)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		UserID:  in_wellet.IdUser,
		Balance: in_wellet.Balance,
		Credit:  in_wellet.Credit,
	}, nil
}
func (w *welletUsecaseSystem) DebitWelletTarifeVendedorInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error) {
	dt_wellet := entities.Wellet{}

	atual_balance, err := w.welletRepository.FindWelletBYIdUser(welletInput.UserID)
	if err != nil {
		return nil, err
	}

	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, &dt_wellet)

	in_wellet, _ := dt_wellet.DebitWelletTarifeVendedor(welletInput.Balance)

	err = w.welletRepository.DebitWelletTarifeVendedor(in_wellet.DebitAmount, in_wellet.CreditAmount, in_wellet.Debit, in_wellet.Tarife)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		Balance:      in_wellet.Balance,
		CreditAmount: in_wellet.CreditAmount,
		DebitAmount:  in_wellet.DebitAmount,
		Tarife:       in_wellet.Tarife,
	}, nil
}
func (w *welletUsecaseSystem) CreditWelletTarifeEbiznoInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error) {
	dt_wellet := entities.Wellet{}

	atual_balance, err := w.welletRepository.FindWelletBYIdUser(welletInput.UserID)
	if err != nil {
		return nil, err
	}

	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, &dt_wellet)

	in_wellet := dt_wellet.CreditWelletTarifeEbizno(welletInput.Balance)

	err = w.welletRepository.CreditWelletTarifeEbizno(in_wellet.Credit, in_wellet.Tarife)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		Balance: in_wellet.Balance,
		Credit:  in_wellet.Credit,
		Tarife:  in_wellet.Tarife,
	}, nil
}
func (w *welletUsecaseSystem) DebitWelletIVAVendedorOut(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error) {
	dt_wellet := entities.Wellet{}

	atual_balance, err := w.welletRepository.FindWelletBYIdUser(welletInput.UserID)
	if err != nil {
		return nil, err
	}

	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, &dt_wellet)

	in_wellet := dt_wellet.DebitWelletIVAVendedor(welletInput.Balance)

	err = w.welletRepository.DebitWelletIVAVendedor(in_wellet.Debit, in_wellet.Iva)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		Balance: in_wellet.Balance,
		Credit:  in_wellet.Credit,
		Iva:     in_wellet.Iva,
	}, nil
}
func (w *welletUsecaseSystem) CreditWelletIVAEbiznoInt(welletInput *dto.WelletDtoInput) (*dto.WelletDtoOutput, error) {
	dt_wellet := entities.Wellet{}

	atual_balance, err := w.welletRepository.FindWelletBYIdUser(welletInput.UserID)
	if err != nil {
		return nil, err
	}

	productJson, _ := json.Marshal(atual_balance)
	json.Unmarshal(productJson, &dt_wellet)

	in_wellet := dt_wellet.CreditWelletIVAEbizno(welletInput.Balance)

	err = w.welletRepository.CreditWelletIVAEbizno(in_wellet.Credit, in_wellet.Iva)
	if err != nil {
		return nil, err
	}
	return &dto.WelletDtoOutput{
		Balance: in_wellet.Balance,
		Credit:  in_wellet.Credit,
		Iva:     in_wellet.Iva,
	}, nil
}
