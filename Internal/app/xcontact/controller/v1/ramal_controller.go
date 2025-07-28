package controller

import (
	"ProjetoGustavo/Internal/app/xcontact/dto"
	"ProjetoGustavo/Internal/app/xcontact/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateRamalHandler godoc
// @Summary 			Criar ramal
// @Description 		Criar um ramal
// @Tags 				Ramal
// @Accept       		json
// @Produce      		json
// @Param				request		body 		dto.RamalRequest true "Ramal data"
// @Success      		200     	{object}  	dto.RamalResponse
// @Failure      		400
// @Failure      		404
// @Failure     		500
// @Router       /api/v2/ramal [post]
func CriarRamal(c *gin.Context) {
	var ramalRequest dto.RamalRequest

	if erro := c.ShouldBindJSON(&ramalRequest); erro != nil {
		c.JSON(400, gin.H{"erro: JSON Inválido, detalhe: ": erro.Error()})
		return
	}

	_, erro := service.AdicionarRamal(ramalRequest)

	if erro != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"erro: Falha ao adicionar ramal, detalhe: ": erro.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Ramal criado com sucesso!"})
}

// GetAllRamalHandler godoc
// @Summary      Listar Ramais
// @Description  Listar todos os ramais
// @Tags         Ramal
// @Accept       json
// @Produce      json
// @Success      200 {object} dto.RamalResponse
// @Failure      400
// @Failure      500
// @Router       /api/v2/ramais [get]
func ListarRamais(c *gin.Context) {

	ramais, erro := service.ListarRamais()

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro: Falha ao listar ramais, detalhe: ": erro.Error()})
		return
	}
	c.JSON(http.StatusOK, ramais)

}

// GetByNameRamalHandler godoc
// @Summary      Buscar Ramal
// @Description  Buscar Ramal pelo Nome
// @Tags         Ramal
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do Ramal"
// @Success      200 {object} dto.RamalResponse
// @Failure      400
// @Failure      500
// @Router       /api/v2/ramal/{id} [get]
func BuscarRamalPorId(c *gin.Context) {
	idParametro := c.Param("id")
	id, erro := strconv.Atoi(idParametro)
	fmt.Printf("Token no ")
	if erro != nil {
		fmt.Printf("ID: %s", idParametro)
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}
	ramal, erro := service.BuscarRamalPorId(id)

	if erro != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"erro: Falha ao buscar ramal, detalhe: ": erro.Error()})
		return
	}
	c.JSON(http.StatusOK, ramal)

}

// UpdateRamalHandler godoc
// @Summary      Atualizar Ramal
// @Description  Atualiza os dados de um ramal existente pelo ID
// @Tags         Ramal
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do Ramal"
// @Param        ramal body dto.RamalRequest true "Dados do Ramal para atualização"
// @Success      200 {object} dto.RamalResponse
// @Failure      400
// @Failure      500
// @Router       /api/v2/ramal/{id} [put]
func AtualizarRamal(c *gin.Context) {
	idParametro := c.Param("id")
	id, erro := strconv.Atoi(idParametro)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var ramalRequest dto.RamalRequest

	if erro := c.ShouldBindJSON(&ramalRequest); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	ramalResponse, erro := service.AtualizarRamal(id, ramalRequest)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, ramalResponse)
}

/*
// DeleteRamalHandler godoc
// @Summary      Atualizar Ramal
// @Description  Exclui os dados de um ramal existente pelo ID
// @Tags         Ramal
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do Ramal"
// @Success      200 {object} dto.RamalResponse
// @Failure      400
// @Failure      500
// @Router       /api/v2/ramal/{id} [delete]
func ExcluirRamal(c *gin.Context) {
	idParametro := c.Param("id")

	id, erro := strconv.Atoi(idParametro)

	ramal := service.DeletarRamal(id)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, ramal)
}

*/

// DeleteRamalHandler godoc
// @Summary      Atualizar Ramal
// @Description  Exclui os dados de um ramal existente pelo ID
// @Tags         Ramal
// @Accept       json
// @Produce      json
// @Param        id path int true "ID do Ramal"
// @Success      200 {object} dto.RamalResponse
// @Failure      400
// @Failure      500
// @Router       /api/v2/ramal/{id} [delete]
func ExcluirRamal(c *gin.Context) {
	idParametro := c.Param("id")
	id, err := strconv.Atoi(idParametro)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = service.DeletarRamal(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ramal deletado com sucesso"})
}
