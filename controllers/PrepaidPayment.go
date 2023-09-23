package controllers

// func PrepaidAPI(c *gin.Context) {

// 	var req models.ReqstBody
// 	data, _ := c.Get("decryptedText")
// 	json.Unmarshal(data.([]byte), &req)

// 	db := config.InitDB()

// 	db.AutoMigrate(&models.ReqBody{})

// 	if pk := db.Create(&req); pk.Error != nil {
// 		log.Fatal(pk.Error)
// 		return
// 	}
// db := config.Database.ConnectToDB()
// defer db.Close()

// // Insert data into the table
// _, err = db.Exec("INSERT INTO prepaidpayments (chId, refId, clientRequestId, agentId, initiatingChannel, amount, currency, custConvFee, couCustConvFee, billerId, mobileNumber, circle, id, email, aadhaar, pan, mobileNo, remarks, paymentMode, quickPay, splitPay, offusPay, planType, planId) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
// 	req.ChID, req.RefID, req.ClientRequestID, req.AgentDetails.AgentID, req.AgentDetails.DeviceTags[0].Value, req.AmountDetails.Amount, req.AmountDetails.Currency, req.AmountDetails.CustConvFee, req.AmountDetails.CouCustConvFee, req.BillDetails.BillerID, req.BillDetails.CustomerParams[0].Value, req.BillDetails.CustomerParams[1].Value, req.BillDetails.CustomerParams[2].Value, req.CustDetails.CustomerTags[0].Value,
// 	req.CustDetails.CustomerTags[1].Value, req.CustDetails.CustomerTags[2].Value, req.CustDetails.MobileNo, req.PaymentDetails.PaymentInfo[0].Value, req.PaymentDetails.PaymentMode, req.PaymentDetails.QuickPay, req.PaymentDetails.SplitPay, req.PaymentDetails.OffusPay, req.PlanDetails.PlanDetail.Type, req.PlanDetails.PlanDetail.ID,
// )
// if err != nil {
// 	log.Fatal(err)
// }

// query := "SELECT * FROM your_table WHERE your_column = ?"
// rows, err := config.DB.Query(query, requestBody.AgentDetails.AgentID)
// if err != nil {
//     c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//     return
// }
// defer rows.Close()

// Process the query results here.

// Construct the response based on success or failure.
//  var response interface{}
// if req.ReqstBody.ClientRequestID != "" {

// 	var response models.SuccesResponse
// 	result := db.First(&response)

// 	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
// 		// No records exist, so create a new user
// 		response = models.SuccesResponse{
// 			RespCode: "1",
// 			Status:   "SUCCESS",
// 			Response: models.SuccessData{
// 				ChID:           req.ChID,
// 				RefID:          "3W0YSK5MJCQW9QWUEUSB3RKTY0221861853",
// 				ApprovalRefNum: "AB12345001",
// 				ResponseCode:   "000",
// 				ResponseReason: "Successful",
// 				BillerPlanResponse: models.BillerPlanResponse{
// 					PlanInfo: []models.PlanInfo{
// 						{
// 							Type: "ACTIVATED",
// 							PlanInfoTags: []models.PlanInfoTag{
// 								{
// 									Name:  "Id",
// 									Value: "10",
// 								},
// 								{
// 									Name:  "Plan Type",
// 									Value: "Recharge",
// 								},
// 								{
// 									Name:  "Talktime",
// 									Value: "0",
// 								},
// 								{
// 									Name:  "Validity",
// 									Value: "84 Days",
// 								},
// 								{
// 									Name:  "Data",
// 									Value: "100 GB",
// 								},
// 								{
// 									Name:  "Circle",
// 									Value: "Andhra Pradesh",
// 								},
// 								{
// 									Name:  "amountInRupees",
// 									Value: "108",
// 								},
// 								{
// 									Name:  "planDescription",
// 									Value: "Tarrif Calls - Local/STD/LL @ 1P/sec",
// 								},
// 							},
// 						},
// 					},
// 				},
// 				TxnDateTime:    "2022-07-04 13:02:52",
// 				TxnReferenceID: "YB412186NKGZG30XV5RE",
// 			},
// 		}
// 		db.Create(&response)
// 		response = models.SuccesResponse{}
// 		fmt.Println("Inserted new user as there was no data in the User table.")
// 	} else {
// 		fmt.Println("Data already exists in the User table.")
// 	}

// 	db.Find(&response)

// 	c.JSON(http.StatusOK, response)
// }

// // func GetPrepaidResponse(c *gin.Context) {}

// // else {
// // 	response = models.FailureResponseData{
// // 		RespCode: "0",
// // 		Status:   "FAILURE",
// // 		Response: models.FailureResponse{
// // 			ChID:             reqs.ChID,
// // 			RefID:            "refId for failure",
// // 			ApprovalRefNum:   "approvalRefNum for failure",
// // 			ResponseCode:     "failure response code",
// // 			ResponseReason:   "Failure",
// // 			ComplianceReason: "Compliance reason for failure",
// // 			ComplianceRespCd: "ComplianceRespCd for failure",
// // 			TxnDateTime:      "2022-07-04 13:02:52",
// // 			TxnReferenceID:   "TxnReferenceID for failure",
// // 		},
// // 	}

// // }

// // 	c.JSON(http.StatusOK, response)
// // }
