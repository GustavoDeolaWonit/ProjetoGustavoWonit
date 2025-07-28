package service_test

import (
	"ProjetoGustavo/Internal/app/xcontact/dto"
	"ProjetoGustavo/Internal/app/xcontact/model"
	"ProjetoGustavo/Internal/app/xcontact/repositories"
	"ProjetoGustavo/Internal/app/xcontact/service"
	"errors"
	"testing"
)

func TestAdicionarRamal(t *testing.T) {
	// Salva a função original do repositório (se não fizer isso ele vai continuar mockado em outros testes)
	original := repositories.AdicionarRamal
	// O mock é removido ao final do teste por isso o defer e a função é restaurada no final do teste
	// para não afetar outros testes que rodam dps
	defer func() { repositories.AdicionarRamal = original }()

	t.Run("Sucesso ao adicionar ramal", func(t *testing.T) {
		repositories.AdicionarRamal = func(ramal model.Ramal) error {
			// Simula sucesso, mas não altera o ramal (não é ponteiro), a função original é substituida por um mock
			return nil
		}
		//Monta o request do mock
		request := dto.RamalRequest{
			Numero: "101",
			Nome:   "Ramal Teste",
			Senha:  "1234",
			Grupo:  "Grupo A",
			Allow:  "all",
		}

		// Monta o valor de response que estou esperando
		// Como não é atribuido o ramal.ID na struct "ramal", emtão simulo que o banco atribuiria ID 1
		esperado := dto.RamalResponse{
			Id:     0, // <- O ID será 0 porque o service não atribui o valor
			Numero: "101",
			Nome:   "Ramal Teste",
			Senha:  "1234",
			Grupo:  "Grupo A",
			Allow:  "all",
		}
		//Aqui ele monta a requisição para fazer o POST na API
		resultado, erro := service.AdicionarRamal(request)
		if erro != nil {
			t.Errorf("Não era esperado erro, mas veio: %v", erro)
		}
		//Verifica se o resultado é o mesmo do esperado
		if resultado.Numero != esperado.Numero ||
			resultado.Nome != esperado.Nome ||
			resultado.Senha != esperado.Senha ||
			resultado.Grupo != esperado.Grupo ||
			resultado.Allow != esperado.Allow {
			t.Errorf("Resultado inesperado. Esperado %+v, veio %+v", esperado, resultado)
		}
	})

	t.Run("Erro ao adicionar ramal", func(t *testing.T) {
		repositories.AdicionarRamal = func(ramal model.Ramal) error {
			return errors.New("erro ao salvar no banco")
		}
		//Subteste para simular o comportamento de falha do banco
		request := dto.RamalRequest{
			Numero: "102",
			Nome:   "Erro Ramal",
			Senha:  "1234",
		}

		//aqui é enviado uma requisição invalida
		_, erro := service.AdicionarRamal(request)
		if erro == nil {
			t.Errorf("Esperado erro, mas não ocorreu")
		}
		//dps a mensagem esperada do erro do banco
		mensagemEsperada := "erro ao salvar no banco"
		if erro.Error() != mensagemEsperada {
			t.Errorf("Mensagem de erro inesperada. Esperado: %s, veio: %s", mensagemEsperada, erro.Error())
		}
	})
}

func TestListarRamais(t *testing.T) {
	original := repositories.ListarRamais

	defer func() { repositories.ListarRamais = original }()

	t.Run("Sucesso ao listar ramal", func(t *testing.T) {
		repositories.ListarRamais = func() ([]model.Ramal, error) {
			return []model.Ramal{
				{
					Id:     1,
					Numero: "101",
					Nome:   "Ramal Teste",
					Senha:  "1234",
					Grupo:  "Grupo A",
					Allow:  "all",
				},
				{
					Id:     2,
					Numero: "102",
					Nome:   "Ramal Teste 2",
					Senha:  "12346",
					Grupo:  "Grupo A",
					Allow:  "all",
				},
				{
					Id:     3,
					Numero: "103",
					Nome:   "Ramal Teste 3",
					Senha:  "12345",
					Grupo:  "Grupo B",
					Allow:  "all",
				},
			}, nil
		}

		resultado, erro := repositories.ListarRamais()

		if erro != nil {
			t.Errorf("Erro ao listar ramal. Erro: %v", erro)
		}

		if len(resultado) != 2 {
			t.Errorf("Esperado %d resultado", len(resultado))
		}

		if resultado[0].Numero != "101" || resultado[1].Nome != "Ramal 2" {
			t.Errorf("Dados dos ramais incorretos: %v", resultado)
		}

	})

}

func TestBuscarRamalPorId(t *testing.T) {
	original := repositories.BuscarRamalPorId
	defer func() { repositories.BuscarRamalPorId = original }()

	t.Run("Sucesso ao buscar ramal", func(t *testing.T) {
		repositories.BuscarRamalPorId = func(id int) (*model.Ramal, error) {
			return &model.Ramal{
				Id:     id,
				Numero: "101",
				Nome:   "João",
				Senha:  "123",
				Grupo:  "GRUPO A",
				Allow:  "all",
			}, nil

		}
		resultado, erro := service.BuscarRamalPorId(1)
		if erro != nil {
			t.Errorf("Não era esperado erro, mas veio: %v", erro)
		}
		if resultado.Numero != "101" || resultado.Nome != "João" {
			t.Errorf("Esperado 101, veio %s", resultado.Numero)
		}

	})

	t.Run("Erro ao buscar ramal por ID", func(t *testing.T) {
		repositories.BuscarRamalPorId = func(id int) (*model.Ramal, error) {
			return nil, errors.New("ramal não encontrado")
		}

		_, erro := service.BuscarRamalPorId(999)

		if erro == nil {
			t.Errorf("Era esperado erro, mas veio nil")
		}

		mensagemEsperada := "ramal não encontrado"
		if erro.Error() != mensagemEsperada {
			t.Errorf("Mensagem de erro inesperada. Esperado: %s, veio: %s", mensagemEsperada, erro.Error())
		}
	})
}

func TestAtualizarRamal(t *testing.T) {
	original := repositories.AtualizarRamal
	defer func() { repositories.AtualizarRamal = original }()

	t.Run("Sucesso ao atualizar ramal", func(t *testing.T) {
		repositories.AtualizarRamal = func(id int, ramal model.Ramal) error {
			return nil
		}

		request := dto.RamalRequest{
			Numero: "201",
			Nome:   "Ramal Atualizado",
			Senha:  "senha123",
			Grupo:  "Grupo X",
			Allow:  "all",
		}

		id := 10

		resultado, erro := service.AtualizarRamal(id, request)
		if erro != nil {
			t.Errorf("Não era esperado erro, mas veio: %v", erro)
		}

		if resultado.Numero != request.Numero || resultado.Nome != request.Nome {
			t.Errorf("Atualização falhou. Esperado nome: %s, veio: %s", request.Nome, resultado.Nome)
		}
	})

	t.Run("Erro ao atualizar ramal", func(t *testing.T) {
		repositories.AtualizarRamal = func(id int, ramal model.Ramal) error {
			return errors.New("falha ao atualizar no banco")
		}

		request := dto.RamalRequest{
			Numero: "202",
			Nome:   "Ramal Erro",
			Senha:  "senha456",
		}

		id := 20

		_, erro := service.AtualizarRamal(id, request)

		if erro == nil {
			t.Errorf("Era esperado um erro, mas veio nil")
		}

		mensagemEsperada := "falha ao atualizar no banco"
		if erro.Error() != mensagemEsperada {
			t.Errorf("Mensagem de erro inesperada. Esperado: %s, veio: %s", mensagemEsperada, erro.Error())
		}
	})
}
