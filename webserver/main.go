package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}

func main() {
	// Criando router com Gin
	router := gin.Default()

	// gin.SetMode("release") // Opção para o Gin em produção

	// Criando arquivo de log do servidor
	f, err := os.OpenFile("gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// Capturar o request GET "/hello-world"
	router.GET("/hello-world", helloHandler)
	router.Run()

}
