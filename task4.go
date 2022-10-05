package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
	course     []string
}

func NewStudent(rollno int, name string, address string, cour []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.course = cour
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, course []string) *Student {
	st := NewStudent(rollno, name, address, course)
	ls.list = append(ls.list, st)
	return st
}

func PrintStudentList(ls *StudentList) {

	for i := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))
		fmt.Printf("student rollno              %d\n", ls.list[i].rollnumber)
		fmt.Printf("student name                %s\n", ls.list[i].name)
		fmt.Printf("student address             %s\n", ls.list[i].address)
		fmt.Printf("student course             %s\n", ls.list[i].course)

		result := strconv.Itoa(ls.list[i].rollnumber) + ls.list[i].name + ls.list[i].address
		sum := sha256.Sum256([]byte(result))
		fmt.Printf("%x\n", sum) //hexadecimal
	}

}

func main() {
	student := new(StudentList)

	student.CreateStudent(24, "Asim", "AAAAAAA", []string{"English", "Urdu"})
	student.CreateStudent(25, "Naveed", "BBBBBB", []string{"Blockchain", "Math", "Urdu"})
	PrintStudentList(student)
}
