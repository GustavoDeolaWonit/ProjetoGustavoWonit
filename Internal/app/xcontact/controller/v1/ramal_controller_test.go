package controller

import (
	"ProjetoGustavo/Internal/app/xcontact/dto"
	"ProjetoGustavo/Internal/app/xcontact/service"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCriarRamal(t *testing.T) {
	gin.SetMode(gin.TestMode)
	//Subteste isolado para Sucesso ao criar Ramal
	t.Run("Sucesso ao criar Ramal", func(t *testing.T) {
		//Mock da função AdicionarRamal do service

		service.AdicionarRamal = func(r dto.RamalRequest) (dto.RamalResponse, error) {
			return dto.RamalResponse{}, nil
		}
		//Criando o corpo da requisição para enviar no POST
		requestBody := dto.RamalRequest{
			Numero: "1001",
			Nome:   "Teste",
			Senha:  "1234",
			Grupo:  "Grupo A",
			Allow:  "all",
		}
		//Converte o corpo da requisição (struct) em JSON
		bodyBytes, _ := json.Marshal(requestBody)
		//Por estar usando Gin aqui eu posso fazer um httptest.NewRecorder()
		//Mas faz a mesma coisa que httpClient := mockServer() (Cria um Http fake)
		w := httptest.NewRecorder()
		//Cria o Contexto no Gin com a resposta do httptest. Ele é o contexto que é passado pro
		// handler pra "simular uma requisição real"
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("POST", "/api/v2/ramal", bytes.NewBuffer(bodyBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		//aqui o handler real é executado
		CriarRamal(c)

		if w.Code != http.StatusCreated {
			t.Errorf("Esperado status 201, veio %d", w.Code)
		}

		//Verifica o JSON Body se contém o texto "Ramal Criado com Sucesso"
		assert.Contains(t, w.Body.String(), "Ramal criado com sucesso!")
	})
	//Outro subteste para validar o corpo JSON
	t.Run("Erro ao enviar JSON inválido", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Corpo inválido (faltando aspas no JSON)
		body := []byte(`{"Numero": 1001, "Nome": Teste}`)
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/v2/ramal", bytes.NewBuffer(body))
		c.Request.Header.Set("Content-Type", "application/json")

		CriarRamal(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "JSON Inválido")
	})
	//Subteste para tratar de erro ao adicionar ramal
	t.Run("Erro interno ao adicionar ramal", func(t *testing.T) {
		// Mocka a função AdicionarRamal para retornar erro
		service.AdicionarRamal = func(r dto.RamalRequest) (dto.RamalResponse, error) {
			return dto.RamalResponse{}, errors.New("falha no banco de dados")
		}

		reqBody := dto.RamalRequest{
			Numero: "1002",
			Nome:   "Erro",
			Grupo:  "Grupo B", //Adicionei o campo Grupo para não dar erro no teste, mas se remover esse o teste da fail aqui
			Senha:  "1234",
		}
		bodyBytes, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodPost, "/api/v2/ramal", bytes.NewBuffer(bodyBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		CriarRamal(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Falha ao adicionar ramal")
	})
}

func TestListarRamais(t *testing.T) {
	gin.SetMode(gin.TestMode)

	original := service.ListarRamais
	defer func() { service.ListarRamais = original }()

	t.Run("Sucesso ao listar ramais", func(t *testing.T) {
		service.ListarRamais = func() ([]dto.RamalResponse, error) {
			return []dto.RamalResponse{
				{
					Id:     1,
					Numero: "1001",
					Nome:   "Ramal 1",
				},
				{
					Id:     2,
					Numero: "1002",
					Nome:   "Ramal 2",
				},
			}, nil
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/api/v2/ramais", nil)

		ListarRamais(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var resposta []dto.RamalResponse
		erro := json.Unmarshal(w.Body.Bytes(), &resposta)
		assert.NoError(t, erro)
		assert.Len(t, resposta, 2)
		assert.Equal(t, "1001", resposta[0].Numero)
	})

	t.Run("Erro ao listar ramais", func(t *testing.T) {
		service.ListarRamais = func() ([]dto.RamalResponse, error) {
			return nil, errors.New("erro no banco")
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest(http.MethodGet, "/api/v2/ramais", nil)

		ListarRamais(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "erro no banco")
	})
}

func TestBuscarRamalPorId(t *testing.T) {
	// Mocka a função do service
	original := service.BuscarRamalPorId
	defer func() { service.BuscarRamalPorId = original }()

	router := gin.Default()
	router.GET("/api/v2/ramal/:id", BuscarRamalPorId)

	t.Run("Buscar ramal com ID válido", func(t *testing.T) {
		service.BuscarRamalPorId = func(id int) (dto.RamalResponse, error) {
			return dto.RamalResponse{
				Id:     1,
				Numero: "101",
				Nome:   "João",
			}, nil
		}

		req, _ := http.NewRequest("GET", "/api/v2/ramal/1", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Esperado status 200, veio %d", resp.Code)
		}
	})

	t.Run("Buscar ramal com ID inválido (não é número)", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v2/ramal/abc", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Esperado status 400, veio %d", resp.Code)
		}
	})

	t.Run("Erro ao buscar ramal do service", func(t *testing.T) {
		service.BuscarRamalPorId = func(id int) (dto.RamalResponse, error) {
			return dto.RamalResponse{}, errors.New("erro de banco")
		}

		req, _ := http.NewRequest("GET", "/api/v2/ramal/2", nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Esperado status 500, veio %d", resp.Code)
		}
	})
}

func TestAtualizarRamal(t *testing.T) {
	original := service.AtualizarRamal
	defer func() { service.AtualizarRamal = original }()

	body := `{"numero": "105", "nome": "Atualizado", "senha": "1234", "grupo": "Grupo A", "allow": "all"}`
	//body := `{"numero": "105", "nome": "Atualizado"}` // JSON para forçar a falha do teste

	t.Run("Atualizar ramal com sucesso", func(t *testing.T) {
		service.AtualizarRamal = func(id int, r dto.RamalRequest) (dto.RamalResponse, error) {
			return dto.RamalResponse{Id: id, Numero: r.Numero, Nome: r.Nome}, nil
		}

		router := gin.Default()
		router.PUT("/api/v2/ramal/:id", AtualizarRamal)

		request, _ := http.NewRequest("PUT", "/api/v2/ramal/1", strings.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Esperado 200, veio %d", response.Code)
		}
	})

	t.Run("ID inválido", func(t *testing.T) {
		router := gin.Default()
		router.PUT("/api/v2/ramal/:id", AtualizarRamal)

		req, _ := http.NewRequest("PUT", "/api/v2/ramal/abc", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Esperado 400, veio %d", resp.Code)
		}
	})

	t.Run("JSON inválido", func(t *testing.T) {
		router := gin.Default()
		router.PUT("/api/v2/ramal/:id", AtualizarRamal)

		req, _ := http.NewRequest("PUT", "/api/v2/ramal/1", strings.NewReader("{invalid json}"))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Esperado 400, veio %d", resp.Code)
		}
	})

	t.Run("Erro do service", func(t *testing.T) {
		service.AtualizarRamal = func(id int, r dto.RamalRequest) (dto.RamalResponse, error) {
			return dto.RamalResponse{}, errors.New("erro ao atualizar")
		}

		router := gin.Default()
		router.PUT("/api/v2/ramal/:id", AtualizarRamal)

		request, _ := http.NewRequest("PUT", "/api/v2/ramal/1", strings.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusInternalServerError {
			t.Errorf("Esperado 500, veio %d", response.Code)
		}
	})
}

func TestExcluirRamal(t *testing.T) {
	original := service.DeletarRamal
	defer func() { service.DeletarRamal = original }()

	t.Run("Deletar ramal com sucesso", func(t *testing.T) {
		service.DeletarRamal = func(id int) error {
			return nil
		}

		router := gin.Default()
		router.DELETE("/api/v2/ramal/:id", ExcluirRamal)

		request, _ := http.NewRequest("DELETE", "/api/v2/ramal/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("Esperado 200, veio %d", response.Code)
		}
	})

	t.Run("Erro ao deletar", func(t *testing.T) {
		service.DeletarRamal = func(id int) error {
			return errors.New("Erro ao deletar ramal")
		}

		router := gin.Default()
		router.DELETE("/api/v2/ramal/:id", ExcluirRamal)

		request, _ := http.NewRequest("DELETE", "/api/v2/ramal/1", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusInternalServerError {
			t.Errorf("Esperado 500, veio %d", response.Code)
		}
	})

	t.Run("ID Inválido", func(t *testing.T) {
		router := gin.Default()
		router.DELETE("/api/v2/ramal/:id", ExcluirRamal)

		request, _ := http.NewRequest("DELETE", "/api/v2/ramal/abc", nil)
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("Esperado 400, veio %d", response.Code)
		}
	})
}
