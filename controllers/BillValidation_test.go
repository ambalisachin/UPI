package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestBillValidationWithCorrectDetails(t *testing.T) {
	// Create a new Gin router.
	router := gin.New()

	// Set up the route to call the bill validation API.
	router.POST("/validateBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the agentId and billerId are correct.
		if requestBody["agentDetails"].(map[string]interface{})["agentId"].(string) != "AM01YB41BSC519046456" ||
			requestBody["billDetails"].(map[string]interface{})["billerId"].(string) != "OUSH40000NAT02" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bill Validation failed"})
			return
		}

		// Simulate a successful response.
		c.JSON(http.StatusOK, gin.H{"message": "Bill Validation successful"})
	})

	// Create a test request with correct details.
	requestBody := `{
        "agentDetails": {
            "agentId": "AM01YB41BSC519046456"
        },
        "billDetails": {
            "billerId": "OUSH40000NAT02",
            "customerParams": [
                {
                    "name": "Registered Mobile Number / Viewing Card Number",
                    "value": "9987654321"
                }
            ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/validateBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	router.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}

	expectedResponse := `{"message":"Bill Validation successful"}`
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

func TestBillValidationWithInvalidBillerID(t *testing.T) {
	// Create a new Gin router.
	router := gin.New()

	// Set up the route to call the bill validation API.
	router.POST("/validateBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the agentId and billerId are correct.
		if requestBody["agentDetails"].(map[string]interface{})["agentId"].(string) != "AM01YB41BSC519046456" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bill Validation failed"})
			return
		}

		if requestBody["billDetails"].(map[string]interface{})["billerId"].(string) != "INVALID_BILLER_ID" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Biller"})
			return
		}

		// Simulate a failure response with an error message.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bill Validation failed - Invalid Biller"})
	})

	// Create a test request with an invalid billerId.
	requestBody := `{
        "agentDetails": {
            "agentId": "AM01YB41BSC519046456"
        },
        "billDetails": {
            "billerId": "INVALID_BILLER_ID", // Invalid billerId
            "customerParams": [
                {
                    "name": "Registered Mobile Number / Viewing Card Number",
                    "value": "9987654321"
                }
            ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/validateBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	router.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
	}

	expectedResponse := `{"error":"Bill Validation failed - Invalid Biller"}`
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

func TestBillValidationWithMissingMandatoryCustomerParam(t *testing.T) {
	// Create a new Gin router.
	router := gin.New()

	// Set up the route to call the bill validation API.
	router.POST("/validateBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the agentId and billerId are correct.
		if requestBody["agentDetails"].(map[string]interface{})["agentId"].(string) != "AM01YB41BSC519046456" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bill Validation failed"})
			return
		}

		// Check for mandatory customer parameter "Registered Mobile Number / Viewing Card Number".
		customerParams := requestBody["billDetails"].(map[string]interface{})["customerParams"].([]interface{})
		paramFound := false
		for _, param := range customerParams {
			paramName := param.(map[string]interface{})["name"].(string)
			paramValue := param.(map[string]interface{})["value"].(string)
			if paramName == "Registered Mobile Number / Viewing Card Number" && paramValue == "9987654321" {
				paramFound = true
				break
			}
		}

		if !paramFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Mandatory Customer Param Not provided"})
			return
		}

		// Simulate a success response.
		c.JSON(http.StatusOK, gin.H{"message": "Bill Validation successful"})
	})

	// Create a test request with missing mandatory customer parameter.
	requestBody := `{
        "agentDetails": {
            "agentId": "AM01YB41BSC519046456"
        },
        "billDetails": {
            "billerId": "OUSH40000NAT02",
            "customerParams": [ ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/validateBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	router.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
	}

	expectedResponse := `{"error":"Mandatory Customer Param Not provided"}`
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

// func TestBillValidationWithInvalidRegexForCustomerParam(t *testing.T) {
// 	// Create a new Gin router.
// 	router := gin.New()

// 	// Set up the route to call the bill validation API.
// 	router.POST("/validateBill", func(c *gin.Context) {
// 		// Parse the JSON request body.
// 		var requestBody map[string]interface{}
// 		if err := c.ShouldBindJSON(&requestBody); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Check if the agentId and billerId are correct.
// 		if requestBody["agentDetails"].(map[string]interface{})["agentId"].(string) != "AM01YB41BSC519046456" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Bill Validation failed"})
// 			return
// 		}

// 		// Check for the customer parameter "Registered Mobile Number / Viewing Card Number".
// 		customerParams := requestBody["billDetails"].(map[string]interface{})["customerParams"].([]interface{})
// 		for _, param := range customerParams {
// 			paramName := param.(map[string]interface{})["name"].(string)
// 			paramValue := param.(map[string]interface{})["value"].(string)
// 			if paramName == "Registered Mobile Number / Viewing Card Number" {
// 				// Check if the value matches the expected regex pattern.
// 				if !isValidMobileNumber(paramValue) {
// 					c.JSON(http.StatusBadRequest, gin.H{"error": "Bill Validation failed - Invalid Customer Param Value"})
// 					return
// 				}
// 			}
// 		}

// 		// Simulate a success response.
// 		c.JSON(http.StatusOK, gin.H{"message": "Bill Validation successful"})
// 	})

// 	// Create a test request with an invalid customer parameter value (alpha-numeric).
// 	requestBody := `{
//         "agentDetails": {
//             "agentId": "AM01YB41BSC519046456"
//         },
//         "billDetails": {
//             "billerId": "OUSH40000NAT02",
//             "customerParams": [
//                 {
//                     "name": "Registered Mobile Number / Viewing Card Number",
//                     "value": "invalid123" // Invalid value with alpha-numeric characters
//                 }
//             ]
//         }
//     }`

// 	req, _ := http.NewRequest(http.MethodPost, "/validateBill", bytes.NewBufferString(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a response recorder to record the response.
// 	resp := httptest.NewRecorder()

// 	// Serve the request using the Gin router.
// 	router.ServeHTTP(resp, req)

// 	// Check the response status code and message.
// 	if resp.Code != http.StatusBadRequest {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
// 	}

// 	expectedResponse := `{"error":"Bill Validation failed - Invalid Customer Param Value"}`
// 	if resp.Body.String() != expectedResponse {
// 		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
// 	}
// }

// // Function to validate the mobile number using a regex pattern.
// func IsValidMobNumber(value string) bool {
// 	// Define the regex pattern for a valid mobile number.
// 	pattern := `^[0-9]{10}$`
// 	match, _ := regexp.MatchString(pattern, value)
// 	return match
// }
