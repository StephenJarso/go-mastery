package restapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestockProduct(t *testing.T) {
	repo := NewProductRepository()
	router := SetupRouter(repo)
	RegisterPracticeRoute(router, repo)

	// Seed product
	prod, _ := repo.Create(Product{
		SKU:      "SKU-ABC",
		Name:     "Initial Product",
		Price:    10.00,
		Quantity: 5,
	})

	// Case 1: Valid restock
	payload := `{"amount": 10}`
	req := httptest.NewRequest("POST", "/api/v1/products/"+prod.ID+"/restock", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}

	var updatedProduct Product
	json.Unmarshal(rr.Body.Bytes(), &updatedProduct)
	if updatedProduct.Quantity != 15 {
		t.Errorf("expected updated quantity to be 15, got %d", updatedProduct.Quantity)
	}

	// Case 2: Invalid restock (amount = 0)
	payloadBad := `{"amount": 0}`
	reqBad := httptest.NewRequest("POST", "/api/v1/products/"+prod.ID+"/restock", bytes.NewBufferString(payloadBad))
	reqBad.Header.Set("Content-Type", "application/json")
	rrBad := httptest.NewRecorder()
	router.ServeHTTP(rrBad, reqBad)

	if rrBad.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", rrBad.Code)
	}

	// Case 3: Invalid product ID
	reqNoProd := httptest.NewRequest("POST", "/api/v1/products/prod-999/restock", bytes.NewBufferString(payload))
	reqNoProd.Header.Set("Content-Type", "application/json")
	rrNoProd := httptest.NewRecorder()
	router.ServeHTTP(rrNoProd, reqNoProd)

	if rrNoProd.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", rrNoProd.Code)
	}
}
