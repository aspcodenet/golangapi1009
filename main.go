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

func handleUpdateEmployee(c *gin.Context) { // PUT Bodyn JSON {"Age":16, "City":"Nacka"}"
	// /api/employee/21
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employeeFromDB := data.GetEmployee(numId) // DEN SOM REDAN FINNS I DATABASEN !!!
	if employeeFromDB == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		var employeeJson data.Employee

		if err := c.BindJSON(&employeeJson); err != nil {
			return
		}
		employeeJson.Id = numId
		data.UpdateEmployee(employeeJson)
		//data. Save(&employeeJson)
		c.IndentedJSON(http.StatusOK, employeeJson)
	}
}

func handleDeleteEmployee(c *gin.Context) { // PUT Bodyn JSON {"Age":16, "City":"Nacka"}"
	// /api/employee/21
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employeeFromDB := data.GetEmployee(numId) // DEN SOM REDAN FINNS I DATABASEN !!!
	if employeeFromDB == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		data.DeleteEmployee(employeeFromDB)
		//data. Save(&employeeJson)
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Deleted ok"})
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

func handleStart(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body>Hello</body></html>"))
}

type PageView struct {
	Title  string
	Rubrik string
}

func handleAbout(c *gin.Context) {

	c.HTML(http.StatusOK, "about.html", &PageView{Title: "test", Rubrik: data.GetEmployee(1).Namn})
}

var config Config

func main() {
	readConfig(&config)

	data.Init(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	r := gin.Default()
	r.LoadHTMLGlob("templates/**")
	// GET all
	// GET one (by id)
	// POST Create new
	// PUT, DELETE

	r.GET("/", handleStart)
	r.GET("/about", handleAbout)
	r.GET("/api/employee", handleGetAllEmployees)
	r.GET("/api/employee/:id", handleGetOneEmployee)
	r.POST("/api/employee", handleNewEmployees) // SKA JU Employee skickas med som JSON
	// PUT = replace (ALLA properties)
	// PATCH = update a few properties
	r.PUT("/api/employee/:id", handleUpdateEmployee)    // SKA JU Employee skickas med som JSON
	r.DELETE("/api/employee/:id", handleDeleteEmployee) // SKA JU Employee skickas med som JSON

	// r.GET("/omoss", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 		"nu":      "rast",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
