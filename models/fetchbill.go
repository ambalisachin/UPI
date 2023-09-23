package models

// type FetchBillRequest struct {
// 	ChID            int          `gorm:"primaryKey" json:"chId"`
// 	IsRealTimeFetch bool         `json:"isRealTimeFetch"`
// 	CustDetails     CustDetails  `json:"custDetails" `
// 	AgentDetails    AgentDetails `json:"agentDetails"`
// 	BillDetail      BillDetail   `json:"billDetails"`
// }

// type CustDetails struct {
//     MobileNo     string       `json:"mobileNo"`
//     CustomerTags []CustomerTag `json:"customerTags"`
// }

// type CustomerTag struct {
// 	Name  string `json:"name"`
// 	Value string `json:"value"`
// }

// type AgentDetails struct {
//     AgentID     string       `json:"agentId"`
//     DeviceTags  []DeviceTag  `json:"deviceTags"`
// }

// type DeviceTag struct {
//     Name  string `json:"name"`
//     Value string `json:"value"`
// }

// type BillDetail struct {
// 	BillerID       string          `json:"billerId"`
// 	CustomerParams []CustomerParam `json:"customerParams"`
// }

// type CustomerParam struct {
//     Name  string `json:"name"`
//     Value string `json:"value"`
// }

type FetchBillResponse struct {
	RespCode string   `json:"respCode"`
	Status   string   `json:"status"`
	Response RespData `json:"response"`
}

type RespData struct {
	ChID             int             `json:"chId"`
	RefID            string          `json:"refId"`
	ApprovalRefNum   string          `json:"approvalRefNum"`
	ResponseCode     string          `json:"responseCode"`
	ResponseReason   string          `json:"responseReason"`
	ComplianceReason string          `json:"complianceReason"`
	ComplianceRespCd string          `json:"complianceRespCd"`
	BillDetails      []CustomerParam `json:"billDetails"`
	BillerResponse   BillerResponse  `json:"billerResponse"`
	AdditionalInfo   []interface{}   `json:"additionalInfo"`
}

type BillParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BillerResponse struct {
	CustomerName string        `json:"customerName"`
	Amount       string        `json:"amount"`
	DueDate      string        `json:"dueDate"`
	CustConvFee  string        `json:"custConvFee"`
	CustConvDesc string        `json:"custConvDesc"`
	BillDate     string        `json:"billDate"`
	BillNumber   string        `json:"billNumber"`
	BillPeriod   string        `json:"billPeriod"`
	BillTags     []interface{} `json:"billTags"`
}

// type FetchBillRequest struct {
// 	ChID            int          `gorm:"primaryKey" json:"chId"`
// 	IsRealTimeFetch bool         `json:"isRealTimeFetch"`
// 	CustDetailsID   uint         // This acts as a foreign key for CustDetails
// 	CustDetails     CustDetails  `gorm:"foreignKey:CustDetailsID" json:"custDetails"`
// 	AgentDetails    AgentDetails `json:"agentDetails"`
// 	BillDetail      BillDetails  `json:"billDetails"`
// }

// type CustDetails struct {
// 	ID           uint           `gorm:"primaryKey"`
// 	CustomerTags []customerTags `gorm:"foreignKey:CustDetailsID" json:"customerTags"`
// 	MobileNo     string         `json:"mobileNo"`
// }

// type FetchBillRequest struct {
// 	ChID            int          `gorm:"primaryKey" json:"chId"`
// 	IsRealTimeFetch bool         `json:"isRealTimeFetch"`
// 	CustDetailsID   uint         // This acts as a foreign key for CustDetails
// 	CustDetails     CustDetails  `gorm:"foreignKey:CustDetailsID" json:"custDetails"`
// 	AgentDetailsID  uint         // This acts as a foreign key for AgentDetails
// 	AgentDetails    AgentDetails `gorm:"foreignKey:AgentDetailsID" json:"agentDetails"`
// 	BillDetailID    uint         // This acts as a foreign key for BillDetail
// 	BillDetail      BillDetail   `gorm:"foreignKey:BillDetailID" json:"billDetails"`
// }
// type CustDetails struct {
// 	ID           uint           `gorm:"primaryKey"`          // Primary key for CustDetails
// 	CustomerTags []customerTags `gorm:"foreignKey:DetailID"` // Foreign key relationship
// 	MobileNo     string         `json:"mobileNo"`
// 	DetailID     uint           // This acts as a reference for CustomerTag
// }
// type customerTags struct {
// 	ID       uint   `gorm:"primaryKey"` // Primary key for CustomerTag
// 	Name     string `json:"name"`
// 	Value    string `json:"value"`
// 	DetailID uint   // This will act as the foreign key referring to CustDetails
// }

// type AgentDetails struct {
// 	ID         uint         `gorm:"primaryKey"`
// 	AgentID    string       `json:"agentId"`
// 	DeviceTags []DeviceTags `gorm:"foreignKey:AgentDetailsID" json:"deviceTags"`
// }

// type DeviceTags struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	Name     string `json:"name"`
// 	Value    string `json:"value"`
// 	DetailID uint   // This will act as the foreign key referring to DeviceTags
// }

// type BillDetail struct {
// 	ID             uint             `gorm:"primaryKey"`
// 	BillerID       string           `json:"billerId"`
// 	CustomerParams []CustomerParams `gorm:"foreignKey:BillDetailID" json:"customerParams"`
// }

// type CustomerParams struct {
// 	ID           uint   `gorm:"primaryKey"`
// 	Name         string `json:"name"`
// 	Value        string `json:"value"`
// 	BillDetailID uint   // This will act as the foreign key referring to BillDetail

// }

// type CustomerTag struct {
// 	ID            uint   `gorm:"primaryKey"`
// 	Name          string `json:"name"`
// 	Value         string `json:"value"`
// 	CustDetailsID uint   // This will act as the foreign key referring to CustDetails
// }
//////////////////////////////////////
// type FetchBillRequest struct {
// 	ChID            int          `gorm:"primaryKey" json:"chId"`
// 	IsRealTimeFetch bool         `json:"isRealTimeFetch"`
// 	CustDetailsID   uint         // This acts as a foreign key for CustDetails
// 	CustDetails     CustDetails  `gorm:"foreignKey:CustDetailsID" json:"custDetails"`
// 	AgentDetailsID  uint         // This acts as a foreign key for AgentDetails
// 	AgentDetails    AgentDetails `gorm:"foreignKey:AgentDetailsID" json:"agentDetails"`
// 	BillDetailID    uint         // This acts as a foreign key for BillDetail
// 	BillDetail      BillDetail   `gorm:"foreignKey:BillDetailID" json:"billDetails"`
// }
// type CustDetails struct {
// 	ID           uint           `gorm:"primaryKey"`          // Primary key for CustDetails
// 	CustomerTags []customerTags `gorm:"foreignKey:DetailID"` // Foreign key relationship
// 	MobileNo     string         `json:"mobileNo"`
// 	DetailID     uint           // This acts as a reference for CustomerTag
// }
// type customerTags struct {
// 	ID       uint   `gorm:"primaryKey"` // Primary key for CustomerTag
// 	Name     string `json:"name"`
// 	Value    string `json:"value"`
// 	DetailID uint   // This will act as the foreign key referring to CustDetails
// }

// type AgentDetails struct {
// 	ID             uint         `gorm:"primaryKey"`
// 	AgentID        string       `json:"agentId"`
// 	DeviceTags     []DeviceTags `gorm:"foreignKey:AgentDetailsID" json:"deviceTags"`
// 	AgentDetailsID uint
// }

// type DeviceTags struct {
// 	ID             uint   `gorm:"primaryKey"`
// 	Name           string `json:"name"`
// 	Value          string `json:"value"`
// 	AgentDetailsID uint
// 	DetailID       uint // This will act as the foreign key referring to DeviceTags
// }

// type BillDetail struct {
// 	ID             uint             `gorm:"primaryKey"`
// 	BillerID       string           `json:"billerId"`
// 	CustomerParams []CustomerParams `gorm:"foreignKey:BillDetailID" json:"customerParams"`
// 	BillDetailID   uint
// }

// type CustomerParams struct {
// 	ID           uint   `gorm:"primaryKey"`
// 	Name         string `json:"name"`
// 	Value        string `json:"value"`
// 	BillDetailID uint   // This will act as the foreign key referring to BillDetail

// }
type CustomerTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CustDetails struct {
	MobileNo     string        `json:"mobileNo"`
	CustomerTags []CustomerTag `json:"customerTags"`
}

type DeviceTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AgentDetails struct {
	AgentId    string      `json:"agentId"`
	DeviceTags []DeviceTag `json:"deviceTags"`
}

type CustomerParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type BillDetails struct {
	BillerId       string          `json:"billerId"`
	CustomerParams []CustomerParam `json:"customerParams"`
}

type FetchBillRequest struct {
	ChId            int          `json:"chId"`
	IsRealTimeFetch bool         `json:"isRealTimeFetch"`
	CustDetails     CustDetails  `json:"custDetails"`
	AgentDetails    AgentDetails `json:"agentDetails"`
	BillDetails     BillDetails  `json:"billDetails"`
}

// YourModel represents the table structure in the database
// type YourModel struct {
// 	gorm.Model // This includes the default fields: ID, CreatedAt, UpdatedAt, DeletedAt
// }
