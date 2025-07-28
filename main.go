// @title API Wonit
// @version 2.0
// @Description Api para cadastro de consulta de ramais
package main

import (
	"ProjetoGustavo/Internal/app/xcontact/controller/v1"
	_ "ProjetoGustavo/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/v2/ramal", controller.CriarRamal)
	r.GET("/api/v2/ramais", controller.ListarRamais)
	r.GET("/api/v2/ramal/:id", controller.BuscarRamalPorId)
	r.PUT("/api/v2/ramal/:id", controller.AtualizarRamal)
	r.DELETE("/api/v2/ramal/:id", controller.ExcluirRamal)

	err := r.Run()

	if err != nil {
		log.Fatal(err)
	}
}
