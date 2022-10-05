package main

import "fmt"

type employee struct {
	name     string
	salary   int
	position string
}

type company struct {
	companyName string
	employees   []employee
}

func main() {
	emp1 := employee{"Amir", 80000, "Full-Stack Developer"}
	emp2 := employee{"Zoha", 40000, "Assistant"}
	emp3 := employee{"Khan", 66000, "Manager"}

	emplys := []employee{emp1, emp2, emp3}

	c1 := company{"ZK", emplys}

	fmt.Printf("Company=  %v  ", c1)
}
