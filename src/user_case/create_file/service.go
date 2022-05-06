package create_file

import (
	"CreateFilePDF/src/infra/adapters/gorm/repository"
)

type CreateService struct {
	repository repository.FaceCreateRepository
}

func NewCreateService(repository repository.FaceCreateRepository) *CreateService {
	return &CreateService{repository}
}

func (c *CreateService) CreatePDF() error {
	return nil
}
