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

	bodyInfo(pdf, 20, 120, "Name", txt.Name)
	bodyInfo(pdf, 17, 135, "CPF", txt.CPF)
	bodyInfo(pdf, 15, 150, "RG", txt.RG)
	bodyInfo(pdf, 28, 165, "Birth Date", txt.BirthDate)
	bodyInfo(pdf, 20, 180, "Email", txt.Email)
	bodyInfo(pdf, 33, 195, "Code Postal", txt.CodePostal)
	bodyInfo(pdf, 26, 210, "Address", txt.Address)
	bodyInfo(pdf, 24, 225, "Number", txt.Number)
	bodyInfo(pdf, 23, 240, "District", txt.District)
	bodyInfo(pdf, 17, 255, "City", txt.City)
	bodyInfo(pdf, 19, 270, "State", txt.State)
	bodyInfo(pdf, 17, 285, "Cell", txt.Cell)
	bodyInfo(pdf, 20, 300, "Telephone", txt.Telephone)

}

func bodyInfo(pdf *gofpdf.Fpdf, x, h float64, title, name string) {
	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(5, 5)
	pdf.Cell(0, h, title+":")
	pdf.SetFont("Arial", "", 12)
	pdf.MoveTo(x, 5)
	pdf.Cell(0, h, name)
}
