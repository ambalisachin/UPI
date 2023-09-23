package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"BBT/models"

	"github.com/gin-gonic/gin"
)

// FetchBill gives billdetails
func FetchBill(c *gin.Context) {
	var reqs models.FetchBillRequest
	data, _ := c.Get("decryptedText")
	json.Unmarshal(data.([]byte), &reqs)

	db := NewDataBase()
	db.Create(&reqs)

	if !db.Validate(&reqs) {

		log.Fatal(errors.New("invalid data"))
		return
	}

	//*********************** this part should be removed after real database is created *******end************//

	resp, err := db.GetBill(&reqs)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
