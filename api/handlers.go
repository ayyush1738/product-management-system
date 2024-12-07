package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"product-management-system/db"

	"github.com/gorilla/mux"
)

// Product represents a product entity
type Product struct {
	ID                 int      `json:"id"`
	UserID             int      `json:"user_id"`
	ProductName        string   `json:"product_name"`
	ProductDescription string   `json:"product_description"`
	ProductImages      []string `json:"product_images"`
	CompressedImages   []string `json:"compressed_product_images"`
	ProductPrice       float64  `json:"product_price"`
	CreatedAt          string   `json:"created_at"`
}

// CreateProduct handles the creation of a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO products (user_id, product_name, product_description, product_images, product_price)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`
	var productID int
	err := db.DB.QueryRow(query, product.UserID, product.ProductName, product.ProductDescription, product.ProductImages, product.ProductPrice).Scan(&productID)
	if err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"id":` + strconv.Itoa(productID) + `}`))
}

// GetProductByID retrieves a product by its ID
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product Product
	query := `SELECT * FROM products WHERE id = $1`
	err = db.DB.Get(&product, query, productID)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProductsByUser retrieves all products for a specific user
func GetProductsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	var products []Product
	query := `SELECT * FROM products WHERE user_id = $1`
	err := db.DB.Select(&products, query, userID)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
