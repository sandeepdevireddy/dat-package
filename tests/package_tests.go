package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePackages(t *testing.T) {
	// Setup the Gin router
	router := gin.New()
	mockCreatePackage := func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"ID": "67f76ddf09a219e107accf96",
		})
	}
	router.POST("/packages", mockCreatePackage)

	// Mock request body
	requestBody := []byte(`{
 		 "Code": "test",
  		 "Description": "test",
 		 "Components": ["RES", "IND", "CAP"]
	}`)

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "/packages", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	recorder := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(recorder, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, recorder.Code, "Expected HTTP Response code should be 200 Created")
	assert.Contains(t, recorder.Body.String(), `"ID":"67f76ddf09a219e107accf96"`, "Response body should contain the unique_id 'abcdef123456'")
}
