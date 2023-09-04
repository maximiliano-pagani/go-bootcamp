package ex1

import "fmt"

// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

type Student struct {
	Name     string
	Lastname string
	DNI      int
	Date     string
}

func (a *Student) Detail() {
	fmt.Println("Nombre:", a.Name)
	fmt.Println("Apellido:", a.Lastname)
	fmt.Println("DNI:", a.DNI)
	fmt.Println("Fecha:", a.Date)
}

func Ex1() {
	var students []Student

	students = []Student{
		Student{Name: "Mike", Lastname: "Dafoe", DNI: 32494, Date: "1994-05-05"},
		Student{Name: "Tom", Lastname: "Williams", DNI: 37575, Date: "1999-07-04"},
	}

	for _, student := range students {
		student.Detail()
	}
}
