package controllers

// import (
// 	"BBT/config"
// 	"BBT/models"
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func BillValidationAPI(c *gin.Context) {
// 	var requestBody models.ReqBody

// 	data, _ := c.Get("decryptedText")
// 	json.Unmarshal(data.([]byte), &requestBody)
// 	db := config.Database.ConnectToDB()
// 	defer db.Close()

// 	// Insert data into the table.
// 	_, err = db.Exec(
// 		"INSERT INTO billvalidations (agentId, billerId, registeredMobileNumber) VALUES (?, ?, ?)",
// 		requestBody.AgentDetails.AgentId,
// 		requestBody.BillDetails.BillerID,
// 		requestBody.BillDetails.CustomerParams[0].Value,
// 	)
// 	if err != nil {
// 		log.Fatal(err)

// 	}
// 	// query := "SELECT * FROM your_table WHERE your_column = ?"
// 	// rows, err := config.db.Query(query, requestBody.AgentDetail.AgentID)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 	return
// 	// }
// 	// defer rows.Close()

// 	var response interface{}
// 	if requestBody.BillDetails.BillerID == "" {
// 		response = models.BillValidationResponse{
// 			RespCode: "1",
// 			Status:   "SUCCESS",
// 			Response: models.BillValidationData{
// 				ChID:             1,
// 				ApprovalRefNum:   "",
// 				ResponseCode:     "000",
// 				ResponseReason:   "Successful",
// 				ComplianceReason: "",
// 			},
// 		}
// 	} else {
// 		response = models.BillValidationResponse{
// 			RespCode: "0",
// 			Status:   "FAILURE",
// 			Response: models.BillValidationData{
// 				ChID:             1,
// 				ApprovalRefNum:   "AB123456",
// 				ResponseCode:     "200",
// 				ResponseReason:   "FAILURE",
// 				ComplianceReason: "Bill response not received",
// 			},
// 		}
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // func BillValidate(c *gin.Context) {
// // 	var validationRequest models.Validaterequest

// // 	data, _ := c.Get("decryptedText")
// // 	json.Unmarshal(data.([]byte), &validationRequest)
// // 	db := config.Database.ConnectToDB()
// // 	defer db.Close()

// // 	if validationRequest.RequestedAmount == 1000 {
// // 		c.JSON(http.StatusOK, gin.H{"message": "Amount Validated  successfully"})
// // 	} else {
// // 		c.JSON(http.StatusOK, gin.H{"message": "Invalid amount"})
// // 	}

// // Convert the amount from string to float64 for comparison
// // requestedAmount, err := strconv.ParseFloat(validationRequest.AmountDetails.Amount, 64)
// // if err != nil {
// // 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
// // 	return
// // }

// // Your validation logic for the amount
// // if requestedAmount != 1000 {
// // 	//  if requestedAmount != ValidationRequest.AmountDetails.Amount {

// // 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
// // 	return
// // }

// // Query the database to get the stored amount
// // 	var storedAmount models.Amount
// // 	if err := config.Database.Last(&storedAmount).Error; err != nil {
// // 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// // 		return
// // 	}

// // 	// Compare the requested amount with the stored amount
// // 	if validationRequest.RequestedAmount == int(storedAmount.Value) {
// // 		c.JSON(http.StatusOK, gin.H{"message": "Validated amount successfully"})
// // 	} else {
// // 		c.JSON(http.StatusOK, gin.H{"message": "Invalid amount"})
// // 	}
// // }

// // c.JSON(http.StatusOK, gin.H{"message": "Amount validation successful"})
// // c.JSON(http.StatusCreated, AESEncrypt("ValidateBill  Successfully.....", []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))
// // }
