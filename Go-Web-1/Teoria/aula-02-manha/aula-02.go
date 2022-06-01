package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MENSAGENS HTTP

// REQUEST
// Formado por 5 propriedades: Method, URL, Version, Header, Body

// Method - GET - PUT - POST - DELETE
// URL - equivalente ao path, caminho que queremos acessar.
// Version - Versão do protocolo http que nossa API irá utilizar
// Header - É o cabeçalho e contém os metadados que nossa aplicação pode pedir. ("Infos, cookies, sessões")
// Body - É um bloco de dados arbitrários que vamos enviar em formato JSON

// RESPONSE - Resposta do servidor

// Formado por - Version, Status, Headers e Body

// Respostas HTTP:
// https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Status

// 200 - OK
// 401 - UNAUTHORIZED - EX: Tentou acessar API sem autorização
// 404 - NOT FOUND - EX: Tentou acessar um dado que não existe
// 500 - INTERNAL ERROR - EX: Erros que ainda não tratamos
// 502 - BAD GATEWAY - EX: Não conseguiu completar o caminho

// GIN CONTEXT

// ROUTER
// É um objeto que representa um caminho de URL para um handler.
// no gin é o router := gin.Default()
// router.GET("/", rootHandler)
// router.Run()

// Agrupamento de rotas
// utilizamos o .Group ex:

// group := router.Group("/gophers")
// {
// 	group.GET("/", rootHandler)
// 	group.GET("/hello-world", helloHandler)
// }

func rootHandler(c *gin.Context) {
	//O corpo, cabeçalho e método estão contidos no contexto gin.
	body := c.Request.Body
	header := c.Request.Header
	method := c.Request.Method

	fmt.Println("Eu recebi algo!")
	fmt.Printf("\tMétodo: %s\n", method)
	fmt.Printf("\tConteúdo do cabeçalho:\n")

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}

	fmt.Printf("\tO body é um io.ReadCloser:(%s), e para trabalhar com ele teremos que ler depois\n", body)
	fmt.Println("¡Yay!")
	c.String(200, "Eu recebi!") //Respondemos ao cliente com 200 OK e uma mensagem.
}
func SearchEmployee(c *gin.Context) {
	employee, ok := employee[c.Param("id")]
	if ok {
		c.String(200, "Informação do empregado %s, nome: %s", c.Param("id"), employee)
	} else {
		c.String(404, "Informação do empregado não existente!")
	}
}

// PARAMETRO

// Alguns recursos, como serviços de bando de dados podem ser consultados por um recurso específico
// a partir de uma solicitação. Quais informações são solicitadas são obtidas por meio de parâmetros.

// Primeiro parametro sempre utiliza ? e é algo que queremos concatenar
// Segundo utilizamos um &
// Ex: localhost:8080/users/?name=Vinicius&lastname=Silva

// Definimos um pseudo base de dados onde iremos consultar a informação

// Exemplo de path de params
var employee = map[string]string{
	"644": "Employee A",
	"755": "Employee B",
	"777": "Employee C",
}

func helloHandlers(c *gin.Context) {
	// URL Param ( oq eu vem depois do ":" -> :id)
	// Como só receber numbers no request
	id := c.Param("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "id is not a number",
		})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello: " + id,
	})
}

// Utilizando QueryString

func rootHandlerQuery(c *gin.Context) {
	name := c.Query("name")
	lastName := c.Query("lastName")
	log.Println("name: ", name)
	log.Println("lastName: ", lastName)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello" + name + lastName,
	})
}

// Lendo o contéudo do body

type Employee struct {
	Name     string `json:"name"`
	Password string `json:"name"`
	Id       string `json:"id"`
	Active   string `json:"active"`
}

func employeeHandler(c *gin.Context) {
	var employee Employee
	if err2 := c.ShouldBind(&employee); err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err2.Error()})
		return
	}
	if employee.Name != "user1" || employee.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status: ": "Autorizado"})
}

func main() {
	gin.SetMode("debug")

	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/employees/:id", SearchEmployee)

	// Agrupamento de EndPoints
	group := router.Group("/api/v1")
	{
		group.GET("/hello-world/:id", helloHandlers)
		group.GET("/get")
		group.GET("/info")
	}

	router.Run()
}
