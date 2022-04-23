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
		Message: "Hello Word",
	}
}
