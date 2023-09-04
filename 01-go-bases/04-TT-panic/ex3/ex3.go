package ex3

import (
	"errors"
	"fmt"
)

// El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos son:
// - Legajo
// - Nombre
// - DNI
// - Número de teléfono
// - Domicilio

// - Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array. En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
// 	1.- generar un panic;
// 	2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución del programa normalmente.
// - Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
// - Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “Fin de la ejecución” y “Se detectaron varios errores en tiempo de ejecución”. Utilizá defer para cumplir con este requerimiento.

// - Requerimientos generales:
// Utilizá recover para recuperar el valor de los panics que puedan surgir
// Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
// Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error retornado).

type Client struct {
	id      int
	name    string
	dni     int
	tel     string
	address string
}

var ErrInvalidClientData error = errors.New("Error: los datos del cliente no son válidos")

func (c *Client) isClientDataValid() (bool, error) {
	if c.id == 0 ||
		c.name == "" ||
		c.dni == 0 ||
		c.tel == "" ||
		c.address == "" {
		return false, ErrInvalidClientData
	}

	return true, nil
}

var clientsDatabase []Client = []Client{
	{
		id:      1,
		name:    "Pepe",
		dni:     23949221,
		tel:     "1123889212",
		address: "J. B. Justo 1754",
	},
}

func Ex3() {
	newClientsSampleData := []Client{
		{
			id:      1,
			name:    "Pepe",
			dni:     23949221,
			tel:     "1123889212",
			address: "J. B. Justo 1754",
		},
		{
			id:   3,
			name: "Mario",
		},
		{
			id:      3,
			name:    "Mario",
			dni:     4892392,
			tel:     "1184289235",
			address: "Av. Cabildo 934",
		},
	}

	for _, client := range newClientsSampleData {
		addClient(&client)
	}

}

func addClient(newClient *Client) {
	fmt.Println("\nAgregando nuevo cliente", newClient)
	doesClientAlreadyExists(newClient)

	defer func() {
		fmt.Println("Fin de la ejecución")

		err := recover()

		if err != nil {
			fmt.Println(err)
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
	}()

	_, err := newClient.isClientDataValid()

	if errors.Is(err, ErrInvalidClientData) {
		panic(err)
	}

	fmt.Println("Cliente agregado")
}

func doesClientAlreadyExists(newClient *Client) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, client := range clientsDatabase {
		if *newClient == client {
			panic("Error: el cliente ya existe")
		}
	}
}
