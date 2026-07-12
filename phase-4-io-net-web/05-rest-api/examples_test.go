package restapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductCRUDWorkflow(t *testing.T) {
	repo := NewProductRepository()
	router := SetupRouter(repo)

	// 1. Create a product (POST)
	createPayload := `{
		"sku": "GO-PLUSH",
		"name": "Original Gopher Plush",
		"price": 19.99,
		"quantity": 10
	}`

	reqCreate := httptest.NewRequest("POST", "/api/v1/products", bytes.NewBufferString(createPayload))
	reqCreate.Header.Set("Content-Type", "application/json")
	rrCreate := httptest.NewRecorder()
	router.ServeHTTP(rrCreate, reqCreate)

	if rrCreate.Code != http.StatusCreated {
		t.Fatalf("expected POST status 201, got %d. Response: %s", rrCreate.Code, rrCreate.Body.String())
	}

	var created Product
	err := json.Unmarshal(rrCreate.Body.Bytes(), &created)
	if err != nil {
		t.Fatalf("failed to parse response body: %v", err)
	}

	if created.ID == "" {
		t.Error("expected generated product ID, got empty string")
	}
	if created.SKU != "GO-PLUSH" {
		t.Errorf("expected SKU 'GO-PLUSH', got %q", created.SKU)
	}

	// 2. Attempt duplicate SKU creation (POST - Conflict 409)
	reqDup := httptest.NewRequest("POST", "/api/v1/products", bytes.NewBufferString(createPayload))
	reqDup.Header.Set("Content-Type", "application/json")
	rrDup := httptest.NewRecorder()
	router.ServeHTTP(rrDup, reqDup)

	if rrDup.Code != http.StatusConflict {
		t.Errorf("expected POST duplicate status 409, got %d", rrDup.Code)
	}

	// 3. Attempt invalid creation (POST - Bad Request 400)
	badPayload := `{"sku": "", "price": -5.00}` // Invalid SKU (empty), invalid price
	reqBad := httptest.NewRequest("POST", "/api/v1/products", bytes.NewBufferString(badPayload))
	reqBad.Header.Set("Content-Type", "application/json")
	rrBad := httptest.NewRecorder()
	router.ServeHTTP(rrBad, reqBad)

	if rrBad.Code != http.StatusBadRequest {
		t.Errorf("expected POST invalid status 400, got %d", rrBad.Code)
	}

	// 4. Retrieve the product by ID (GET)
	reqGet := httptest.NewRequest("GET", "/api/v1/products/"+created.ID, nil)
	rrGet := httptest.NewRecorder()
	router.ServeHTTP(rrGet, reqGet)

	if rrGet.Code != http.StatusOK {
		t.Errorf("expected GET status 200, got %d", rrGet.Code)
	}

	// 5. List all products (GET)
	reqList := httptest.NewRequest("GET", "/api/v1/products", nil)
	rrList := httptest.NewRecorder()
	router.ServeHTTP(rrList, reqList)

	if rrList.Code != http.StatusOK {
		t.Errorf("expected GET list status 200, got %d", rrList.Code)
	}

	var list []Product
	json.Unmarshal(rrList.Body.Bytes(), &list)
	if len(list) != 1 {
		t.Errorf("expected list length 1, got %d", len(list))
	}

	// 6. Update the product (PUT)
	updatePayload := `{
		"price": 24.99,
		"quantity": 15
	}`
	reqUpdate := httptest.NewRequest("PUT", "/api/v1/products/"+created.ID, bytes.NewBufferString(updatePayload))
	reqUpdate.Header.Set("Content-Type", "application/json")
	rrUpdate := httptest.NewRecorder()
	router.ServeHTTP(rrUpdate, reqUpdate)

	if rrUpdate.Code != http.StatusOK {
		t.Errorf("expected PUT status 200, got %d", rrUpdate.Code)
	}

	var updated Product
	json.Unmarshal(rrUpdate.Body.Bytes(), &updated)
	if updated.Price != 24.99 || updated.Quantity != 15 {
		t.Errorf("expected updated values (price=24.99, quantity=15), got (price=%f, quantity=%d)", updated.Price, updated.Quantity)
	}

	// 7. Delete the product (DELETE)
	reqDelete := httptest.NewRequest("DELETE", "/api/v1/products/"+created.ID, nil)
	rrDelete := httptest.NewRecorder()
	router.ServeHTTP(rrDelete, reqDelete)

	if rrDelete.Code != http.StatusNoContent {
		t.Errorf("expected DELETE status 204 No Content, got %d", rrDelete.Code)
	}

	// 8. Verify product is gone (GET -> 404)
	reqVerify := httptest.NewRequest("GET", "/api/v1/products/"+created.ID, nil)
	rrVerify := httptest.NewRecorder()
	router.ServeHTTP(rrVerify, reqVerify)

	if rrVerify.Code != http.StatusNotFound {
		t.Errorf("expected GET post-delete status 404, got %d", rrVerify.Code)
	}
}
