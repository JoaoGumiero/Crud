package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/joaog/API_Test/models"
)

// trazer todos os templates da pasta
var temp = template.Must(template.ParseGlob("templates/*.html"))

// Arrumar essa função aqui junto com o HTML
func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}


func NewProductHandler(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func InsertProductHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price") // Here, the price and amount fields are retrieved in the string format, so we have to parse them to their respective field types.
		amount := r.FormValue("amount")

		priceFloated, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Failed to convert the price object", err)
		}
		amountInted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Failed to convert the amount object", err)
		}
		models.CreateProduct(name, description, priceFloated, amountInted)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	models.DeleteProduct(productID)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func EditProductHandler(w http.ResponseWriter, r *http.Request) {
	productID := r.URL.Query().Get("id")
	product := models.EditProduct(productID)
	temp.ExecuteTemplate(w, "Edit", product)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price") // Here, the price and amount fields are retrieved in the string format, so we have to parse them to their respective field types.
		amount := r.FormValue("amount")

		idInted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Failed to convert the amount object", err)
		}
		priceFloated, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Failed to convert the price object", err)
		}
		amountInted, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Failed to convert the amount object", err)
		}
		models.UpdateProduct(idInted,name, description, priceFloated, amountInted)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}