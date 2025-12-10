package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/database"
)

func Run() {
	// Criação do router
	router := gin.Default();
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Colocar o db aqui!
	db, err := database.NewConnection(); if err != nil {
		fmt.Println(err);
		return;
	} else {
		fmt.Println("Conectado.");
	}

	testApi(router);
	getAllProducts(router, db);
	addProduct(router, db);
	updateProduct(router, db);
	getAllImage(router, db);

	// Files
	setupFileRoutes(router);

	router.Run(":8080");
}


// Funções de Produto
func testApi(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Api funcionando",
		})
	})
}

func getAllProducts(router *gin.Engine, db *gorm.DB) {
	router.GET("/products", func(c *gin.Context) {

		alldata, err := database.GetAllProduct(db); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
		}

		c.JSON(http.StatusOK, gin.H{
			"data": alldata,
		})
	})
}

func addProduct(router *gin.Engine, db *gorm.DB) {
	
	router.POST("/product", func(c *gin.Context) {

		product := database.Product{};

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		database.AddProduct(db, product);

		c.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})
}

func updateProduct(router *gin.Engine, db *gorm.DB) {



	router.PUT("/product/:id", func(c *gin.Context) {

		product := database.Product{};

		str_id := c.Param("id");

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		database.UpdateProduct(db, uuid.MustParse(str_id), product);

		c.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})
}

// Funções de Arquivo
func setupFileRoutes(router *gin.Engine) {
	router.GET("/files/:filename", func (c *gin.Context)  {
		filename := c.Param("filename");

		filePath := filepath.Join("..", "uploads", filename);

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "File not found",
			})
			return;
		}

		c.File(filePath)
	})
}

// Funções de Imagem
func getAllImage(router *gin.Engine, db *gorm.DB) {
	router.GET("/images", func(c *gin.Context) {

		alldata, err := database.GetAllImage(db); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
		}

		c.JSON(http.StatusOK, gin.H{
			"data": alldata,
		})
	})
}