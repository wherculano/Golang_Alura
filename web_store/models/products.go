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
	selectAllProducts, err := db.Query("select * from products order by id asc")
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

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		all_products = append(all_products, product)
	}
	defer db.Close()
	return all_products
}

func CreateNewProduct(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	insertData, err := db.Prepare("insert into products (name, description, price, amount) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProductById(id string) {
	db := db.ConnectDB()
	deleteData, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteData.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	product, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for product.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := product.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Amount = amount

	}

	defer db.Close()
	return productToUpdate
}

func UpdateProductFields(id int, name, description string, price float64, amount int) {
	db := db.ConnectDB()
	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, amount, id)
	defer db.Close()
}
