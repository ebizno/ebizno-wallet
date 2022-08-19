package entities

import (
	"time"

	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/user/infrastructure/database/entities"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

type Wellet struct {
	WelletID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"wellet_id"`
	UserID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();column:user_id"`
	Balance      int       `gorm:"column:balance"`
	DebitAmount  int       `gorm:"column:debit_amount"`
	CreditAmount int       `gorm:"column:credit_amount"`
	Debit        int       `gorm:"column:debit"`
	Credit       int       `gorm:"column:credit"`
	Tarife       int       `gorm:"column:tarife"`
	Iva          int       `gorm:"column:iva"`
	Cancel       bool      `gorm:"column:cancel"`
	Base
	User []entities.User
}
