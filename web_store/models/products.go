package models

import "Alura/web_store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Value       float64
	Amount      int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()
	selectAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}        // instance of one Product
	all_products := []Product{} // slice of Product

	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var value float64

		// scans all lines from DB and save the variables
		err := selectAllProducts.Scan(&id, &name, &description, &value, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Name = name
		product.Description = description
		product.Value = value
		product.Amount = amount

		all_products = append(all_products, product)
	}
	defer db.Close()
	return all_products
}
