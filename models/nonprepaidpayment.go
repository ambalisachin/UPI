package models

// type AgentDetails struct {
// 	AgentID    string      `json:"agentId"`
// 	DeviceTags []DeviceTag `json:"deviceTags"`
// }

// type DeviceTag struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type AmountDetails struct {
// 	Amount         string `json:"amount"`
// 	Currency       string `json:"currency"`
// 	CustConvFee    string `json:"custConvFee"`
// 	CouCustConvFee string `json:"couCustConvFee"`
// }

// type BillDetails struct {
// 	BillerID       string          `json:"billerId"`
// 	CustomerParams []CustomerParam `json:"customerParams"`
// }

// type CustomerParam struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type CustDetails struct {
// 	CustomerTags []CustomerTag `json:"customerTags"`
// 	MobileNo     string        `json:"mobileNo" gorm:"primaryKey foreignKey:ChID;references:ChID"`
// }

// type CustomerTag struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type PaymentDetails struct {
// 	PaymentInfo []PaymentInfo `json:"paymentInfo"`
// 	PaymentMode string        `json:"paymentMode"`
// 	QuickPay    string        `json:"quickPay"`
// 	SplitPay    string        `json:"splitPay"`
// 	OffusPay    string        `json:"offusPay"`
// }

// type PaymentInfo struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type ReqtBody struct {
// 	AgentDetails    AgentDetails   `json:"agentDetails"`
// 	AmountDetails   AmountDetails  `json:"amountDetails"`
// 	BillDetails     BillDetails    `json:"billDetails"`
// 	ChId            int            `json:"chId"`
// 	CustDetails     CustDetails    `json:"custDetails"`
// 	PaymentDetails  PaymentDetails `json:"paymentDetails"`
// 	RefId           string         `json:"refId"`
// 	ClientRequestId string         `json:"clientRequestId"`
// }

// type SuccessResponse struct {
// 	RespCode string       `json:"respCode"`
// 	Status   string       `json:"status"`
// 	Response ResponseData `json:"response"`
// }

// type FailureResponse struct {
// 	RespCode string              `json:"respCode"`
// 	Status   string              `json:"status"`
// 	Response FailureResponseData `json:"response"`
// }

// type ResponseData struct {
// 	ChID           int    `json:"chId"`
// 	RefID          string `json:"refId"`
// 	ApprovalRefNum string `json:"approvalRefNum"`
// 	ResponseCode   string `json:"responseCode"`
// 	ResponseReason string `json:"responseReason"`
// 	TxnDateTime    string `json:"txnDateTime"`
// 	TxnReferenceID string `json:"txnReferenceId"`
// }

// type FailureResponseData struct {
// 	ChID             int    `json:"chId"`
// 	RefID            string `json:"refId"`
// 	ApprovalRefNum   string `json:"approvalRefNum"`
// 	ResponseCode     string `json:"responseCode"`
// 	ResponseReason   string `json:"responseReason"`
// 	ComplianceReason string `json:"complianceReason"`
// 	ComplianceRespCd string `json:"complianceRespCd"`
// 	TxnDateTime      string `json:"txnDateTime"`
// 	TxnReferenceID   string `json:"txnReferenceId"`
// }
