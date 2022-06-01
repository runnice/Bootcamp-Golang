package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Lastname  string  `json:"lastname	"`
	Email     string  `json:"email"`
	Age       uint    `json:"age"`
	Height    float64 `json:"height"`
	IsActive  bool    `json:"isActive"`
	BirthDate string  `json:"birthDate"`
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Vinicius!",
	})
}
func getAllWithJson(c *gin.Context) {
	c.JSON(200, readFile("users.json"))

}

func readFile(fileName string) User {
	arquivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer arquivo.Close()

	byteValue, _ := ioutil.ReadAll(arquivo)

	var u User

	if err := json.Unmarshal(byteValue, &u); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(u)
	return u

}

func getAll(c *gin.Context) {

	jsonDataRaw := `[
		{
		"Id": 1,
		"Name": "Joao",
		"Lastmame": "Silva",
		"Email": "joaosilva@gemail.com",
		"Age": 20,
		"Height": 1.80,
		"IsActive": true,
		"BirthDate": "01-01-2020"},
	{
		"Id":2,
		"Name":"Joana",
		"Lastname": "Guedes",
		"Email": "joana@gemail.com",
		"Age": 35,
		"Height": 1.85,
		"IsActive": true,
		"BirthDate": "01/01/1990"
	}
	]`

	var u []User

	if err := json.Unmarshal([]byte(jsonDataRaw), &u); err != nil {
		fmt.Println(err)
	}
	c.JSON(200, u)
	// for _, i := range u {
	// 	c.JSON(200, i)

	// }

}

func main() {

	gin.SetMode("debug")
	router := gin.Default()
	router.GET("/", helloHandler)
	router.GET("/users", getAllWithJson)
	router.Run()

	// if err := json.Unmarshal([]byte(jsonDataRaw), &u); err != nil{
	// 	fmt.Println(err)
	// }
	// for _, i := range u{
	// 	c.JSON(200, i)
	// }

	// jsonData, err := os.OpenFile("users.json", os.O_RDONLY, 0644)
	// if err2 := json.Unmarshal([]byte(jsonData), &users); err2 != nil {
	// 	log.Fatal(err2)
	// }

}
