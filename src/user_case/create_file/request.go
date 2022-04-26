package create_file

import "CreateFilePDF/src/entity"

type RequestCreatePDF struct {
	FilePDF    int64  `json:"file_pdf"`
	FileIMG    string `json:"file_img"`
	Name       string `json:"name"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
	BirthDate  string `json:"birth_date"`
	CodePostal string `json:"code_postal"`
	Address    string `json:"address"`
	Number     string `json:"number"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	Email      string `json:"email"`
	Cell       string `json:"cell"`
	Telephone  string `json:"telephone"`
}

func (r *RequestCreatePDF) ToDomain() entity.People {
	return entity.People{
		Name:      r.Name,
		CPF:       r.CPF,
		RG:        r.RG,
		BirthDate: r.BirthDate,
		Address: entity.Address{
			CodePostal: r.CodePostal,
			Address:    r.Address,
			Number:     r.Number,
			District:   r.District,
			City:       r.City,
			State:      r.State,
		},
		Contact: entity.Contact{
			Email:     r.Email,
			Cell:      r.Cell,
			Telephone: r.Telephone,
		},
	}
}
