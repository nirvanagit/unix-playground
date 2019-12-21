package main

type Person struct {
	Name string
	Age  int
}

func main() {
	var me Person
	me.Name = "foo"
	me.Age = 0

	var wife Person
	wife.Name = "bar"
	wife.Age = 0

	var mother Person
	mother.Name = "eoo"
	mother.Age = 1

	var father Person
	father.Name = "goo"
	father.Age = 1

	var family []Person
	family = append(family, me, wife, mother, father)

	birthdayToday(&me)
}

func birthdayToday(person *Person) {
	person.Age = person.Age + 1
}
