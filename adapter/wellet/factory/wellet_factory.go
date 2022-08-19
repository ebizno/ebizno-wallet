package factory

import (
	"github.com/Paulo-Lopes-Estevao/e-bizno/adapter/wellet/infrastructure/database/repository"
	repo "github.com/Paulo-Lopes-Estevao/e-bizno/domain/interfaces/repository"
	"github.com/jinzhu/gorm"
)

type welletFactory struct {
	Db *gorm.DB
}

func NewWelletFactoryRepository(db *gorm.DB) *welletFactory {
	return &welletFactory{
		Db: db,
	}
}

func (w welletFactory) FactoryRepositoryWellet() repo.WelletInterfaceRepository {
	return repository.NewWelletRepository(w.Db)
}
