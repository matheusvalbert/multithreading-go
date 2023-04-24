package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ApiCep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
	Message    string `json:"message"`
}

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro"`
}

func main() {
	if len(os.Args) != 2 {
		panic("Wrong number of arguments provided")
	}

	chApiCep := make(chan ApiCep)
	chViaCep := make(chan ViaCep)

	cep := os.Args[1:]

	go func() {
		req, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep[0] + ".json")
		if err != nil {
			fmt.Println("Error to receive data in ApiCep: ", err)
			return
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error to receive data in cdn ApiCep: ", err)
			return
		}
		var data ApiCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Println("Error to receive data in cdn ApiCep: ", err)
			return
		}
		if data.Ok == false {
			fmt.Println("Error to receive data in cdn ApiCep: ", data)
			return
		}
		chApiCep <- data
	}()

	go func() {
		req, err := http.Get("https://viacep.com.br/ws/" + cep[0] + "/json/")
		if err != nil {
			fmt.Println("Error to receive data in ViaCep: ", err)
			return
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("Error to receive data in ViaCep: ", err)
			return
		}
		var data ViaCep
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Println("Error to receive data in ViaCep: ", err)
			return
		}
		if data.Erro == true {
			fmt.Println("Error to receive data in ViaCep: ", data)
			return
		}
		chViaCep <- data
	}()

	select {
	case apiCep := <-chApiCep:
		fmt.Println("Response received from cdn ApiCep:", apiCep)
	case viaCep := <-chViaCep:
		fmt.Println("Response received from ViaCep:", viaCep)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
