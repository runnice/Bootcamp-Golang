package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ENTENDENDO O POST

// POST - Metódo é usado para criar um novo registro ou seja, completamente novo.
// No post podemos enviar arquivos, como imagens vídeos ou áudios.

// GERAR SOLICITAÇÕES POST DO GO

// ENDENDA  O QUE SÃO HEADERS

//  context.GetHeader("meu_header")

// type Products struct {
// 	Products []Product `json:"products"`
// }

type Product struct {
	Id         int64   `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

type Products []Product

var lastID int64
var products Products

func listProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": &products})
}

func produtos(c *gin.Context) {

	// Setando que o token tem que ser igual a 123456, caso contrário não vai prosseguir.
	token := c.GetHeader("token")

	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token inválido",
		})
		var produto Product

		if err := c.ShouldBindJSON(&produto); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++
		produto.Id = lastID
		products = append(products, produto)

		c.JSON(200, gin.H{
			"data": produto,
		})
	}
}
func main() {
	gin.SetMode("debug")
	router := gin.Default()

	// GRUPO DE POST PRODUTOS
	group := router.Group("/produtos")
	{
		group.GET("/", listProducts)
		group.POST("/", produtos)
	}

	// router.POST("/", produtos)
	router.Run()
}
