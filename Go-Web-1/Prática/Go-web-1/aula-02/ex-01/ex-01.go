package main

import "github.com/gin-gonic/gin"

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

type Users struct {
	Users []User `json:"users"`
}

func getAll(c *gin.Context) {
	c.File("./users.json")
}

func main() {
	gin.SetMode("debug")
	router := gin.Default()
	router.GET("/", helloHandler)
	router.GET("/users", getAll)
	router.Run()

}
