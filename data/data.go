package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetAllEmployees() []Employee {
	var employees []Employee
	db.Find(&employees) // select * from Employee
	return employees
}

func CreateNewEmployee(newEmployee Employee) {
	db.Create(&newEmployee)
}

func UpdateEmployee(employee Employee) {
	db.Save(&employee)
}

func DeleteEmployee(employee *Employee) {
	db.Delete(&employee)
}

func GetEmployee(id int) *Employee {
	var employee Employee
	db.First(&employee, "id = ?", id)
	return &employee
}

var db *gorm.DB

func Init(file, server, database, username, password string, port int) {
	db, _ = gorm.Open(sqlite.Open(file), &gorm.Config{})
	db.AutoMigrate(&Employee{})

	var antal int64
	db.Model(&Employee{}).Count(&antal) // select count(*) from Employee
	if antal == 0 {
		db.Create(&Employee{Age: 50, Namn: "Stefan Holmberg", City: "Teststad"})
		db.Create(&Employee{Age: 14, Namn: "Oliver Holmberg", City: "Teststad"})
		db.Create(&Employee{Age: 20, Namn: "Josefine Holmberg", City: "Uppsala"})
	}
}
