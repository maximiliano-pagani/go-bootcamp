package ex2

import "fmt"

// Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos empleados. Los objetivos son:
// Crear una estructura Person con los campos ID, Name, DateOfBirth.
// Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
// Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
// Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().
// Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	Person
	ID       int
	Position string
}

func (e *Employee) Print() {
	fmt.Println("Employee ID", e.ID)
	fmt.Println(e.Position)
	e.Person.Print()

}

func (p *Person) Print() {
	fmt.Println("Person ID", p.ID)
	fmt.Println(p.Name)
	fmt.Println(p.DateOfBirth)
}

func Ex2() {
	person := Person{
		ID:          535,
		Name:        "Lucas",
		DateOfBirth: "1990-01-01",
	}

	employee := Employee{
		Person:   person,
		ID:       142,
		Position: "Developer",
	}

	employee.Print()
}
