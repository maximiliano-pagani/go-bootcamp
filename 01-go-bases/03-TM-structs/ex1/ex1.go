package ex1

import "fmt"

// Crear un programa que cumpla los siguiente puntos:
// Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
// Tener un slice global de Product llamado Products instanciado con valores.
// 2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
// Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
// Ejecutar al menos una vez cada método y función definido desde main().

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p *Product) GetAll() {
	for _, product := range Products {
		product.Print()
	}
}

func (p *Product) Print() {
	fmt.Println(p.ID, p.Name, p.Price, p.Description, p.Category)
}

func getById(id int) {
	for _, product := range Products {
		if product.ID == id {
			fmt.Println("Found with ID", id, ":")
			product.Print()
		}
	}
}

var Products []Product

func Ex1() {
	p := Product{ID: 1, Name: "Apple", Price: 84}
	p.Save()

	p = Product{ID: 2, Name: "Banana", Price: 543}
	p.Save()

	p = Product{ID: 3, Name: "Orange", Price: 161}
	p.Save()

	p.GetAll()
	getById(2)
}
