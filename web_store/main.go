package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq" // _ means excute just during it is running
)


func connectDB() *sql.DB{
	conn := "user=postgres dbname=alura_store password=123qweasd host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}


type Product struct{
	Id int
	Name string
	Description string
	Value float64
	Amount int
}


var temp8 = template.Must(template.ParseGlob("templates/*.html"))

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	db := connectDB()
	selectAllProducts, err := db.Query("select * from products")
	if err != nil{
		panic(err.Error())
	}
	
	product := Product{} // instance of one Product
	all_products := []Product{}  // slice of Product

	for selectAllProducts.Next(){
		var id, amount int
		var name, description string
		var value float64
		
		// scans all lines from DB and save the variables
		err := selectAllProducts.Scan(&id, &name, &description, &value, &amount)
		if err != nil{
			panic(err.Error())
		}

		product.Name = name
		product.Description = description
		product.Value = value
		product.Amount = amount

		all_products = append(all_products, product)
	}

	temp8.ExecuteTemplate(w, "Index", all_products)
	defer db.Close() // Executes all lines and then runs this line
}