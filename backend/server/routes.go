package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/database"
)

func Run() {
	// Criação do router
	router := gin.Default();

	// Colocar o db aqui!
	db, err := database.NewConnection(); if err != nil {
		fmt.Println(err);
		return;
	} else {
		fmt.Println("Conectado.");
	}

	testApi(router);
	getAllData(router, db);

	router.Run(":8080");
}


// Funções de Rota
func testApi(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Api funcionando",
		})
	})
}

func getAllData(router *gin.Engine, db *gorm.DB) {
	router.GET("/products", func(c *gin.Context) {

		alldata, err := database.GetAllData(db); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
			"data": alldata,
		})
	})
}