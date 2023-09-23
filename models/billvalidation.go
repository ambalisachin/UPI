package models

// // Define Golang models to represent the data.

// type AgentDetail struct {
// 	AgentId string `json:"agentId"`
// }

// type ReqBody struct {
// 	AgentDetails AgentDetail `json:"agentDetails"`
// 	BillDetails  BillDetails `json:"billDetails"`
// }

// type ResponsBody struct {
// 	RespCode string `json:"respCode"`
// 	Status   string `json:"status"`
// 	Response struct {
// 		ChId             int    `json:"chId"`
// 		ApprovalRefNum   string `json:"approvalRefNum"`
// 		ResponseCode     string `json:"responseCode"`
// 		ResponseReason   string `json:"responseReason"`
// 		ComplianceReason string `json:"complianceReason"`
// 	} `json:"response"`
// }

// type BillValidationResponse struct {
// 	RespCode string             `json:"respCode"`
// 	Status   string             `json:"status"`
// 	Response BillValidationData `json:"response"`
// }
// type BillValidationData struct {
// 	ChID             int    `json:"chId"`
// 	ApprovalRefNum   string `json:"approvalRefNum"`
// 	ResponseCode     string `json:"responseCode"`
// 	ResponseReason   string `json:"responseReason"`
// 	ComplianceReason string `json:"complianceReason"`
// }

// // Define the MySQL table schema for your data.

// // type BillValidationRequest struct {
// // 	AgentDetail AgentDetail `json:"agentDetail"`
// // 	BillDetails BillDetails `json:"billDetails"`
// // }

// // type AgentDetail struct {
// // 	AgentID string `json:"agentID"`
// // }

// // // type BillDetails struct {
// // // 	BillerID       string          `json:"billerId"`
// // // 	CustomerParams []CustomerParam `json:"customerParams"`
// // // }

// // // type CustomerParam struct {
// // // 	Name  string `json:"name"`
// // // 	Value string `json:"value"`
// // // }

// // type BillValidationResponse struct {
// // 	RespCode string             `json:"respCode"`
// // 	Status   string             `json:"status"`
// // 	Response BillValidationData `json:"response"`
// // }

// // type BillValidationData struct {
// // 	ChID             int    `json:"chId"`
// // 	ApprovalRefNum   string `json:"approvalRefNum"`
// // 	ResponseCode     string `json:"responseCode"`
// // 	ResponseReason   string `json:"responseReason"`
// // 	ComplianceReason string `json:"complianceReason"`
// // }

// // // type Validaterequest struct {
// // // 	RequestedAmount int `json:"requestedAmount"`
// // // }

// // // type Amount struct {
// // // 	gorm.Model
// // // 	Value int
// // // }

// // //////////////////////////////////////////
// // // type ReqBody struct {
// // // 	AgentDetails AgentDetails `json:"agentDetails"`
// // // 	BillDetails  BillDetails  `json:"billDetails"`
// // // }
// // // type AgentDetails struct {
// // // 	AgentId string `json:"agentId"`
// // // }

// // //	type BillDetails struct {
// // //		BillerId       string         `json:"billerId"`
// // //		CustomerParams []CustomerParam `json:"customerParams"`
// // //	}
// // //
// // // Define the MySQL table schema for your data.
