package main

import "fmt"

type student struct {
	name string
	roll int
	num  int
}

func main() {
	var s1 student
	s1.name = "Zoha"
	s1.roll = 1234
	print(s1)

}

func print(s1 student) {

	fmt.Println("Name: is ", s1.name)
	fmt.Println("roll: ", s1.roll)
}
