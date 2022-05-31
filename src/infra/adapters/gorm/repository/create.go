package repository

import (
	"CreateFilePDF/src/infra/adapters/gorm/model"
	"gorm.io/gorm"
)

type FaceCreateRepository interface {
	Create(data model.Create) error
	UpdateStatus(data model.Create) error
}

type CreateRepository struct {
	database *gorm.DB
}

func NewCreateRepository(database *gorm.DB) *CreateRepository {
	return &CreateRepository{database}
}

func (c CreateRepository) Create(data model.Create) error {
	err := c.database.Create(data)
	if err != nil {
		return err.Error
	}
	return nil
}

func (c CreateRepository) UpdateStatus(data model.Create) error {
	err := c.database.Model(data).Where("id = ?", data.ID).Update("status", "complete")
	if err != nil {
		return err.Error
	}
	return nil
}
