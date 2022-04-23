package assemble_func

import (
	"CreateFilePDF/src/entity"
	"github.com/jung-kurt/gofpdf"
)

func Init() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "cm", "A4", "")
	pdf.AddPage()
	return pdf
}

func Logo(pdf *gofpdf.Fpdf, NameImg string) {
	pdf.Image(NameImg, 0.9, 0.5, 0, 3.65,
		false, "", 0, "http://www.fpdf.org")
}

func Title(pdf *gofpdf.Fpdf) {
	pdf.SetTitle("Registration Form", false)
	pdf.SetAuthor("ZSystem", false)

	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(1, 4.5)
	pdf.SetFillColor(200, 220, 255)
	pdf.CellFormat(0, 1, "Registration Form", "1", 1, "C", true, 0, "")
}

func Body(pdf *gofpdf.Fpdf, txt entity.People) {

	bodyInfo(pdf, 2.5, 2.7, "Name", txt.Name)
	bodyInfo(pdf, 2.5, 3.9, "CPF", txt.CPF)
	bodyInfo(pdf, 2.5, 5.1, "RG", txt.RG)
	bodyInfo(pdf, 3.5, 6.3, "Birth Date", txt.BirthDate)
	bodyInfo(pdf, 2.5, 7.5, "Email", txt.Email)
	bodyInfo(pdf, 4.0, 8.7, "Code Postal", txt.CodePostal)
	bodyInfo(pdf, 3.0, 9.9, "Address", txt.Address)
	bodyInfo(pdf, 3.0, 11.1, "Number", txt.Number)
	bodyInfo(pdf, 3.0, 12.3, "District", txt.District)
	bodyInfo(pdf, 2.5, 13.5, "City", txt.City)
	bodyInfo(pdf, 2.5, 14.7, "State", txt.State)
	bodyInfo(pdf, 2.5, 15.9, "Cell", txt.Cell)
	bodyInfo(pdf, 3.5, 17.1, "Telephone", txt.Telephone)

}

func bodyInfo(pdf *gofpdf.Fpdf, x, h float64, title, name string) {
	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(0.9, 5)
	pdf.Cell(0, h, title+":")
	pdf.SetFont("Arial", "", 12)
	pdf.MoveTo(x, 5)
	pdf.Cell(0, h, name)
}
