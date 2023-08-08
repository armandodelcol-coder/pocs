package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, zipcode := range os.Args[1:] {
		fmt.Printf("Buscando o CEP: %s\n", zipcode)
		url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição %v\n", err)
			continue
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a resposta %v\n", err)
			continue
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta %v\n", err)
			continue
		}

		file, err := os.OpenFile("zipcode.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao abir o arquivo %v\n", err)
			continue
		}

		toWrite := fmt.Sprintf(
			"CEP: %s; Logradouro: %s, Bairro: %s, Cidade: %s, UF: %s\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf)
		if _, err := file.Write([]byte(toWrite)); err != nil {
			file.Close()
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo %v\n", err)
			continue
		}

		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fechar o arquivo %v\n", err)
		}
	}
}
