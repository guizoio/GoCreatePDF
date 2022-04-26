package generator

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator/assemble_func"
	"CreateFilePDF/src/infra/adapters/gorm/model"
	"CreateFilePDF/src/infra/adapters/gorm/repository"
	"encoding/json"
	"errors"
	"github.com/gofrs/uuid"
)

type CreatePDF struct {
	HeaderPDF  entity.HeadlerPDF
	People     entity.People
	Company    entity.Company
	Repository repository.FaceCreateRepository
}

func NewCreatePDF(HeaderPDF entity.HeadlerPDF, Message entity.People, Company entity.Company, Repository repository.FaceCreateRepository) CreatePDF {
	return CreatePDF{HeaderPDF, Message, Company, Repository}
}

func (c *CreatePDF) CreatePDF() error {
	if c.HeaderPDF.FilePDF == 1 {
		return c.convertPdfPeople()
	} else if c.HeaderPDF.FilePDF == 2 {
		return c.convertPdfCompany()
	} else {
		return errors.New("Error: file code pdf ")
	}

}

func (c *CreatePDF) convertPdfPeople() error {
	pdf := assemble_func.Init()
	assemble_func.Logo(pdf, c.HeaderPDF.FileIMG)
	assemble_func.Title(pdf, "Registration Form")
	assemble_func.Body(pdf, c.People)

	buffer, _ := json.Marshal(c.People)
	err := c.infoCreateDB(buffer, "PEOPLE")
	if err != nil {
		return err
	}

	return pdf.OutputFileAndClose("Registration.pdf")
}

func (c *CreatePDF) convertPdfCompany() error {
	pdf := assemble_func.InitCompany()
	assemble_func.LogoCompany(pdf, c.HeaderPDF.FileIMG)
	assemble_func.TitleCompany(pdf)
	assemble_func.BodyCompany(pdf, c.Company)

	buffer, _ := json.Marshal(c.Company)
	err := c.infoCreateDB(buffer, "COMPANY")
	if err != nil {
		return err
	}

	return pdf.OutputFileAndClose("RegistrationCompany.pdf")
}

func (c *CreatePDF) infoCreateDB(buffer []byte, TxtType string) error {
	textUUID, _ := uuid.NewV4()
	data := model.Create{
		ID:      textUUID.String(),
		Name:    TxtType,
		Content: buffer,
	}
	return c.Repository.Create(data)
}
