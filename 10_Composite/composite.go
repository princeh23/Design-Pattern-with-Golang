package main

import "fmt"

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (Employee) Count() int {
	return 1
}

type Department struct {
	Name string

	SubOrganizations []IOrganization
}

func (d Department) Count() int {
	c := 0
	for _, org := range d.SubOrganizations {
		c += org.Count()
	}
	return c
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func main() {
	root := Department{Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(Employee{Name: "sub" + string('0'+i)})
	}
	fmt.Printf("%+v", root.SubOrganizations)

}
