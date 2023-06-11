package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/apkriege/outdoorsy-backend/common/db"
	"github.com/apkriege/outdoorsy-backend/controllers/rentals"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestEnvironment() {
	err := db.Init()
	if err != nil {
		panic(err)
	}
}

func TestGetRental(t *testing.T) {
	SetupTestEnvironment()
	router := gin.Default()

	// Define the API route and handler
	router.GET("/rentals/:id", rentals.GetRental)

	// Make an API request to retrieve a user
	req, err := http.NewRequest("GET", "/rentals/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use httptest package to send the request to the server
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Verify the response status code
	assert.Equal(t, http.StatusOK, rr.Code, "unexpected status code")
}

func TestGetRentals(t *testing.T) {
	// 	// Create a new Gin router
	SetupTestEnvironment()
	router := gin.Default()
	router.GET("/rentals", rentals.GetRentals)

	// Create a new HTTP request for the /rentals endpoint
	req, err := http.NewRequest(http.MethodGet, "/rentals", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set query parameters if needed
	q := req.URL.Query()
	q.Add("near", "33.64,-117.93")
	q.Add("price_min", "9000")
	q.Add("price_max", "75000")
	q.Add("limit", "3")
	q.Add("offset", "6")
	q.Add("sort", "price")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rr.Code, "unexpected status code")
}
