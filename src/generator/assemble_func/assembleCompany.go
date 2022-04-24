package assemble_func

import (
	"CreateFilePDF/src/entity"
	"github.com/jung-kurt/gofpdf"
)

func InitCompany() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "cm", "A4", "")
	pdf.AddPage()
	return pdf
}

func LogoCompany(pdf *gofpdf.Fpdf, NameImg string) {
	pdf.Image(NameImg, 0.9, 0.5, 0, 3.65,
		false, "", 0, "http://www.fpdf.org")
}

func TitleCompany(pdf *gofpdf.Fpdf) {
	pdf.SetTitle("Registration Form", false)
	pdf.SetAuthor("ZSystem", false)

	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(1, 4.5)
	pdf.SetFillColor(200, 220, 255)
	pdf.CellFormat(0, 1, "Registration Form Company", "1", 1, "C", true, 0, "")
}

func BodyCompany(pdf *gofpdf.Fpdf, txt entity.Company) {

	bodyInfoCompany(pdf, 4.0, 2.7, "Name", txt.Name)
	bodyInfoCompany(pdf, 4.0, 4.2, "CNPJ", txt.CNPJ)
	bodyInfoCompany(pdf, 4.0, 5.7, "Registration", txt.StateRegistration)
	bodyInfoCompany(pdf, 4.0, 7.2, "Opening Date", txt.OpeningDate)
	bodyInfoCompany(pdf, 4.0, 8.7, "Site", txt.Site)
	bodyInfoCompany(pdf, 4.0, 10.2, "Email", txt.Contact.Email)
	bodyInfoCompany(pdf, 4.0, 11.7, "Code Postal", txt.Address.CodePostal)
	bodyInfoCompany(pdf, 4.0, 13.2, "Address", txt.Address.Address)
	bodyInfoCompany(pdf, 4.0, 14.7, "Number", txt.Address.Number)
	bodyInfoCompany(pdf, 4.0, 16.2, "District", txt.Address.District)
	bodyInfoCompany(pdf, 4.0, 17.7, "City", txt.Address.City)
	bodyInfoCompany(pdf, 4.0, 19.2, "State", txt.Address.State)
	bodyInfoCompany(pdf, 4.0, 20.7, "Cell", txt.Contact.Cell)
	bodyInfoCompany(pdf, 4.0, 22.2, "Telephone", txt.Contact.Telephone)

}

func bodyInfoCompany(pdf *gofpdf.Fpdf, x, h float64, title, name string) {
	pdf.SetFont("Arial", "B", 12)
	pdf.MoveTo(0.9, 5)
	pdf.Cell(0, h, title+":")
	pdf.SetFont("Arial", "", 12)
	pdf.MoveTo(x, 5)
	pdf.Cell(0, h, name)
}
