package main

import (
	"CreateFilePDF/src/generator"
	"fmt"
)

func main() {
	fmt.Println("init convert pdf")

	c := generator.NewCreatePDF(mock().FilePDF,
		mock().FileIMG, mock().Message)

	err := c.Convert()
	if err != nil {
		fmt.Println(err)
	}
}

func mock() generator.CreatePDF {
	return generator.CreatePDF{
		FilePDF: "Registration",
		FileIMG: "lim.png",
		Message: struct {
			Name       string
			CPF        string
			RG         string
			BirthDate  string
			Email      string
			CodePostal string
			Address    string
			Number     string
			District   string
			City       string
			State      string
			Cell       string
			Telephone  string
		}{
			Name:       "Isabela Carolina Bernardes",
			CPF:        "340.836.925-96",
			RG:         "25.366.824-4",
			BirthDate:  "14/02/2002",
			Email:      "isabelacarolinabernardes@itatiaia.net",
			CodePostal: "89165-302",
			Address:    "Rua Felix Deeke Junior",
			Number:     "113",
			District:   "Budag",
			City:       "Rio do Sul",
			State:      "SP",
			Cell:       "(47) 98463-2761",
			Telephone:  "(47) 2637-5687",
		},
	}
}
