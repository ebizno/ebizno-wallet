package repository

import (
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/infrastructure/database/entities"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type welletRepository struct {
	Db *gorm.DB
}

func NewWelletRepository(db *gorm.DB) *welletRepository {
	return &welletRepository{Db: db}
}

func (w *welletRepository) CreateWellet(id_user string, balance int) error {
	iduser_uuid, _ := uuid.FromString(id_user)
	wellet := entities.Wellet{UserID: iduser_uuid, Balance: balance}
	err := w.Db.Create(&wellet).Error
	return err
}

func (w *welletRepository) DebitWelletCliente(balance, debit int, id_user string) error {
	var wellet entities.Wellet
	iduser_uuid, _ := uuid.FromString(id_user)
	w.Db.Find(&wellet, "user_id=?", iduser_uuid)
	wellet.Balance = balance
	wellet.Debit = debit
	err := w.Db.Save(&wellet).Error
	return err
}
func (w *welletRepository) CreditWelletCliente(amount int) error {
	wellet := entities.Wellet{Balance: amount}
	err := w.Db.Model(&wellet).Update("balance", amount).Error
	return err
}

func (w *welletRepository) CreditWelletVendedor(balance, credit int, id_user string) error {
	var wellet entities.Wellet
	iduser_uuid, _ := uuid.FromString(id_user)
	w.Db.Find(&wellet, "user_id=?", iduser_uuid)
	wellet.Balance = balance
	wellet.Credit = credit
	err := w.Db.Save(&wellet).Error
	return err
}

func (w *welletRepository) DebitWelletTarifeVendedor(debitamount, creditamount, debit, tarife int) error {
	var wellet entities.Wellet
	w.Db.Find(&wellet)
	wellet.DebitAmount = debitamount
	wellet.CreditAmount = creditamount
	wellet.Debit = debit
	wellet.Tarife = tarife
	err := w.Db.Save(&wellet).Error
	return err
}

func (w *welletRepository) CreditWelletTarifeEbizno(credit, tarife int) error {
	var wellet entities.Wellet
	w.Db.Find(&wellet)
	wellet.Credit = credit
	wellet.Tarife = tarife
	err := w.Db.Save(&wellet).Error
	return err
}

func (w *welletRepository) DebitWelletIVAVendedor(debit, iva int) error {
	var wellet entities.Wellet
	w.Db.Find(&wellet)
	wellet.Debit = debit
	wellet.Iva = iva
	err := w.Db.Save(&wellet).Error
	return err
}

func (w *welletRepository) CreditWelletIVAEbizno(credit, iva int) error {
	var wellet entities.Wellet
	w.Db.Find(&wellet)
	wellet.Credit = credit
	wellet.Iva = iva
	err := w.Db.Save(&wellet).Error
	return err
}

func (w *welletRepository) FindWelletBYIdUser(id_user string) (interface{}, error) {
	var wellet entities.Wellet
	iduser_uuid, _ := uuid.FromString(id_user)
	err := w.Db.Find(&wellet, "user_id=?", iduser_uuid).Error
	return wellet, err
}
