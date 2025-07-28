package auth

import (
	"ProjetoGustavo/Internal/app/xcontact/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func efetuaLogin() (string, error) {
	body, _ := json.Marshal(model.LoginRequestAPI{
		Nome:  "gustavo",
		Senha: "gustavo",
	})

	request, _ := http.NewRequest("POST", "https://ws.wonit.net.br:8004/api/v4/login/supervisor", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	if response.StatusCode >= 400 {
		return "", fmt.Errorf("Error API Xcontact: %s", string(bodyBytes))
	}

	var result struct {
		Token string `json:"token"`
	}

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", err
	}

	return result.Token, nil
}

func GetLogin() string {
	token, _ := efetuaLogin()
	return token
}
