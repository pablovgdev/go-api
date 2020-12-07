package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAt   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

// Products list
type Products []*Product

// ToJSON encode product list
func (products *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(products)
}

// FromJSON decode product
func (product *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(product)
}

// GetProducts list
func GetProducts() Products {
	return productList
}

// AddProduct to product list
func AddProduct(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)
}

// UpdateProduct list
func UpdateProduct(id int, product *Product) error {
	pos, err := findProductPosition(id)

	if err != nil {
		return err
	}

	product.ID = id
	productList[pos] = product

	return nil
}

func findProductPosition(id int) (int, error) {
	for pos, product := range productList {
		if product.ID == id {
			return pos, nil
		}
	}

	return 0, fmt.Errorf("Product not found")
}

func getNextID() int {
	return len(productList) + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
	},
}
