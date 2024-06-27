package main

import (
	"fmt"
	"math/rand"
	"unicode/utf8"
)

var (
	eHRDept     *HRDept
	eTeckDept   *TeckDept
	eFinaceDept *FinaceDept
)

func randomName() string {
	buf := make([]byte, utf8.UTFMax)
	for i := 0; i < 5; i++ {
		buf = append(buf, byte(rand.Intn(24)+int('a')))
	}
	return string(buf)
}

func init() {
	eHRDept = &HRDept{
		Employees: make([]*Employee, 0),
	}
	for i := 0; i < 100; i++ {
		eHRDept.Employees = append(eHRDept.Employees, &Employee{
			No:          10000 + i,
			Name:        randomName(),
			Performance: rand.Intn(100),
			Skill:       rand.Intn(100),
			Salary:      rand.Intn(100),
		})
	}
	eTeckDept = &TeckDept{eHRDept}
	eFinaceDept = &FinaceDept{eHRDept}
}

type Employee struct {
	No          int
	Name        string
	Performance int // 0 - 100
	Skill       int // 0 - 100
	Salary      int // 0 - 100
}

type HRDept struct {
	Employees []*Employee
}

func (hr *HRDept) GoodEmployeesRange(yield func(*Employee) bool) {
	for _, e := range hr.Employees {
		if e.Performance > 80 {
			if !yield(e) {
				return
			}
		}
	}
}

type TeckDept struct {
	*HRDept
}

func (dept *TeckDept) GoodEmployees() {
	for e := range dept.GoodEmployeesRange {
		if e.Skill > 80 {
			fmt.Println("teck good employee:", e.Name)
		}
	}
}

type FinaceDept struct {
	*HRDept
}

func (dept *FinaceDept) GoodEmployees() {
	for e := range dept.GoodEmployeesRange {
		if e.Salary < 20 {
			fmt.Println("finace good employee: ", e.Name)
		}
	}
}
