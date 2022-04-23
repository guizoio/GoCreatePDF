package main

import (
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator"
	"fmt"
)

func main() {
	fmt.Println("init convert pdf")

	c := generator.NewCreatePDF(
		mockAll().FilePDF,
		mockAll().FileIMG,
		mockAll().People,
		mockAll().Company,
	)

	err := c.CreatePDF()
	if err != nil {
		fmt.Println(err)
	}
}

func mockAll() generator.CreatePDF {
	return generator.CreatePDF{
		FilePDF: 1,
		FileIMG: "lim.png",
		People: struct {
			Name      string
			CPF       string
			RG        string
			BirthDate string
			Address   entity.Address
			Contact   entity.Contact
		}{
			Name:      "Isabela Carolina Bernardes",
			CPF:       "340.836.925-96",
			RG:        "25.366.824-4",
			BirthDate: "14/02/2002",
			Address: struct {
				CodePostal string
				Address    string
				Number     string
				District   string
				City       string
				State      string
			}{
				CodePostal: "89165-302",
				Address:    "Rua Felix Deeke Junior",
				Number:     "113",
				District:   "Budag",
				City:       "Rio do Sul",
				State:      "SP",
			},
			Contact: struct {
				Email     string
				Cell      string
				Telephone string
			}{
				Email:     "isabelacarolinabernardes@itatiaia.net",
				Cell:      "(47) 98463-2761",
				Telephone: "(47) 2637-5687",
			},
		},
		Company: struct {
			Name              string
			CNPJ              string
			StateRegistration string
			OpeningDate       string
			Site              string
			Address           entity.Address
			Contact           entity.Contact
		}{
			Name:              "Ryan e Kevin Lavanderia Ltda",
			CNPJ:              "89.814.507/0001-92",
			StateRegistration: "071.849.090.248",
			OpeningDate:       "14/05/2017",
			Site:              "www.ryanekevinlavanderialtda.com.br",
			Address: struct {
				CodePostal string
				Address    string
				Number     string
				District   string
				City       string
				State      string
			}{
				CodePostal: "13872-551",
				Address:    "Estrada Vicinal para João Batista Merlin",
				Number:     "582",
				District:   "Jardim Itália",
				City:       "São João da Boa Vista",
				State:      "SP",
			},
			Contact: struct {
				Email     string
				Cell      string
				Telephone string
			}{
				Email:     "seguranca@ryanekevinlavanderialtda.com.br",
				Cell:      "(19) 98528-5736",
				Telephone: "(19) 3524-5663",
			},
		},
	}
}
