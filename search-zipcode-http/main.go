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
	http.HandleFunc("/", SearchZipcodeHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipcodeParam := r.URL.Query().Get("zipcode")
	if zipcodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	zipcode, err := SearchZipcode(zipcodeParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zipcode)
}

func SearchZipcode(zipcode string) (*ViaCEP, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipcode)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição para ViaCep %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao recuperar a response de ViaCep %v\n", err)
		return nil, err
	}

	var c ViaCEP
	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao transformar o retorno em struct ViaCEP %v\n", err)
		return nil, err
	}

	return &c, nil
}
