package generator

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator/assemble_func"
)

type CreatePDF struct {
	FilePDF string
	FileIMG string
	Message entity.People
}

func NewCreatePDF(FilePDF string, FileIMG string, Message entity.People) *CreatePDF {
	return &CreatePDF{FilePDF, FileIMG, Message}
}

func (c *CreatePDF) Convert() error {
	pdf := assemble_func.Init()
	assemble_func.Logo(pdf, c.FileIMG)
	assemble_func.Title(pdf)
	assemble_func.Body(pdf, c.Message)
	return pdf.OutputFileAndClose(c.FilePDF + ".pdf")
}
