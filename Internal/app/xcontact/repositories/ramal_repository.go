package repositories

import (
	"ProjetoGustavo/Internal/app/xcontact/model"
	"ProjetoGustavo/Internal/app/xcontact/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var HttpClient = &http.Client{}

var AdicionarRamal = func(ramal model.Ramal) error {
	body, _ := json.Marshal(ramal)
	request, _ := http.NewRequest("POST", "https://ws.wonit.net.br:8004/api/v2/ramal", bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+util.GetToken())
	request.Header.Set("Content-Type", "application/json")

	response, err := HttpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return fmt.Errorf("Erro API: %s", string(bodyBytes))
	}
	return nil
}

var ListarRamais = func() ([]model.Ramal, error) {

	request, erro := http.NewRequest("GET", "https://ws.wonit.net.br:8004/api/v2/ramais", nil)
	if erro != nil {
		return nil, erro
	}

	request.Header.Set("Authorization", "Bearer "+util.GetToken())
	request.Header.Set("Content-Type", "application/json")

	response, erro := HttpClient.Do(request)
	if erro != nil {
		return nil, erro
	}
	defer response.Body.Close()

	bodyBytes, erro := io.ReadAll(response.Body)
	if erro != nil {
		return nil, erro
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("Erro API Xcontact: %s", string(bodyBytes))
	}

	var ramais []model.Ramal
	erro = json.Unmarshal(bodyBytes, &ramais)
	if erro != nil {
		return nil, erro
	}

	return ramais, nil
}

var BuscarRamalPorId = func(id int) (*model.Ramal, error) {

	var ramal model.Ramal

	request, erro := http.NewRequest("GET", fmt.Sprintf("https://ws.wonit.net.br:8004/api/v2/ramal/%d", id), nil)

	if erro != nil {
		return nil, erro
	}

	request.Header.Set("Authorization", "Bearer "+util.GetToken())
	request.Header.Set("Content-Type", "application/json")

	response, erro := HttpClient.Do(request)

	if erro != nil {
		return nil, erro
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("erro na API Xcontact")
	} else if response.StatusCode == 404 {
		return nil, fmt.Errorf("Ramal não encontrado")
	}

	if erro = json.NewDecoder(response.Body).Decode(&ramal); erro != nil {
		return nil, erro
	}

	return &ramal, nil
}

var AtualizarRamal = func(id int, ramal model.Ramal) error {

	body, erro := json.Marshal(ramal)
	if erro != nil {
		return fmt.Errorf("Erro ao converter para JSON: %s", erro.Error())
	}

	request, erro := http.NewRequest("PUT", fmt.Sprintf("https://ws.wonit.net.br:8004/api/v2/ramal/%d", id), bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+util.GetToken())
	request.Header.Set("Content-Type", "application/json")

	response, erro := HttpClient.Do(request)
	if erro != nil {
		return erro
	}
	defer response.Body.Close()

	bodyBytes, erro := io.ReadAll(response.Body)
	if erro != nil {
		return erro
	}

	if response.StatusCode >= 400 {
		return fmt.Errorf("Error API Xcontact: %s", string(bodyBytes))
	}

	return nil
}

var DeletarRamal = func(id int) error {

	request, erro := http.NewRequest("DELETE", fmt.Sprintf("https://ws.wonit.net.br:8004/api/v2/ramal/%d", id), nil)

	if erro != nil {
		return erro
	}

	request.Header.Set("Authorization", "Bearer "+util.GetToken())
	request.Header.Set("Content-Type", "application/json")

	//Manda a requisição pro servidor HTTP usando o Client e armazena a resposta
	//na variavel response
	response, erro := HttpClient.Do(request)

	//caso não houver um response ou algum erro de servidor, retorna o erro:
	if erro != nil {
		return erro
	}

	//O Corpo do JSON é fechado ao final da função para evitar vazamentos de memoria e conexões HTTP
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return fmt.Errorf("erro na API Xcontact")
	} else if response.StatusCode == 404 {
		return fmt.Errorf("Ramal não encontrado")
	}

	return nil
}
