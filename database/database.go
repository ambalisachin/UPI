package database

// import (
// 	"BBT/models"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// )

// type Database interface {
// 	AddBill(bill *models.FetchBillRequest) error
// 	GetBill(bill *models.FetchBillRequest) (*models.FetchBillResponse, error)
// 	GetBiilerId() (*models.FetchBillResponse, error)
// }

// // Database Implementations

// type RealDatabase struct {
// 	// ... connection info, etc.
// }

// func (dbs *RealDatabase) GetBiilerId(db *sql.DB) (*models.FetchBillResponse, error) {
// 	defer db.Close()

// 	row, err := db.Query("SELECT * FROM billers")
// 	if err != nil {
// 		return nil, err
// 	}
// 	//Iterates over a collection of rows from a SQL query and stores each row into the "biller" variable which is of type Models.Biller.
// 	//It does this by scanning each row and assigning the values to the ID, Title, and Description fields of the biller variable.
// 	// If an error is encountered, the error is printed to the writer.
// 	var biller models.Biller
// 	for row.Next() {
// 		// if err := row.Scan(&biller.ID, &biller.BillerID, &biller.BillerName, &biller.BillerCategory, &biller.BillerChannel, &biller.BillerSubCategoryName, &biller.BillerCustomerInfo, &biller.BillerMinAmount, &biller.BillerFetchBill, &biller.BillerCategoryKey, &biller.BillerCategoryName, &biller.CoverageCity, &biller.CoverageState, &biller.CoveragePincode, &biller.BillerUpdatedDate, &biller.Status, &biller.IsAvailable); err != nil {
// 		if err := row.Scan(&biller.ID, &biller.BillerID, &biller.BillerName, &biller.BillerAliasName, &biller.BillerCategoryName, &biller.BillerSubCategoryName, &biller.BillerMode, &biller.BillerAcceptsAdhoc, &biller.ParentBiller, &biller.ParentBillerID, &biller.BillerCoverage, &biller.FetchRequirement, &biller.SupportBillValidation, &biller.BillerEffctvFrom, &biller.BillerPymtModes, &biller.IntChngFee, &biller.Status, &biller.AdditonalInfo, &biller.PlanAdditionalInfo); err != nil {
// 			return nil, err
// 		}

// 	}
// 	//Adds a "biller" item to the list of "billers".
// 	//aappend func takes 2 arguments: the list of existing billers and the new biller item that is to be added to the list.
// 	//Func then adds the new biller item to the end of the existing list and returns the new list.
// 	// 	billers = append(billers, biller)
// 	// }
// 	data, _ := json.Marshal(biller)
// 	fmt.Println(string(data))

// 	return nil, nil
// }

// func (dbs *RealDatabase) AddBill(db *sql.DB, reqs *models.FetchBillRequest) error {
// 	// Here would be the real database logic.
// 	defer db.Close()

// 	insertQuery := `INSERT INTO  fetchbills (chId ,isRealTimeFetch ,mobileNo ,email ,agendId ,initiatingchannel , mobile ,geocode ,postalcode ,terminalId,billerId , consumermobileno ,uid ,accountno )VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?);`
// 	_, err := db.Exec(insertQuery,
// 		reqs.ChID, reqs.IsRealTimeFetch, reqs.CustDetails.MobileNo,
// 		reqs.CustDetails.CustomerTags[0].Value,
// 		reqs.AgentDetails.AgentID, reqs.AgentDetails.DeviceTags[0].Value,
// 		reqs.AgentDetails.DeviceTags[1].Value,
// 		reqs.AgentDetails.DeviceTags[2].Value, reqs.AgentDetails.DeviceTags[3].Value,
// 		reqs.AgentDetails.DeviceTags[4].Value, reqs.BillDetail.BillerID,
// 		reqs.BillDetail.CustomerParams[0].Value,
// 		reqs.BillDetail.CustomerParams[1].Value,
// 		reqs.BillDetail.CustomerParams[2].Value)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (dbs *RealDatabase) GetBill(bill *models.FetchBillRequest) (*models.FetchBillResponse, error) {
// 	return nil, nil
// }

// type mockDatabase struct{}

// func (db *mockDatabase) AddBill(bill *models.FetchBillRequest) error {

// 	return nil
// }

// // Handlers

// // func GetUserHandler(db Database) gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		idStr := c.Param("id")
// // 		id, err := strconv.Atoi(idStr)
// // 		if err != nil {
// // 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
// // 			return
// // 		}

// // 		user, err := db.GetUser(id)
// // 		if err != nil {
// // 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
// // 			return
// // 		}

// // 		c.JSON(http.StatusOK, user)
// // 	}
// // }
