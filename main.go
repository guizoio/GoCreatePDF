package main

import (
	"CreateFilePDF/src/cmd"
	"CreateFilePDF/src/configs"
	"CreateFilePDF/src/configs/database"
	"CreateFilePDF/src/entity"
	"CreateFilePDF/src/generator"
	"CreateFilePDF/src/infra/adapters/gorm/repository"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	configs.LoadEnv()
}

func main() {
	config := database.Config{
		Hostname: os.Getenv("DB_HOST_LOCAL"),
		Port:     os.Getenv("DB_PORT"),
		UserName: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	db := database.InitGorm(&config)

	cmdMakeMigrations := &cobra.Command{
		Use:   "MakeMigrations",
		Short: "Run MakeMigrations",
		Run: func(cli *cobra.Command, args []string) {
			makeMigration := cmd.NewDatabaseMakeMigrations(db)
			makeMigration.MakeMigrations()
		},
	}

	cmdHttpServer := &cobra.Command{
		Use:   "httpserver",
		Short: "Run httpserver",
		Run: func(cli *cobra.Command, args []string) {
			fmt.Println("httpserver")
		},
	}

	createPdf := &cobra.Command{
		Use:   "createPdf",
		Short: "Run createPdf",
		Run: func(cli *cobra.Command, args []string) {
			fmt.Println("init convert pdf")

			repositoryCreate := repository.NewCreateRepository(db)

			c := generator.NewCreatePDF(
				mockPeople().FilePDF,
				mockPeople().FileIMG,
				mockPeople().People,
				mockPeople().Company,
				repositoryCreate,
			)

			err := c.CreatePDF()
			if err != nil {
				fmt.Println("mockPeople:", err)
			}

			c = generator.NewCreatePDF(
				mockCompany().FilePDF,
				mockCompany().FileIMG,
				mockCompany().People,
				mockCompany().Company,
				repositoryCreate,
			)
			err = c.CreatePDF()
			if err != nil {
				fmt.Println("mockCompany:", err)
			}
		},
	}

	var rootCmd = &cobra.Command{Use: "APP"}
	rootCmd.AddCommand(cmdMakeMigrations)
	rootCmd.AddCommand(cmdHttpServer)
	rootCmd.AddCommand(createPdf)
	rootCmd.Execute()

}

func mockPeople() generator.CreatePDF {
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

func mockCompany() generator.CreatePDF {
	return generator.CreatePDF{
		FilePDF: 2,
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
				Address:    "Estrada Vicinal para Joao Batista Merlin",
				Number:     "582",
				District:   "Jardim Italia",
				City:       "Sao Joao da Boa Vista",
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
