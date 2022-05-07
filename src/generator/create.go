package generator

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator/assemble_func"
	"CreateFilePDF/src/infra/adapters/gorm/model"
	"CreateFilePDF/src/infra/adapters/gorm/repository"
	"context"
	"encoding/json"
	"errors"
	"github.com/gofrs/uuid"
	"os"
	"time"
)

type PublishMessage interface {
	PublishMessage(ctx context.Context, id string, message interface{}, topic string, headers map[string]string) error
}

type CreatePDF struct {
	HeaderPDF     entity.HeadlerPDF
	People        entity.People
	Company       entity.Company
	Repository    repository.FaceCreateRepository
	messageBroker PublishMessage
}

func NewCreatePDF(HeaderPDF entity.HeadlerPDF, Message entity.People, Company entity.Company, Repository repository.FaceCreateRepository, messageBroker PublishMessage) CreatePDF {
	return CreatePDF{HeaderPDF, Message, Company, Repository, messageBroker}
}

func (c *CreatePDF) CreatePDF(fileName string) error {
	if c.HeaderPDF.FilePDF == 1 {
		return c.convertPdfPeople(fileName)
	} else if c.HeaderPDF.FilePDF == 2 {
		return c.convertPdfCompany()
	} else {
		return errors.New("Error: file code pdf ")
	}

}

func (c *CreatePDF) convertPdfPeople(fileName string) error {
	pdf := assemble_func.Init()
	assemble_func.Logo(pdf, c.HeaderPDF.FileIMG)
	assemble_func.Title(pdf, "Registration Form")
	assemble_func.Body(pdf, c.People)

	buffer, _ := json.Marshal(c.People)
	err := c.infoCreateDB(buffer, "PEOPLE")
	if err != nil {
		return err
	}

	return pdf.OutputFileAndClose(fileName)
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
	var pe *entity.People
	json.Unmarshal(buffer, &pe)

	aa := entity.PublishMessageKafkaPeople{
		ID:   textUUID.String(),
		Date: time.Now(),
		People: entity.People{
			Name:      pe.Name,
			CPF:       pe.CPF,
			RG:        pe.RG,
			BirthDate: pe.BirthDate,
			Address: entity.Address{
				CodePostal: pe.Address.CodePostal,
				Address:    pe.Address.Address,
				Number:     pe.Address.Number,
				District:   pe.Address.District,
				City:       pe.Address.City,
				State:      pe.Address.State,
			},
			Contact: entity.Contact{
				Email:     pe.Contact.Email,
				Cell:      pe.Contact.Cell,
				Telephone: pe.Contact.Telephone,
			},
		},
	}

	go c.messageBroker.PublishMessage(context.TODO(), os.Getenv("KAFKA_GROUP"), aa, os.Getenv("KAFKA_TOPIC"), nil)

	return c.Repository.Create(data)
}
