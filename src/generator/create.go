package generator

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator/assemble_func"
	"errors"
)

type CreatePDF struct {
	FilePDF int64
	FileIMG string
	People  entity.People
	Company entity.Company
}

func NewCreatePDF(FilePDF int64, FileIMG string, Message entity.People, Company entity.Company) *CreatePDF {
	return &CreatePDF{FilePDF, FileIMG, Message, Company}
}

func (c *CreatePDF) CreatePDF() error {
	if c.FilePDF == 1 {
		return c.convertPdfPeoble()
	} else if c.FilePDF == 2 {
		return c.convertPdfCompany()
	} else {
		return errors.New("Error: file code pdf ")
	}

}

func (c *CreatePDF) convertPdfPeoble() error {
	pdf := assemble_func.Init()
	assemble_func.Logo(pdf, c.FileIMG)
	assemble_func.Title(pdf, "Registration Form")
	assemble_func.Body(pdf, c.People)
	return pdf.OutputFileAndClose("Registration.pdf")
}

func (c *CreatePDF) convertPdfCompany() error {
	pdf := assemble_func.InitCompany()
	assemble_func.LogoCompany(pdf, c.FileIMG)
	assemble_func.TitleCompany(pdf)
	assemble_func.BodyCompany(pdf, c.Company)
	return pdf.OutputFileAndClose("RegistrationCompany.pdf")
}
