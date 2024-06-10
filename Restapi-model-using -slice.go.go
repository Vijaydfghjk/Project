package main

import "fmt"

type employee struct {
	id int

	Name string

	Depat string
}

type company struct {
	temp []employee
}

func myrepo() *company {

	return &company{

		temp: make([]employee, 0),
	}
}

func (a *company) insert(val employee) {

	a.temp = append(a.temp, val)
}

func (a *company) getting() []employee {

	return a.temp
}

func (a *company) change(idd int, newname string) {

	for i, emp := range a.temp {

		if emp.id == idd {

			a.temp[i].Name = newname
			break
		}
	}
}

func (a *company) delete(idd int) {

	for i, emp := range a.temp {

		if emp.id == idd {
			fmt.Println("Value of i", i)
			a.temp = append(a.temp[:i], a.temp[i+1:]...)
			break
		}
	}

}

func (a *company) myfilter(idd int) (employee, bool) {

	for i, emp := range a.temp {

		if emp.id == idd {

			return a.temp[i], true
		}

	}
	return employee{}, false
}

func main() {

	my1 := employee{id: 11, Name: "Mohan", Depat: "BBA"}
	my2 := employee{id: 12, Name: "Mono", Depat: "BcA"}
	my3 := employee{id: 14, Name: "jj", Depat: "MCA"}

	vj := myrepo()

	vj.insert(my1)
	vj.insert(my2)
	vj.insert(my3)

	val := vj.getting()

	fmt.Println("over all :", val)

	vj.change(12, "Raj")

	val2 := vj.getting()

	fmt.Println("Update :", val2)

	vj.delete(14)
	fmt.Println("After deleting")
	val3 := vj.getting()

	fmt.Println("delete:", val3)

	take, found := vj.myfilter(11)

	if found {

		fmt.Println("Found Employee:", take)
	} else {

		fmt.Println("Employee with ID 20 not found")

	}

}
