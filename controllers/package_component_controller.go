package controllers

import (
	"context"
	"log"
	"net/http"

	"sandeepdevireddy/dat-package/db"
	"sandeepdevireddy/dat-package/models"

	"github.com/gin-gonic/gin"
)

// CreatePackage handles the POST /package/components route
func CreatePackageComponents(c *gin.Context) {
	log.Println("Creating Package Components")
	var PackageComponent models.PackageComponent
	if err := c.ShouldBindJSON(&PackageComponent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.PackageComponentsCollection.InsertOne(context.Background(), PackageComponent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
