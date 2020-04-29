package main

import "fmt"

type Analyst struct {
	Name  string
	Email string
}

func (a Analyst) Division() string {
	return "dev-team"
}

func (a Analyst) Contact() string {
	return a.Name + " - " + a.Email + " - " + a.Division()
}

type Manager struct {
	Analyst
	MDivision string
}

func (m Manager) Division() string {
	return m.MDivision
}

func main() {

	m := Manager{
		Analyst: Analyst{
			Name:  "Hugo",
			Email: "h@g.com",
		},
		MDivision: "Management",
	}

	fmt.Println(m.Contact())
}
