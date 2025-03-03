package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// Warehouse представляет склад
type Warehouse struct {
	Name string `json:"name"`
}

// Product представляет продукт
type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

// Глобальные переменные для хранения складов и продуктов
var (
	warehouses []Warehouse
	products   []Product
	mu         sync.Mutex // Мьютекс для обработки конкурентного доступа
)

// HTMLHandler обслуживает страницу управления складом
func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Warehouse Management</title>
    <link rel="stylesheet" href="/static/styles.css">
</head>
<body>
    <header>
        <h1>Warehouse Management System</h1>
    </header>
    <main>
        <section>
            <h2>Add Warehouse</h2>
            <form id="warehouseForm">
                <input type="text" id="warehouseName" placeholder="Warehouse Name" required>
                <button type="submit">Add</button>
            </form>
        </section>
        <section>
            <h2>Add Product</h2>
            <form id="productForm">
                <input type="text" id="productName" placeholder="Product Name" required>
                <input type="number" id="productQuantity" placeholder="Quantity" required>
                <button type="submit">Add</button>
            </form>
        </section>
        <section>
            <h2>Warehouses</h2>
            <table id="warehouseTable">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody id="warehouseList"></tbody>
            </table>
        </section>
        <section>
            <h2>Products</h2>
            <table id="productTable">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Quantity</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody id="productList"></tbody>
            </table>
        </section>
    </main>
    <script src="/static/script.js"></script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	tmplParsed, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmplParsed.Execute(w, nil); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// AddWarehouse обрабатывает добавление нового склада
func AddWarehouse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var warehouse Warehouse
	if err := json.NewDecoder(r.Body).Decode(&warehouse); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	warehouses = append(warehouses, warehouse)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(warehouse)
}

// AddProduct обрабатывает добавление нового продукта
func AddProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var product Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	mu.Lock()
	products = append(products, product)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// Main функция для запуска сервера
func main() {
	http.HandleFunc("/", HTMLHandler)
	http.HandleFunc("/api/warehouses", AddWarehouse)
	http.HandleFunc("/api/products", AddProduct)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
