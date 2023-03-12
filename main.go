package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCep struct {
	// used this tool to convert JSON to Go struct: https://mholt.github.io/json-to-go/
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

func createCEPFileAndAddText(cep string, text string) {
	folder := "cep-searches"
	file, err := os.Create(folder + "/" + cep + ".txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error trying created new .txt file: %v\n", err)
	}
	defer file.Close()
	file.WriteString(text)
}

func main() {
	for _, cep := range os.Args[1:] {
		url := "http://viacep.com.br/ws/" + cep + "/json/"

		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to make the request: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to read the results: %v\n", err)
		}

		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error trying to unmarshal the results: %v\n", err)
		}

		info := fmt.Sprintf(
			"CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s",
			data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf,
		)

		createCEPFileAndAddText(cep, info)

		fmt.Println(data)

	}

}
