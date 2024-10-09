package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"systementor.se/apidemo1009/data"
)

// request, response
func handleGetAllEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func handleGetOneEmployee(c *gin.Context) {
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employee := data.GetEmployee(numId)

	if employee == nil { // INTE HITTAT  /api/employee/
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func handleNewEmployees(c *gin.Context) {
	// TODO Add new
	// försöka få fram den JSON Employee som man skickat in
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	data.CreateNewEmployee(employee)
	c.IndentedJSON(http.StatusCreated, employee)
}

func main() {
	data.Init()

	r := gin.Default()
	// GET all
	// GET one (by id)
	// POST Create new
	// PUT, DELETE

	r.GET("/api/employee", handleGetAllEmployees)
	r.GET("/api/employee/:id", handleGetOneEmployee)
	r.POST("/api/employee", handleNewEmployees) // SKA JU Employee skickas med som JSON

	// r.GET("/omoss", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 		"nu":      "rast",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
