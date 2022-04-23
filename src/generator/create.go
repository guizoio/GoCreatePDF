package generator

import (
	"CreateFilePDF/src/generator/assemble_func"
)

type CreatePDF struct {
	FilePDF string
	FileIMG string
	Message string
}

func NewCreatePDF(FilePDF string, FileIMG string, Message string) *CreatePDF {
	return &CreatePDF{FilePDF, FileIMG, Message}
}

func (c *CreatePDF) Convert() error {
	pdf := assemble_func.Init()
	assemble_func.Logo(pdf, c.FileIMG)
	assemble_func.Title(pdf)
	assemble_func.Body(pdf, c.Message)
	return pdf.OutputFileAndClose(c.FilePDF + ".pdf")
}
