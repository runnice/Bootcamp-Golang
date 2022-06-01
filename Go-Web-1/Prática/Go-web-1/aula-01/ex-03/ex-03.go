package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Lastname  string  `json:"lastname	"`
	Email     string  `json:"email"`
	Age       uint    `json:"age"`
	Height    float64 `json:"height"`
	IsActive  bool    `json:"isActive"`
	BirthDate string  `json:"birthDate"`
}

type Users struct {
	Users []User `json:"users"`
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Vinicius!",
	})
}
func getAllWithJson(c *gin.Context) {
	c.File("./users.json")
}

func getAll(c *gin.Context) {

	jsonDataRaw := `[
		{
		"id": 1,
		"name": "Joao",
		"lastmame": "Silva",
		"email": "joaosilva@gemail.com",
		"age": 20,
		"height": 1.80,
		"isActive": true,
		"birthDate": "01-01-2020"},
	{
		"id":2,
		"name":"Joana",
		"lastname": "Guedes",
		"email": "joana@gemail.com",
		"age": 35,
		"height": 1.85,
		"isActive": true,
		"birthDate": "01/01/1990"
	}
	]`

	var u []User

	if err := json.Unmarshal([]byte(jsonDataRaw), &u); err != nil {
		fmt.Println(err)
	}
	c.JSON(200, u)

}

// Exercício 02 - GoWeb-1 Aula 2
func filterJSON(fileName string) Users {
	arquivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer arquivo.Close()

	byteValue, _ := ioutil.ReadAll(arquivo)

	var u Users
	var u2 Users

	if err := json.Unmarshal(byteValue, &u); err != nil {
		fmt.Println("Error: ", err)
	}

	// m := make(map[string]interface{})

	for _, v := range u.Users {
		if v.Age == 30 {
			u2.Users = append(u2.Users, v)
		}

	}
	return u2

}

func readJson(fileName string) Users {
	arquivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer arquivo.Close()

	byteValue, _ := ioutil.ReadAll(arquivo)

	var u Users

	if err := json.Unmarshal(byteValue, &u); err != nil {
		fmt.Println("Error: ", err)
	}

	// m := make(map[string]interface{})

	return u

}

// Exercício 2

func GetId(c *gin.Context) {
	id := c.Param("id")
	convertId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "id must be a number"})
		log.Println(err)

	}
	users := readJson("./users.json")
	for _, users := range users.Users {
		if users.Id == convertId {
			c.JSON(http.StatusOK, gin.H{
				"user": users,
			})

		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user not found",
	})

}

func filter(c *gin.Context) {

	c.JSON(200, filterJSON("./users.json"))

}
func main() {

	gin.SetMode("debug")
	router := gin.Default()
	router.GET("/", helloHandler)
	router.GET("/users", getAllWithJson)
	router.GET("/user", getAll)
	router.GET("/filter", filter)
	router.GET("/:id", GetId)

	group := router.Group("/usuarios")
	{
		group.GET("/", getAllWithJson)
		group.GET("/:id", GetId)
	}

	router.Run()
}
