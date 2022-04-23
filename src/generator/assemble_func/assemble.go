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

func Title(pdf *gofpdf.Fpdf, txt string) {
	pdf.SetTitle("Registration Form", false)
	pdf.SetAuthor("ZSystem", false)

	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(1, 4.5)
	pdf.SetFillColor(200, 220, 255)
	pdf.CellFormat(0, 1, txt, "1", 1, "C", true, 0, "")
}

func Body(pdf *gofpdf.Fpdf, txt entity.People) {

	bodyInfo(pdf, 4.0, 2.7, "Name", txt.Name)
	bodyInfo(pdf, 4.0, 4.2, "CPF", txt.CPF)
	bodyInfo(pdf, 4.0, 5.7, "RG", txt.RG)
	bodyInfo(pdf, 4.0, 7.2, "Birth Date", txt.BirthDate)
	bodyInfo(pdf, 4.0, 8.7, "Email", txt.Contact.Email)
	bodyInfo(pdf, 4.0, 10.2, "Code Postal", txt.Address.CodePostal)
	bodyInfo(pdf, 4.0, 11.7, "Address", txt.Address.Address)
	bodyInfo(pdf, 4.0, 13.2, "Number", txt.Address.Number)
	bodyInfo(pdf, 4.0, 14.7, "District", txt.Address.District)
	bodyInfo(pdf, 4.0, 16.2, "City", txt.Address.City)
	bodyInfo(pdf, 4.0, 17.7, "State", txt.Address.State)
	bodyInfo(pdf, 4.0, 19.2, "Cell", txt.Contact.Cell)
	bodyInfo(pdf, 4.0, 20.7, "Telephone", txt.Contact.Telephone)

}

func bodyInfo(pdf *gofpdf.Fpdf, x, h float64, title, name string) {
	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(0.9, 5)
	pdf.Cell(0, h, title+":")
	pdf.SetFont("Arial", "", 12)
	pdf.MoveTo(x, 5)
	pdf.Cell(0, h, name)
}
