package ex2

import (
	"fmt"
	"os"
)

// A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
// Ahora que el archivo sí existe, el panic no debe ser lanzado.
// 1. Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
// 2. Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga. En el caso de no poder leerlo, se debe lanzar un “panic”.
// Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.

func Ex2() {
	file, err := os.Open("./customers-ex2.txt")

	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}

	bytes := make([]byte, 32)
	bytesRead, err := file.Read(bytes)

	defer func() {
		file.Close()
		fmt.Println("Ejecución finalizada")
	}()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes[:bytesRead]))
}
