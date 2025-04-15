package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/DAT-PACKAGE/db"
	"github.com/DAT-PACKAGE/models"
	"github.com/gin-gonic/gin"
)

// CreatePackage handles the POST /packages route
func CreatePackages(c *gin.Context) {
	log.Println("Creating Package")
	var Package models.Package
	if err := c.ShouldBindJSON(&Package); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.PackagesCollection.InsertOne(context.Background(), Package)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}
