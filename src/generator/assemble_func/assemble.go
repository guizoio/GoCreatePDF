package assemble_func

import (
	"CreateFilePDF/src/entity"
	"github.com/jung-kurt/gofpdf"
)

func Init() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	return pdf
}

func Logo(pdf *gofpdf.Fpdf, NameImg string) {
	pdf.Image(NameImg, 5, 5, 200, 40,
		false, "", 0, "")
}

func Title(pdf *gofpdf.Fpdf) {
	pdf.SetTitle("Registration Form", false)
	pdf.SetAuthor("ZSystem", false)

	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(5, 50)
	pdf.SetFillColor(200, 220, 255)
	pdf.CellFormat(200, 8, "Registration Form", "1", 1, "C", true, 0, "")
	pdf.Ln(0)
}

func Body(pdf *gofpdf.Fpdf, txt entity.People) {
	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(5, 5)
	pdf.Cell(0, 120, "Name:")

	pdf.SetFont("Arial", "", 12)
	pdf.MoveTo(20, 5)
	pdf.Cell(0, 120, txt.Name)

	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(5, 5)
	pdf.Cell(0, 135, "Address:")

	pdf.SetFont("Arial", "", 12)
	pdf.MoveTo(25, 5)
	pdf.Cell(0, 135, txt.Address)
}
