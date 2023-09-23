package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"BBT/config"
// 	"BBT/models"

// 	"github.com/gin-gonic/gin"
// )

// func ValidateAPI(c *gin.Context) {
// 	var req models.RequestBodys
// 	// if err := c.ShouldBindJSON(&req); err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	// 	return
// 	// }
// 	body, _ := io.ReadAll(c.Request.Body)
// 	json.Unmarshal(body, &req)
// 	db := config.Database.ConnectToDB()
// 	defer db.Close()

// 	_, e := db.Query("CREATE TABLE IF NOT EXISTS validationbills(agentId int ,billerId varchar(20),name varchar(20),value varchar(20))")
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// _, err := db.Query("insert into billers(agentId,billerId,name, value) values(?,?,?,?)", req.agentId, req.billerId, req.name, req.value)

// if err != nil {
// 	c.JSON(http.StatusBadRequest, err)
// 	return
// }
// Perform validation logic here based on reqBody
// Replace the validation logic below with your specific requirements
// var response interface{}
// if req.AgentDetails.AgentID == "AM01YB41BSC519046456" && req.BillDetails.BillerID == "OUSH40000NAT02" {
// 	// Validation passed
// 	var response = models.ResponseBodys{
// 		RespCode: "1",
// 		Status:   "SUCCESS",
// 		Response: struct {
// 			ChID             int    `json:"chId"`
// 			ApprovalRefNum   string `json:"approvalRefNum"`
// 			ResponseCode     string `json:"responseCode"`
// 			ResponseReason   string `json:"responseReason"`
// 			ComplianceReason string `json:"complianceReason"`
// 		}{
// 			ChID:             1,
// 			ApprovalRefNum:   "",
// 			ResponseCode:     "000",
// 			ResponseReason:   "Successful",
// 			ComplianceReason: "",
// 		},
// 	}
// 	c.JSON(http.StatusOK, response)
// } else {
// Validation failed
// 		response = models.ResponseBodys{
// 			RespCode: "0",
// 			Status:   "FAILURE",
// 			Response: struct {
// 				ChID             int    `json:"chId"`
// 				ApprovalRefNum   string `json:"approvalRefNum"`
// 				ResponseCode     string `json:"responseCode"`
// 				ResponseReason   string `json:"responseReason"`
// 				ComplianceReason string `json:"complianceReason"`
// 			}{
// 				ChID:             1,
// 				ApprovalRefNum:   "AB123456",
// 				ResponseCode:     "200",
// 				ResponseReason:   "FAILURE",
// 				ComplianceReason: "Bill response not received",
// 			},
// 		}
// 		c.JSON(http.StatusOK, response)
// 		c.JSON(http.StatusCreated, AESEncrypt("response", []byte(c.Request.Header.Get("x-key")), c.Request.Header.Get("x-iv")))

// 	}
// }
