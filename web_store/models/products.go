package models

import "Alura/web_store/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
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
		var price float64

		// scans all lines from DB and save the variables
		err := selectAllProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		all_products = append(all_products, product)
	}
	defer db.Close()
	return all_products
}

func CreateNewProduct(name, description string, price float64, amount int){
	db := db.ConnectDB()

	insertData, err := db.Prepare("insert into products (name, description, price, amount) values ($1, $2, $3, $4)")

	if err != nil{
		panic(err.Error())
	}

	insertData.Exec(name, description, price, amount)
	defer db.Close()
}
