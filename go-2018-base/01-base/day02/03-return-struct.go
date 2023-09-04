package main

import "fmt"

type Person3 struct {
	name     string
	sex      byte
	age      int
	interest []string
}

func initPerson(p *Person3) {
	p.name = "tom"
	p.sex = 'm'
	p.age = 23
	p.interest = append(p.interest, "shopping")
	p.interest = append(p.interest, "watching tv")
}

func getPerson() *Person3 {
	var p *Person3 = new(Person3)
	p.name = "init"
	p.sex = 'f'
	p.age = 18
	return p
}

func main() {
	var p1 Person3
	initPerson(&p1)
	fmt.Println("p1: ", p1)
	var p2 *Person3 = getPerson()
	fmt.Println("p2: ", *p2)

}
