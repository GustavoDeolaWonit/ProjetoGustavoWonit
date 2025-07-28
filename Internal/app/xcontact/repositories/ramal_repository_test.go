package repositories_test

import (
	"ProjetoGustavo/Internal/app/xcontact/model"
	"ProjetoGustavo/Internal/app/xcontact/repositories"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	func TestAdicionarRamal(t *testing.T){
			ramal := model.Ramal {
				Nome: "GustavoCabralDeola",
				Numero: "5",

			}

			var b bytes.Buffer
			erro := json.NewDecoder(&b).Encode(ramal)
			if erro != nil {
				t.Error(erro)
			}


			response := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(response)

			body := `{"numero": "123", "nome": "TesteGuzz"}`
			c.Request = httptest.NewRequest(http.MethodPost, "/api/v2/ramal", bytes.NewBufferString(body)))
			c.Request.Header.Set("Content-Type", "application/json")

			 resp, erro := service.AdicionarRamal()

			if response.Code != 200 || response.Code != 201 {
				t.Error("Status : ", response.Code, "o esperado é 201")
			}

			response, erro := io.ReadAll(response.Body)

			if erro != nil {
				t.Error(erro)
			}
	}
*/

func TestAdicionarRamal(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Esperado POST, veio %s", r.Method)
		}

		var ramal model.Ramal
		erro := json.NewDecoder(r.Body).Decode(&ramal)
		if erro != nil {
			t.Errorf("Erro ao decodificar JSON: %v", erro)
		}

		if ramal.Nome == "" {
			t.Errorf("Nome informado não pode estar vazio")
		} else if ramal.Numero == "" {
			t.Errorf("Número do Ramal não pode estar vazio")
		}

		w.WriteHeader(http.StatusCreated)
	}))
	defer mockServer.Close()

	// Sobrescrever client
	repositories.HttpClient = mockServer.Client()

	// Guardar função original
	original := repositories.AdicionarRamal

	// Mock da função
	repositories.AdicionarRamal = func(ramal model.Ramal) error {
		body, _ := json.Marshal(ramal)
		req, _ := http.NewRequest("POST", mockServer.URL+"/api/v2/ramal", bytes.NewBuffer(body))
		req.Header.Set("Authorization", "Bearer fake-token")
		req.Header.Set("Content-Type", "application/json")

		auth := req.Header.Get("Authorization")
		if auth != "Bearer fake-token" {
			t.Errorf("Token errado: %s", auth)
		}

		resp, err := repositories.HttpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			bodyBytes, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("Erro: %s", string(bodyBytes))
		}
		return nil
	}

	// Chamar função
	err := repositories.AdicionarRamal(model.Ramal{
		Nome:   "Teste",
		Numero: "123",
	})
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	// Restaurar
	repositories.AdicionarRamal = original
}

func TestListarRamais(t *testing.T) {
	//Cria um JSON falso como resposta
	respostaFake := `[{"nome":"Ramal1","numero":"123"},{"nome":"Ramal2","numero":"456"}],{"nome":"Ramal3","numero":"789"}]`

	//Cria um servidor HTTP falso pra simular a API e depois valida se a rota que está sendo passada é um GET se não for, retorna um Erro.
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Esperado GET, veio %s", r.Method)
		}
		//Valida se o header da requisição possui token, mesma coisa la do util.getToken()
		auth := r.Header.Get("Authorization")
		if auth != "Bearer fake-token" {
			t.Errorf("Token errado: %s", auth)
		}

		//Retorna o status 200 da API, envia o JSON fake como resposta e fecha o servidor mockado quando o teste terminar.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respostaFake))

	}))
	defer mockServer.Close()

	//Aponta o HTTPClient pro servidor do Mock
	repositories.HttpClient = mockServer.Client()

	//Guarda o ponteiro da função original para restaurar dps para não afetar os outros testes.
	//(Sobrescreve a função para enviar dados ao mockserver)
	originalListarRamais := repositories.ListarRamais

	//Dps ele chama a função de ListarRamais() e faz uma requisição ao mockServer
	repositories.ListarRamais = func() ([]model.Ramal, error) {
		req, _ := http.NewRequest("GET", mockServer.URL+"/api/v2/ramais", nil)
		req.Header.Set("Authorization", "Bearer fake-token")
		req.Header.Set("Content-Type", "application/json")
		//Cria a requisição teste montando os headers e dps envia o response o mock do HttpClient
		response, erro := repositories.HttpClient.Do(req)
		if erro != nil {
			return nil, erro
		}
		defer response.Body.Close()
		// Le o corpo do response e faz um "JSON.Parse()" pro slice de model.Ramal
		bodyBytes, erro := io.ReadAll(response.Body)

		if erro != nil {
			return nil, erro
		}

		var ramais []model.Ramal
		erro = json.Unmarshal(bodyBytes, &ramais)

		if erro != nil {
			return nil, erro
		}

		return ramais, nil

	}
	// Restaura a função original e executa o mock e caso der alguma falha ele cai em algum desses ifs
	defer func() { repositories.ListarRamais = originalListarRamais }()

	ramais, erro := repositories.ListarRamais()
	if erro != nil {
		t.Errorf("Erro inesperado: %v", erro)
	}
	if erro != nil {
		t.Fatalf("Erro inesperado: %v", erro)
	}

	if len(ramais) != 2 {
		t.Errorf("Esperado 2 ramais, veio %d", len(ramais))
	}
	if ramais[0].Nome != "Ramal1" {
		t.Errorf("Esperado Ramal1, veio %s", ramais[0].Nome)
	}
}

func TestBuscarRamalId(t *testing.T) {
	respostaFake := `{"id":10,"nome":"Ramal1","numero":"123"}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Esperado GET, veio %s", r.Method)
		}

		auth := r.Header.Get("Authorization")
		if auth != "Bearer fake-token" {
			t.Errorf("Token errado: %s", auth)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respostaFake))

	}))
	defer mockServer.Close()

	repositories.HttpClient = mockServer.Client()

	originalBuscarRamalId := repositories.BuscarRamalPorId

	repositories.BuscarRamalPorId = func(id int) (*model.Ramal, error) {
		req, _ := http.NewRequest("GET", mockServer.URL+"/api/v2/ramal/"+fmt.Sprintf("%d", id), nil)
		req.Header.Set("Authorization", "Bearer fake-token")
		req.Header.Set("Content-Type", "application/json")

		response, erro := repositories.HttpClient.Do(req)
		if erro != nil {
			return nil, erro
		}
		defer response.Body.Close()

		bodyBytes, erro := io.ReadAll(response.Body)
		if erro != nil {
			return nil, erro
		}
		var ramal model.Ramal

		erro = json.Unmarshal(bodyBytes, &ramal)
		if erro != nil {
			return nil, erro
		}

		return &ramal, nil
	}

	defer func() { repositories.BuscarRamalPorId = originalBuscarRamalId }()

	ramal, erro := repositories.BuscarRamalPorId(10)
	if ramal.Id != 10 {
		t.Errorf("Esperado 10, veio %d", ramal.Id)
	} else if ramal.Id <= 0 {
		t.Errorf("Esperado id maior que 0, veio %d", ramal.Id)
	}

	if erro != nil {
		t.Errorf("Erro inesperado: %v", erro)
	}

}

func TestAtualizarRamal(t *testing.T) {
	respostaFake := `{"id":10,"nome":"Ramal1","numero":"123"}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Esperado PUT, veio %s", r.Method)
		}

		auth := r.Header.Get("Authorization")
		if auth != "Bearer fake-token" {
			t.Errorf("Token errado: %s", auth)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(respostaFake))
	}))
	defer mockServer.Close()

	repositories.HttpClient = mockServer.Client()

	originalAtualizarRamal := repositories.AtualizarRamal

	repositories.AtualizarRamal = func(id int, ramal model.Ramal) error {
		req, _ := http.NewRequest("PUT", mockServer.URL+"/api/v2/ramal/"+fmt.Sprintf("%d", id), nil)
		req.Header.Set("Authorization", "Bearer fake-token")
		req.Header.Set("Content-Type", "application/json")

		response, erro := repositories.HttpClient.Do(req)

		bodyBytes, erro := io.ReadAll(response.Body)
		if erro != nil {
			return erro
		}

		var ramalAtualizada model.Ramal
		erro = json.Unmarshal(bodyBytes, &ramalAtualizada)
		if erro != nil {
			return erro
		}

		if ramalAtualizada.Id != ramal.Id {
			t.Errorf("Esperado id %d, veio %d", ramal.Id, ramalAtualizada.Id)
		}

		if ramalAtualizada.Nome != ramal.Nome {
			t.Errorf("Esperado nome %s, veio %s", ramal.Nome, ramalAtualizada.Nome)
		}

		if ramalAtualizada.Numero != ramal.Numero {
			t.Errorf("Esperado numero %s, veio %s", ramal.Numero, ramalAtualizada.Numero)
		}
		return nil
	}

	defer func() { repositories.AtualizarRamal = originalAtualizarRamal }()

	ramal := model.Ramal{
		Id:     10,
		Nome:   "Ramal1",
		Numero: "123",
	}
	erro := repositories.AtualizarRamal(ramal.Id, ramal)
	if erro != nil {
		t.Errorf("Erro inesperado: %v", erro)
	}
}

func TestDeletarRamal(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Esperado DELETE, veio %s", r.Method)
		}

		auth := r.Header.Get("Authorization")
		if auth != "Bearer fake-token" {
			t.Errorf("Token errado: %s", auth)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	defer mockServer.Close()

	repositories.HttpClient = mockServer.Client()

	originalDeletarRamal := repositories.DeletarRamal

	repositories.DeletarRamal = func(id int) error {

		req, erro := http.NewRequest("DELETE", mockServer.URL+"/api/v2/ramal/"+fmt.Sprintf("%d", id), nil)
		req.Header.Set("Authorization", "Bearer fake-token")
		req.Header.Set("Content-Type", "application/json")
		if erro != nil {
			return erro
		}

		response, erro := repositories.HttpClient.Do(req)
		if erro != nil {
			return erro
		}

		defer response.Body.Close()
		bodyBytes, erro := io.ReadAll(response.Body)
		if erro != nil {
			return erro
		}

		if string(bodyBytes) != "ok" {
			t.Errorf("Esperado ok, veio %s", string(bodyBytes))
		}
		if response.StatusCode != 200 {
			t.Errorf("Esperado 200, veio %d", response.StatusCode)
		}

		return nil

	}

	defer func() { repositories.DeletarRamal = originalDeletarRamal }()
	erro := repositories.DeletarRamal(9)
	if erro != nil {
		t.Errorf("Erro inesperado: %v", erro)
	}

}
