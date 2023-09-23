package controllers

import (
	"BBT/config"
	"BBT/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// IBillService provides methods to get, create and validate billdata
type IBillService interface {
	Create(reqs *models.FetchBillRequest) error
	GetBill(reqs *models.FetchBillRequest) (*models.FetchBillResponse, error)
	Validate(reqs *models.FetchBillRequest) bool
}

// DataBase gives gorm databases
type DataBase struct {
	db *gorm.DB
}

// Create pushes billdata into database
func (d *DataBase) Create(reqs *models.FetchBillRequest) error {
	d.db.AutoMigrate(&reqs)
	if pk := d.db.Create(&reqs); pk.Error != nil {
		log.Fatal(pk.Error)
		return pk.Error
	}
	return nil
}

// GetBill fetches billdata from database
func (d *DataBase) GetBill(reqs *models.FetchBillRequest) (*models.FetchBillResponse, error) {
	db := d.db
	db.AutoMigrate(&models.FetchBillResponse{})

	//*********************** this part should be removed after real database is cerated *******start************//
	// Check if there's any record in the User table
	var resp models.FetchBillResponse
	result := db.First(&resp)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// No records exist, so create a new user
		resp = models.FetchBillResponse{
			RespCode: "1",
			Status:   "SUCCESS",
			Response: models.RespData{
				ChID:             reqs.ChId,
				RefID:            "LZ2EAV1BVVZ0JJQU05MA3RA59IN21301822",
				ApprovalRefNum:   "12345678",
				ResponseCode:     "000",
				ResponseReason:   "Successful",
				ComplianceReason: "",
				ComplianceRespCd: "",
				BillDetails:      reqs.BillDetails.CustomerParams,
				BillerResponse: models.BillerResponse{
					CustomerName: "PRABHA",
					Amount:       "1000",
					DueDate:      "2015-06-20",
					CustConvFee:  "",
					CustConvDesc: "",
					BillDate:     "2015-06-14",
					BillNumber:   "12303",
					BillPeriod:   "june",
					BillTags:     []interface{}{},
				},
				AdditionalInfo: []interface{}{},
			},
		}
		db.Create(&resp)
		resp = models.FetchBillResponse{
			RespCode: "0",
			Status:   "FAILURE",
			Response: models.RespData{
				ChID:             reqs.ChId,
				RefID:            "B8O1THXV28W2JCN8JOFGASGZUNT21301740",
				ApprovalRefNum:   "",
				ResponseCode:     "002",
				ResponseReason:   "Failure",
				ComplianceReason: "BRP042,CPR014,CPR012,AIN004",
				ComplianceRespCd: "",
				//   BillDetails:      []models.BillDetails{},
				BillDetails: []models.CustomerParam{},
				BillerResponse: models.BillerResponse{
					CustomerName: "",
					Amount:       "NaN",
					DueDate:      "",
					CustConvFee:  "",
					CustConvDesc: "",
					BillDate:     "",
					BillNumber:   "",
					BillPeriod:   "",
					BillTags:     []interface{}{},
				},
				AdditionalInfo: []interface{}{},
			},
		}
		db.Create(&resp)
		resp = models.FetchBillResponse{}
		fmt.Println("Inserted new user as there was no data in the User table.")
	} else {
		fmt.Println("Data already exists in the User table.")
	}

	if reqs.IsRealTimeFetch {
		db.Where("RespCode = ?", 0).First(&resp)
	} else {
		db.Where("RespCode = ?", 1).First(&resp)
	}

	return &resp, nil
}

// NewDataBase return instance of gorm.DB database
func NewDataBase() IBillService {
	return &DataBase{
		db: config.InitDB(),
	}
}

// Validate validates the bill coming from request body
func (d *DataBase) Validate(reqs *models.FetchBillRequest) bool {
	var request models.FetchBillRequest

	// Check billerId, agentId, and customer details in the database, Use Gorm to query the database and check if the records exist

	// Retrieve the billerId from the database using some unique identifier.

	if err := d.db.Where("biller_id = ?", reqs.BillDetails.BillerId).First(&request).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// If record not found, then it might be valid as it's a new entry.
			return true
		}
		panic("Database error")
	}

	if request.BillDetails.BillerId != reqs.BillDetails.BillerId {
		return false
	}
	if request.AgentDetails.AgentId != reqs.AgentDetails.AgentId {
		return false
	}
	if request.CustDetails.MobileNo != reqs.CustDetails.MobileNo {
		return false
	}
	if request.AgentDetails.DeviceTags[0].Value != reqs.AgentDetails.DeviceTags[0].Value {
		return false
	}
	if request.CustDetails.CustomerTags[0].Value != reqs.CustDetails.CustomerTags[0].Value {
		return false
	}
	if request.BillDetails.CustomerParams[0].Value != reqs.BillDetails.CustomerParams[0].Value {
		return false
	}

	return true
}
