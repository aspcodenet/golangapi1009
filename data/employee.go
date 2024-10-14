package data

// POJO - POCO
// POGO

type Employee struct {
	Id   int
	Age  int
	City string `json:"city"`
	Namn string `json:"name"`
}

// MEDLEMSFUNKTION
// emp.CalculateSalary()
func (emp Employee) CalculateSalary() int {
	if emp.Namn == "Stefan" {
		return 1000
	}
	return 10
}

func CalculateSalary(emp Employee) int {
	if emp.Namn == "Stefan" {
		return 1000
	}
	return 10
}
