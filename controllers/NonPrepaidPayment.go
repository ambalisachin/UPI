package controllers

// import (
// 	"BBT/config"
// 	"BBT/models"
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// var err error

// func NonPrepaidAPI(c *gin.Context) {
// 	var req models.ReqtBody
// 	data, _ := c.Get("decryptedText")
// 	json.Unmarshal(data.([]byte), &req)
// 	db := config.Database.ConnectToDB()
// 	defer db.Close()

// 	// Insert data into the MySQL table

// 	insertQuery := `INSERT INTO nonprepaidpayments (agentId,initiatingChannel,mobile,geocode,postalCode,terminalId,amount,currency,custConvFee,couCustConvFee,billerId,consumerMobileNo,uid,accountNo,chId,email,mobileNo,remarks,paymentMode,quickPay,splitPay,offusPay,refId,clientRequestId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

// 	_, err = db.Exec(insertQuery,
// 		req.AgentDetails.AgentID, req.AgentDetails.DeviceTags[0].Value, req.AgentDetails.DeviceTags[1].Value, req.AgentDetails.DeviceTags[2].Value, req.AgentDetails.DeviceTags[3].Value, req.AgentDetails.DeviceTags[4].Value, req.AmountDetails.Amount, req.AmountDetails.Currency, req.AmountDetails.CustConvFee, req.AmountDetails.CouCustConvFee, req.BillDetails.BillerID, req.BillDetails.CustomerParams[0].Value,
// 		req.BillDetails.CustomerParams[1].Value, req.BillDetails.CustomerParams[2].Value, req.ChId, req.CustDetails.CustomerTags[0].Value, req.CustDetails.MobileNo, req.PaymentDetails.PaymentInfo[0].Value, req.PaymentDetails.PaymentMode, req.PaymentDetails.QuickPay, req.PaymentDetails.SplitPay, req.PaymentDetails.OffusPay, req.RefId, req.ClientRequestId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 	Query the database
// 	// query := "SELECT * FROM your_table WHERE your_column = ?"
// 	// rows, err := config.DB.Query(query, req.AgentDetails.AgentID)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 	return
// 	// }
// 	// defer rows.Close()

// 	// Construct the response based on success or failure.
// 	var response interface{}
// 	if req.AmountDetails.Amount != "" {
// 		response = models.SuccessResponse{
// 			RespCode: "1",
// 			Status:   "SUCCESS",
// 			Response: models.ResponseData{
// 				ChID:           req.ChId,
// 				RefID:          "LZ2EAV1BVVZ0JJQU05MA3RA59IN21301822",
// 				ApprovalRefNum: "12345093",
// 				ResponseCode:   "000",
// 				ResponseReason: "Successful",
// 				TxnDateTime:    "2022-05-10 18:22:25",
// 				TxnReferenceID: "AM012130EW74Y0R0D39C",
// 			},
// 		}
// 	} else {
// 		response = models.FailureResponse{
// 			RespCode: "0",
// 			Status:   "FAILURE",
// 			Response: models.FailureResponseData{
// 				ChID:             req.ChId,
// 				RefID:            "8HVJDD3ZQRSBA6N67P3S8M8G5SJ21301824",
// 				ApprovalRefNum:   "",
// 				ResponseCode:     "200",
// 				ResponseReason:   "Failure",
// 				ComplianceReason: "Payment received for the billing period - no bill due",
// 				ComplianceRespCd: "BFR004",
// 				TxnDateTime:      "2022-05-10 18:25:10",
// 				TxnReferenceID:   "AM0121308RXD8X7O91UR",
// 			},
// 		}
// 	}

// 	c.JSON(http.StatusOK, response)
// }
