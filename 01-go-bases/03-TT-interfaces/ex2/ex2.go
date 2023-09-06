package ex2

import "fmt"

// Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
// La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

// Y los costos adicionales son:
// Pequeño: solo tiene el costo del producto
// Mediano: el precio del producto + un 3% de mantenerlo en la tienda
// Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

// El porcentaje de mantenerlo en la tienda es en base al precio del producto.
// El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

// Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.
// Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga

const (
	smallType  = "SMALL"
	mediumType = "MEDIUM"
	largeType  = "LARGE"
)

type Product interface {
	Price() float64
}

type BaseProduct struct {
	price float64
}

func (p *BaseProduct) Price() float64 {
	return p.price
}

type SmallProduct struct {
	Product
}

type MediumProduct struct {
	Product
}

func (p *MediumProduct) Price() float64 {
	return p.Product.Price() + p.Product.Price()*0.03
}

type LargeProduct struct {
	Product
}

func (p *LargeProduct) Price() float64 {
	return p.Product.Price() + p.Product.Price()*0.06 + 2500.0
}

func newProduct(prodType string, price float64) Product {
	baseProduct := &BaseProduct{price: price}
	switch prodType {
	case smallType:
		return &SmallProduct{Product: baseProduct}
	case mediumType:
		return &MediumProduct{Product: baseProduct}
	case largeType:
		return &LargeProduct{Product: baseProduct}
	default:
		return nil
	}
}

func Ex2() {
	fmt.Println()

	var miProducto Product

	miProducto = newProduct(smallType, 1000)
	fmt.Println(miProducto.Price())
	miProducto = newProduct(mediumType, 1000)
	fmt.Println(miProducto.Price())
	miProducto = newProduct(largeType, 1000)
	fmt.Println(miProducto.Price())
}
