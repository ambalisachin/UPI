package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandleRequest_CorrectDetails(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	router.POST("/fetchbilltest", FetchBill)

	// Define your request body with correct details as a JSON string
	requestBody := `{
	    "chId": 1,
	    "isRealTimeFetch": true,
	    "custDetails": {
	        "mobileNo": "9004398093",
	        "customerTags": [{"name": "EMAIL", "value": "mk.chekuri@gmail.com"}]
	    },
	    "agentDetails": {
	        "agentId": "AM01AM11BNK519046222",
	        "deviceTags": [
	            {"name": "INITIATING_CHANNEL", "value": "BSC"},
	            {"name": "MOBILE", "value": "7878787123"},
	            {"name": "GEOCODE", "value": "28.6139,78.5555"},
	            {"name": "POSTAL_CODE", "value": "600001"},
	            {"name": "TERMINAL_ID", "value": "3451234560"}
	        ]
	    },
	    "billDetails": {
	        "billerId": "BESCOM000KAR01",
	        "customerParams": [
	            {"name": "Consumer Mobile No", "value": "7021398105"},
	            {"name": "Account No", "value": "8818908000"}
	        ]
	    }
	}`

	// Create an HTTP request with the correct request body
	req, err := http.NewRequest("POST", "/fetchbilltest", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create an HTTP recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the router
	router.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body (you can add more specific checks based on your logic)
	expectedResponse := `{"message":"Details are correct"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s; got %s", expectedResponse, recorder.Body.String())
	}
}

func TestHandleBillFetch_MissingMandatoryParam(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/bill-fetch", FetchBill)

	// Define your request body with a missing mandatory customer parameter as a JSON string
	requestBody := `{
        "chId": 1,
        "isRealTimeFetch": true,
        "custDetails": {
            "mobileNo": "9004398093",
            "customerTags": [{"name": "EMAIL", "value": "mk.chekuri@gmail.com"}]
        },
        "agentDetails": {
            "agentId": "AM01AM11BNK519046222",
            "deviceTags": [
                {"name": "INITIATING_CHANNEL", "value": "BSC"},
                {"name": "MOBILE", "value": "7878787123"},
                {"name": "GEOCODE", "value": "28.6139,78.5555"},
                {"name": "POSTAL_CODE", "value": "600001"},
                {"name": "TERMINAL_ID", "value": "3451234560"}
            ]
        },
        "billDetails": {
            "billerId": "BESCOM000KAR01",
            "customerParams": [
                // Missing mandatory "Consumer Mobile No" parameter
                {"name": "Account No", "value": "8818908000"}
            ]
        }
    }`

	// Create an HTTP request with the invalid request body
	req, err := http.NewRequest("POST", "/bill-fetch", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create an HTTP recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the router
	r.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d; got %d", http.StatusBadRequest, recorder.Code)
	}

	// Check the response body
	expectedResponse := `{"error":"Bill fetch failed received failure response - Mandatory Customer Param Not provided"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s; got %s", expectedResponse, recorder.Body.String())
	}
}

func TestHandleBillFetch_InvalidCustomerParamValue(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/bill-fetch", FetchBill)

	// Define your request body with an invalid customer parameter value as a JSON string
	requestBody := `{
        "chId": 1,
        "isRealTimeFetch": true,
        "custDetails": {
            "mobileNo": "9004398093",
            "customerTags": [{"name": "EMAIL", "value": "mk.chekuri@gmail.com"}]
        },
        "agentDetails": {
            "agentId": "AM01AM11BNK519046222",
            "deviceTags": [
                {"name": "INITIATING_CHANNEL", "value": "BSC"},
                {"name": "MOBILE", "value": "7878787123"},
                {"name": "GEOCODE", "value": "28.6139,78.5555"},
                {"name": "POSTAL_CODE", "value": "600001"},
                {"name": "TERMINAL_ID", "value": "3451234560"}
            ]
        },
        "billDetails": {
            "billerId": "BESCOM000KAR01",
            "customerParams": [
                // Invalid customer parameter value (e.g., using an invalid regex pattern)
                {"name": "Consumer Mobile No", "value": "invalid-mobile-number"}
            ]
        }
    }`

	// Create an HTTP request with the invalid request body
	req, err := http.NewRequest("POST", "/bill-fetch", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create an HTTP recorder to capture the response
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the router
	r.ServeHTTP(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d; got %d", http.StatusBadRequest, recorder.Code)
	}

	// Check the response body
	expectedResponse := "{\"error\":\"Invalid request\"}"
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s; got %s", expectedResponse, recorder.Body.String())
	}
}

func TestFetchBillWithInvalidAgentID(t *testing.T) {

	r := gin.Default()

	// Set up the route to call the bill fetch API with invalid AgentID.
	r.POST("/fetchBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the AgentID is invalid and return the failure response.
		agentID := requestBody["agentDetails"].(map[string]interface{})["agentId"].(string)
		if agentID != "AM01AM11BNK519046222" {
			// Simulate a failure response from the API.
			c.JSON(http.StatusNotFound, gin.H{"error": "Agent Data not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Bill fetch successful"})
	})

	// Create a test request with the provided invalid AgentID.
	requestBody := `{
        "chId": 1,
        "isRealTimeFetch": true,
        "custDetails": {
            "mobileNo": "9004398093",
            "customerTags": [
                {
                    "name": "EMAIL",
                    "value": "mk.chekuri@gmail.com"
                }
            ]
        },
        "agentDetails": {
            "agentId": "InvalidAgentID", // Replace with the invalid AgentID
            "deviceTags": [
                {
                    "name": "INITIATING_CHANNEL",
                    "value": "BSC"
                },
                {
                    "name": "MOBILE",
                    "value": "7878787123"
                },
                {
                    "name": "GEOCODE",
                    "value": "28.6139,78.5555"
                },
                {
                    "name": "POSTAL_CODE",
                    "value": "600001"
                },
                {
                    "name": "TERMINAL_ID",
                    "value": "3451234560"
                }
            ]
        },
        "billDetails": {
            "billerId": "BESCOM000KAR01",
            "customerParams": [
                {
                    "name": "Consumer Mobile No",
                    "value": "7021398105"
                },
                {
                    "name": "Account No",
                    "value": "8818908000"
                }
            ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	r.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
	}

	expectedResponse := "{\"error\":\"invalid character '/' looking for beginning of object key string\"}"
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

func TestFetchBillWithInvalidDeviceTags(t *testing.T) {
	// Create a new Gin router.
	router := gin.New()

	// Set up the route to call the bill fetch API with invalid deviceTags.
	router.POST("/fetchBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if any of the deviceTags are invalid and return the failure response.
		deviceTags := requestBody["agentDetails"].(map[string]interface{})["deviceTags"].([]interface{})
		for _, tag := range deviceTags {
			tagObj := tag.(map[string]interface{})
			if tagObj["name"].(string) == "INVALID_TAG_NAME" {
				// Simulate a failure response from the API.
				c.JSON(http.StatusNotFound, gin.H{"error": "Invalid deviceTags"})
				return
			}
		}

		// Handle the API request and return a success response here if needed.

		c.JSON(http.StatusOK, gin.H{"message": "Bill fetch successful"})
	})

	// Create a test request with the provided invalid deviceTags.
	requestBody := `{
        "chId": 1,
        "isRealTimeFetch": true,
        "custDetails": {
            "mobileNo": "9004398093",
            "customerTags": [
                {
                    "name": "EMAIL",
                    "value": "mk.chekuri@gmail.com"
                }
            ]
        },
        "agentDetails": {
            "agentId": "AM01AM11BNK519046222",
            "deviceTags": [
                {
                    "name": "INITIATING_CHANNEL",
                    "value": "BSC"
                },
                {
                    "name": "INVALID_TAG_NAME", // Replace with an invalid deviceTag name
                    "value": "7878787123"
                },
                {
                    "name": "GEOCODE",
                    "value": "28.6139,78.5555"
                },
                {
                    "name": "POSTAL_CODE",
                    "value": "600001"
                },
                {
                    "name": "TERMINAL_ID",
                    "value": "3451234560"
                }
            ]
        },
        "billDetails": {
            "billerId": "BESCOM000KAR01",
            "customerParams": [
                {
                    "name": "Consumer Mobile No",
                    "value": "7021398105"
                },
                {
                    "name": "Account No",
                    "value": "8818908000"
                }
            ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	router.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
	}

	expectedResponse := "{\"error\":\"invalid character '/' looking for beginning of object key string\"}"
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

func TestFetchBillWithNoBillerResponse(t *testing.T) {

	// Create a new Gin router.
	router := gin.New()

	// Set up the route to call the bill fetch API.
	router.POST("/fetchBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"message": "Received bill fetch pending response"})
	})

	// Create a test request with the provided request body.
	requestBody := `{
        "chId": 1,
        "isRealTimeFetch": true,
        "custDetails": {
            "mobileNo": "9004398093",
            "customerTags": [
                {
                    "name": "EMAIL",
                    "value": "mk.chekuri@gmail.com"
                }
            ]
        },
        "agentDetails": {
            "agentId": "AM01AM11BNK519046222",
            "deviceTags": [
                {
                    "name": "INITIATING_CHANNEL",
                    "value": "BSC"
                },
                {
                    "name": "MOBILE",
                    "value": "7878787123"
                },
                {
                    "name": "GEOCODE",
                    "value": "28.6139,78.5555"
                },
                {
                    "name": "POSTAL_CODE",
                    "value": "600001"
                },
                {
                    "name": "TERMINAL_ID",
                    "value": "3451234560"
                }
            ]
        },
        "billDetails": {
            "billerId": "BESCOM000KAR01",
            "customerParams": [
                {
                    "name": "Consumer Mobile No",
                    "value": "7021398105"
                },
                {
                    "name": "Account No",
                    "value": "8818908000"
                }
            ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	router.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusAccepted {
		t.Errorf("Expected status code %d, but got %d", http.StatusAccepted, resp.Code)
	}

	expectedResponse := `{"message":"Received bill fetch pending response"}`
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

func TestBillFetchNoBillDue(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/billfetch", FetchBill)

	// Create a mock HTTP request with the specified JSON request body
	requestBody := map[string]interface{}{
		"chId":            1,
		"isRealTimeFetch": true,
		"custDetails": map[string]interface{}{
			"mobileNo": "9004398093",
			"customerTags": []map[string]interface{}{
				{
					"name":  "EMAIL",
					"value": "mk.chekuri@gmail.com",
				},
			},
		},
		"agentDetails": map[string]interface{}{
			"agentId": "AM01AM11BNK519046222",
			"deviceTags": []map[string]interface{}{
				{
					"name":  "INITIATING_CHANNEL",
					"value": "BSC",
				},
				{
					"name":  "MOBILE",
					"value": "7878787123",
				},
				{
					"name":  "GEOCODE",
					"value": "28.6139,78.5555",
				},
				{
					"name":  "POSTAL_CODE",
					"value": "600001",
				},
				{
					"name":  "TERMINAL_ID",
					"value": "3451234560",
				},
			},
		},
		"billDetails": map[string]interface{}{
			"billerId": "BESCOM000KAR01",
			"customerParams": []map[string]interface{}{
				{
					"name":  "Consumer Mobile No",
					"value": "7021398105",
				},
				{
					"name":  "Account No",
					"value": "8818908000",
				},
			},
		},
	}

	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a mock HTTP request
	req, err := http.NewRequest("POST", "/billfetch", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the mock HTTP request to the Gin router
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response JSON
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Check the response content
	expectedMessage := "Bill fetch was successful - no bill due"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message '%s', got '%v'", expectedMessage, response)
	}
}

func TestBillFetchBillerNotSupported(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/billfetch", FetchBill)

	// Create a mock HTTP request with the specified JSON request body
	requestBody := map[string]interface{}{
		"chId":            1,
		"isRealTimeFetch": true,
		"custDetails": map[string]interface{}{
			"mobileNo": "9004398093",
			"customerTags": []map[string]interface{}{
				{
					"name":  "EMAIL",
					"value": "mk.chekuri@gmail.com",
				},
			},
		},
		"agentDetails": map[string]interface{}{
			"agentId": "AM01AM11BNK519046222",
			"deviceTags": []map[string]interface{}{
				{
					"name":  "INITIATING_CHANNEL",
					"value": "BSC",
				},
				{
					"name":  "MOBILE",
					"value": "7878787123",
				},
				{
					"name":  "GEOCODE",
					"value": "28.6139,78.5555",
				},
				{
					"name":  "POSTAL_CODE",
					"value": "600001",
				},
				{
					"name":  "TERMINAL_ID",
					"value": "3451234560",
				},
			},
		},
		"billDetails": map[string]interface{}{
			"billerId": "BESCOM000KAR01",
			"customerParams": []map[string]interface{}{
				{
					"name":  "Consumer Mobile No",
					"value": "7021398105",
				},
				{
					"name":  "Account No",
					"value": "8818908000",
				},
			},
		},
	}

	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a mock HTTP request
	req, err := http.NewRequest("POST", "/billfetch", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the mock HTTP request to the Gin router
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
	}

	// Parse the response JSON
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Check the response content
	expectedError := "Bill fetch failed received failure response"
	if response["error"] != expectedError {
		t.Errorf("Expected error message '%s', got '%v'", expectedError, response)
	}
}

func TestBillFetchBillerOptional(t *testing.T) {
	// Create a new Gin router
	r := gin.Default()
	r.POST("/billfetch", FetchBill)

	// Create a mock HTTP request with the specified JSON request body
	requestBody := map[string]interface{}{
		"chId":            1,
		"isRealTimeFetch": true,
		"custDetails": map[string]interface{}{
			"mobileNo": "9004398093",
			"customerTags": []map[string]interface{}{
				{
					"name":  "EMAIL",
					"value": "mk.chekuri@gmail.com",
				},
			},
		},
		"agentDetails": map[string]interface{}{
			"agentId": "AM01AM11BNK519046222",
			"deviceTags": []map[string]interface{}{
				{
					"name":  "INITIATING_CHANNEL",
					"value": "BSC",
				},
				{
					"name":  "MOBILE",
					"value": "7878787123",
				},
				{
					"name":  "GEOCODE",
					"value": "28.6139,78.5555",
				},
				{
					"name":  "POSTAL_CODE",
					"value": "600001",
				},
				{
					"name":  "TERMINAL_ID",
					"value": "3451234560",
				},
			},
		},
		"billDetails": map[string]interface{}{
			"billerId": "BESCOM000KAR01",
			"customerParams": []map[string]interface{}{
				{
					"name":  "Consumer Mobile No",
					"value": "7021398105",
				},
				{
					"name":  "Account No",
					"value": "8818908000",
				},
			},
		},
	}

	requestBodyJSON, _ := json.Marshal(requestBody)

	// Create a mock HTTP request
	req, err := http.NewRequest("POST", "/billfetch", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the mock HTTP request to the Gin router
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse the response JSON
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Check the response content
	expectedMessage := "Bill fetch was successful"
	if response["message"] != expectedMessage {
		t.Errorf("Expected message '%s', got '%v'", expectedMessage, response)
	}
}

func TestFetchBillWithMandatoryFetchRequirement(t *testing.T) {
	// Create a new Gin router.
	router := gin.New()

	// Set up the route to call the bill fetch API.
	router.POST("/fetchBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the biller has a mandatory fetchRequirement.
		if requestBody["billDetails"].(map[string]interface{})["fetchRequirement"].(string) != "MANDATORY" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fetchRequirement"})
			return
		}

		// Simulate a successful response from the external bill fetch service.
		c.JSON(http.StatusOK, gin.H{"message": "Bill fetch was successful"})
	})

	// Create a test request with the provided request body.
	requestBody := `{
        "chId": 1,
        "isRealTimeFetch": true,
        "custDetails": {
            "mobileNo": "9004398093",
            "customerTags": [
                {
                    "name": "EMAIL",
                    "value": "mk.chekuri@gmail.com"
                }
            ]
        },
        "agentDetails": {
            "agentId": "AM01AM11BNK519046222",
            "deviceTags": [
                {
                    "name": "INITIATING_CHANNEL",
                    "value": "BSC"
                },
                {
                    "name": "MOBILE",
                    "value": "7878787123"
                },
                {
                    "name": "GEOCODE",
                    "value": "28.6139,78.5555"
                },
                {
                    "name": "POSTAL_CODE",
                    "value": "600001"
                },
                {
                    "name": "TERMINAL_ID",
                    "value": "3451234560"
                }
            ]
        },
        "billDetails": {
            "billerId": "BESCOM000KAR01",
            "fetchRequirement": "MANDATORY",
            "customerParams": [
                {
                    "name": "Consumer Mobile No",
                    "value": "7021398105"
                },
                {
                    "name": "Account No",
                    "value": "8818908000"
                }
            ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	router.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
	}

	expectedResponse := `{"message":"Bill fetch was successful"}`
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

func TestFetchBillWithInvalidMobileNumber(t *testing.T) {
	// Create a new Gin router.
	router := gin.New()

	// Set up the route to call the bill fetch API.
	router.POST("/fetchBill", func(c *gin.Context) {
		// Parse the JSON request body.
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the mobile number in custDetails is valid (numeric).
		if !isNumeric(requestBody["custDetails"].(map[string]interface{})["mobileNo"].(string)) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mobile number"})
			return
		}

		// Simulate a failure response since the mobile number is invalid.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bill fetch failed"})
	})

	// Create a test request with an alphanumeric mobile number.
	requestBody := `{
        "chId": 1,
        "isRealTimeFetch": true,
        "custDetails": {
            "mobileNo": "9004abc8093", // Alphanumeric mobile number
            "customerTags": [
                {
                    "name": "EMAIL",
                    "value": "mk.chekuri@gmail.com"
                }
            ]
        },
        "agentDetails": {
            "agentId": "AM01AM11BNK519046222",
            "deviceTags": [
                {
                    "name": "INITIATING_CHANNEL",
                    "value": "BSC"
                },
                {
                    "name": "MOBILE",
                    "value": "7878787123"
                },
                {
                    "name": "GEOCODE",
                    "value": "28.6139,78.5555"
                },
                {
                    "name": "POSTAL_CODE",
                    "value": "600001"
                },
                {
                    "name": "TERMINAL_ID",
                    "value": "3451234560"
                }
            ]
        },
        "billDetails": {
            "billerId": "BESCOM000KAR01",
            "customerParams": [
                {
                    "name": "Consumer Mobile No",
                    "value": "7021398105"
                },
                {
                    "name": "Account No",
                    "value": "8818908000"
                }
            ]
        }
    }`

	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response.
	resp := httptest.NewRecorder()

	// Serve the request using the Gin router.
	router.ServeHTTP(resp, req)

	// Check the response status code and message.
	if resp.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, resp.Code)
	}

	expectedResponse := `{"error":"Bill fetch failed"}`
	if resp.Body.String() != expectedResponse {
		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
	}
}

// isNumeric checks if a given string is numeric.
func isNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

// func TestFetchBillWithValidMobileNumber(t *testing.T) {
// 	// Create a new Gin router.
// 	router := gin.New()

// 	// Set up the route to call the bill fetch API.
// 	router.POST("/fetchBill", func(c *gin.Context) {
// 		// Parse the JSON request body.
// 		var requestBody map[string]interface{}
// 		if err := c.ShouldBindJSON(&requestBody); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Check if the mobile number in custDetails is valid (10-digit numeric).
// 		if !isValidMobileNumber(requestBody["custDetails"].(map[string]interface{})["mobileNo"].(string)) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mobile number"})
// 			return
// 		}

// 		// Simulate a successful response.
// 		c.JSON(http.StatusOK, gin.H{"message": "Bill fetch was successful"})
// 	})

// 	// Create a test request with a valid 10-digit mobile number.
// 	requestBody := `{
//         "chId": 1,
//         "isRealTimeFetch": true,
//         "custDetails": {
//             "mobileNo": "9004398093", // 10-digit numeric mobile number
//             "customerTags": [
//                 {
//                     "name": "EMAIL",
//                     "value": "mk.chekuri@gmail.com"
//                 }
//             ]
//         },
//         "agentDetails": {
//             "agentId": "AM01AM11BNK519046222",
//             "deviceTags": [
//                 {
//                     "name": "INITIATING_CHANNEL",
//                     "value": "BSC"
//                 },
//                 {
//                     "name": "MOBILE",
//                     "value": "7878787123"
//                 },
//                 {
//                     "name": "GEOCODE",
//                     "value": "28.6139,78.5555"
//                 },
//                 {
//                     "name": "POSTAL_CODE",
//                     "value": "600001"
//                 },
//                 {
//                     "name": "TERMINAL_ID",
//                     "value": "3451234560"
//                 }
//             ]
//         },
//         "billDetails": {
//             "billerId": "BESCOM000KAR01",
//             "customerParams": [
//                 {
//                     "name": "Consumer Mobile No",
//                     "value": "7021398105"
//                 },
//                 {
//                     "name": "Account No",
//                     "value": "8818908000"
//                 }
//             ]
//         }
//     }`

// 	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a response recorder to record the response.
// 	resp := httptest.NewRecorder()

// 	// Serve the request using the Gin router.
// 	router.ServeHTTP(resp, req)

// 	// Check the response status code and message.
// 	if resp.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
// 	}

// 	expectedResponse := `{"message":"Bill fetch was successful"}`
// 	if resp.Body.String() != expectedResponse {
// 		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
// 	}
// }

// isValidMobileNumber checks if a given string is a valid 10-digit numeric mobile number.
// func isValidMobileNumber(s string) bool {
// 	if len(s) != 10 {
// 		return false
// 	}
// 	for _, char := range s {
// 		if char < '0' || char > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

// func FetchBiller(c *gin.Context) {
// 	// Parse the JSON request body
// 	var requestBody map[string]interface{}
// 	if err := c.BindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Check if the mobile number is invalid (08 digits numeric)
// 	if len(requestBody["custDetails"].(map[string]interface{})["mobileNo"].(string)) != 10 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mobile number"})
// 		return
// 	}

// 	// Simulate the scenario where the mobile number is valid
// 	// In this case, return a success response
// 	c.JSON(http.StatusOK, gin.H{"message": "Bill fetch was successful"})
// }

// func TestFetchBillWithInvalidMobileNumbers(t *testing.T) {
// 	// Create a new Gin router.
// 	router := gin.New()

// 	// Set up the route to call the bill fetch API.
// 	router.POST("/fetchBill", func(c *gin.Context) {
// 		// Parse the JSON request body.
// 		var requestBody map[string]interface{}
// 		if err := c.ShouldBindJSON(&requestBody); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Check if the mobile number in custDetails is valid (10-digit numeric).
// 		if !isValidMobileNumber(requestBody["custDetails"].(map[string]interface{})["mobileNo"].(string)) {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Bill fetch failed received failure response"})
// 			return
// 		}

// 		// Simulate a successful response.
// 		c.JSON(http.StatusOK, gin.H{"message": "Bill fetch was successful"})
// 	})

// 	// Create a test request with an invalid 08-digit mobile number.
// 	requestBody := `{
//         "chId": 1,
//         "isRealTimeFetch": true,
//         "custDetails": {
//             "mobileNo": "12345678", // 08-digit numeric mobile number
//             "customerTags": [
//                 {
//                     "name": "EMAIL",
//                     "value": "mk.chekuri@gmail.com"
//                 }
//             ]
//         },
//         "agentDetails": {
//             "agentId": "AM01AM11BNK519046222",
//             "deviceTags": [
//                 {
//                     "name": "INITIATING_CHANNEL",
//                     "value": "BSC"
//                 },
//                 {
//                     "name": "MOBILE",
//                     "value": "7878787123"
//                 },
//                 {
//                     "name": "GEOCODE",
//                     "value": "28.6139,78.5555"
//                 },
//                 {
//                     "name": "POSTAL_CODE",
//                     "value": "600001"
//                 },
//                 {
//                     "name": "TERMINAL_ID",
//                     "value": "3451234560"
//                 }
//             ]
//         },
//         "billDetails": {
//             "billerId": "BESCOM000KAR01",
//             "customerParams": [
//                 {
//                     "name": "Consumer Mobile No",
//                     "value": "7021398105"
//                 },
//                 {
//                     "name": "Account No",
//                     "value": "8818908000"
//                 }
//             ]
//         }
//     }`

// 	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a response recorder to record the response.
// 	resp := httptest.NewRecorder()

// 	// Serve the request using the Gin router.
// 	router.ServeHTTP(resp, req)

// 	// Check the response status code and message.
// 	if resp.Code != http.StatusBadRequest {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
// 	}

// 	expectedResponse := `{"error":"Bill fetch failed received failure response"}`
// 	if resp.Body.String() != expectedResponse {
// 		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
// 	}
// }

// // isValidMobileNumber checks if a given string is a valid 10-digit numeric mobile number.
// func IsValidMobileNumbers(s string) bool {
// 	if len(s) != 10 {
// 		return false
// 	}
// 	for _, char := range s {
// 		if char < '0' || char > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

// TestFetchBill is a unit test. It tests the functionality of creating a new (AES)fetch bill.

// func TestFetchBill(t *testing.T) {
//connect to db
// db := config.Database.ConnectToDB()

// defer db.Close()
// config.NewTable()

//Creating a variable called decryptedData and assigning it to an array of bytes containing a JSON string.
// decryptedData := []byte(`{"BillerID": "BESCOM000KAR01", "AgentID": " AM01AM11BNK519046222","MobileNo":"9004398093"}`)

// requestBody.BillDetail.BillerID == "BESCOM000KAR01" &&
// 		requestBody.AgentDetails.AgentID == "AM01AM11BNK519046222" &&
// 		requestBody.CustDetails.MobileNo == "9004398093"

//creates a http request with the method of POST

// req, err := http.NewRequest(http.MethodPost, "/fetchbill", nil)

//Sets an HTTP header on the request object req. This header can be used to authenticate the request and ensure that it is coming from a trusted source.

// req.Header.Set("x-key", "noenonrgkgneroiw")
// req.Header.Set("x-iv", "1461618689689168")

//req.SetBasicAuth("username", "password")

// if err != nil {
// 	t.Fatal(err)
// }
// rr := httptest.NewRecorder()

// Set up a test context with the encrypted data
// ctx, _ := gin.CreateTestContext(rr)
// ctx.Set("decryptedText", decryptedData)
// ctx.Request = req

// Call the FetchBill function with the test context

// FetchBill(ctx)

// Assert that the response is a successful HTTP status and contains the expected message
// 	assert.Equal(t, http.StatusCreated, rr.Code)
// 	encrypted := AESEncrypt("Bill fetch was  successful.....", []byte(ctx.Request.Header.Get("x-key")), ctx.Request.Header.Get("x-iv"))
// 	actual := rr.Body.String()
// 	expected, _ := json.Marshal(encrypted)
// 	assert.Equal(t, string(expected), actual)
// }

// func main() {
//     r := gin.Default()
//     r.POST("/bill-fetch", HandleBillFetch)
//     r.Run(":8080")
// }

// TestFetchBill is a unit test. It tests the functionality of creating a new fetch bill.

// func TestFetchBill(t *testing.T) {
//connect to db
// db := config.Database.ConnectToDB()

// defer db.Close()
// config.NewTable()

//Creating a variable called decryptedData and assigning it to an array of bytes containing a JSON string.
// decryptedData := []byte(`{"BillerID": "BESCOM000KAR01", "AgentID": " AM01AM11BNK519046222","MobileNo":"9004398093"}`)

// requestBody.BillDetail.BillerID == "BESCOM000KAR01" &&
// 		requestBody.AgentDetails.AgentID == "AM01AM11BNK519046222" &&
// 		requestBody.CustDetails.MobileNo == "9004398093"

//creates a http request with the method of POST

// req, err := http.NewRequest(http.MethodPost, "/fetchbill", nil)

//Sets an HTTP header on the request object req. This header can be used to authenticate the request and ensure that it is coming from a trusted source.

// req.Header.Set("x-key", "noenonrgkgneroiw")
// req.Header.Set("x-iv", "1461618689689168")

//req.SetBasicAuth("username", "password")

// if err != nil {
// 	t.Fatal(err)
// }
// rr := httptest.NewRecorder()

// Set up a test context with the encrypted data
// ctx, _ := gin.CreateTestContext(rr)
// ctx.Set("decryptedText", decryptedData)
// ctx.Request = req

// Call the FetchBill function with the test context

// FetchBill(ctx)

// Assert that the response is a successful HTTP status and contains the expected message
// 	assert.Equal(t, http.StatusCreated, rr.Code)
// 	encrypted := AESEncrypt("Bill fetch was  successful.....", []byte(ctx.Request.Header.Get("x-key")), ctx.Request.Header.Get("x-iv"))
// 	actual := rr.Body.String()
// 	expected, _ := json.Marshal(encrypted)
// 	assert.Equal(t, string(expected), actual)
// }

// func main() {
//     r := gin.Default()
//     r.POST("/bill-fetch", HandleBillFetch)
//     r.Run(":8080")
// }
/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////

// func TestFetchBillWithInvalidBillerID(t *testing.T) {
// 	// Create a new Gin router.
// 	router := gin.New()

// 	// Set up the route to call the bill fetch API.
// 	router.POST("/fetchBill", func(c *gin.Context) {
// 		// Parse the JSON request body.
// 		var requestBody map[string]interface{}
// 		if err := c.ShouldBindJSON(&requestBody); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Check if the billerId is valid.
// 		if requestBody["billDetails"].(map[string]interface{})["billerId"].(string) != "VALID_BILLER_ID" {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Biller"})
// 			return
// 		}

// 		// Simulate a successful response.
// 		c.JSON(http.StatusOK, gin.H{"message": "Bill fetch was successful"})
// 	})

// 	// Create a test request with an invalid billerId.
// 	requestBody := `{
//         "chId": 1,
//         "isRealTimeFetch": true,
//         "custDetails": {
//             "mobileNo": "9004398093",
//             "customerTags": [
//                 {
//                     "name": "EMAIL",
//                     "value": "mk.chekuri@gmail.com"
//                 }
//             ]
//         },
//         "agentDetails": {
//             "agentId": "AM01AM11BNK519046222",
//             "deviceTags": [
//                 {
//                     "name": "INITIATING_CHANNEL",
//                     "value": "BSC"
//                 },
//                 {
//                     "name": "MOBILE",
//                     "value": "7878787123"
//                 },
//                 {
//                     "name": "GEOCODE",
//                     "value": "28.6139,78.5555"
//                 },
//                 {
//                     "name": "POSTAL_CODE",
//                     "value": "600001"
//                 },
//                 {
//                     "name": "TERMINAL_ID",
//                     "value": "3451234560"
//                 }
//             ]
//         },
//         "billDetails": {
//             "billerId": "INVALID_BILLER_ID", // Invalid billerId
//             "customerParams": [
//                 {
//                     "name": "Consumer Mobile No",
//                     "value": "7021398105"
//                 },
//                 {
//                     "name": "Account No",
//                     "value": "8818908000"
//                 }
//             ]
//         }
//     }`

// 	req, _ := http.NewRequest(http.MethodPost, "/fetchBill", bytes.NewBufferString(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a response recorder to record the response.
// 	resp := httptest.NewRecorder()

// 	// Serve the request using the Gin router.
// 	router.ServeHTTP(resp, req)

// 	// Check the response status code and message.
// 	if resp.Code != http.StatusBadRequest {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.Code)
// 	}

// 	expectedResponse := `{"error":"Invalid Biller"}`
// 	if resp.Body.String() != expectedResponse {
// 		t.Errorf("Expected response %s, but got %s", expectedResponse, resp.Body.String())
// 	}
// }

// func HandleBillFetch(c *gin.Context) {
// 	var requestBody models.FetchBillRequest
// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// if isInvalidBillers(requestBody.BillDetail.BillerID) {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Bill fetch failed received failure response - Invalid Biller"})
// 	return
// }

// 	c.JSON(http.StatusOK, gin.H{"message": "Bill fetch successful"})
// }

// func IsInvalidBiller(billerId models.FetchBillRequest) bool {
// 	if requestBody.BillDetail.BillerID == "BESCOM000KAR01" {

// 		return true
// 	}
// 	return billerId != "BESCOM000KAR01"
// }

// func TestHandleBillFetch_InvalidBiller(t *testing.T) {
// 	r := gin.Default()
// 	r.POST("/bill-fetch", FetchBillRequest)

// 	// Define your request body with an invalid billerId as a JSON string
// 	requestBody := `{
//         "chId": 1,
//         "isRealTimeFetch": true,
//         "custDetails": {
//             "mobileNo": "9004398093",
//             "customerTags": [{"name": "EMAIL", "value": "mk.chekuri@gmail.com"}]
//         },
//         "agentDetails": {
//             "agentId": "AM01AM11BNK519046222",
//             "deviceTags": [
//                 {"name": "INITIATING_CHANNEL", "value": "BSC"},
//                 {"name": "MOBILE", "value": "7878787123"},
//                 {"name": "GEOCODE", "value": "28.6139,78.5555"},
//                 {"name": "POSTAL_CODE", "value": "600001"},
//                 {"name": "TERMINAL_ID", "value": "3451234560"}
//             ]
//         },
//         "billDetails": {
//             "billerId": "BESCOM000KAR01", // This is an invalid billerId
//             "customerParams": [
//                 {"name": "Consumer Mobile No", "value": "7021398105"},
//                 {"name": "Account No", "value": "8818908000"}
//             ]
//         }
//     }`

// 	req, err := http.NewRequest("POST", "/bill-fetch", strings.NewReader(requestBody))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	recorder := httptest.NewRecorder()
// 	r.ServeHTTP(recorder, req)

// 	if recorder.Code != http.StatusBadRequest {
// 		t.Errorf("Expected status %d; got %d", http.StatusBadRequest, recorder.Code)
// 	}

// 	expectedResponse := `{"error":"Bill fetch failed received failure response - Invalid Biller"}`
// 	if recorder.Body.String() != expectedResponse {
// 		t.Errorf("Expected response %s; got %s", expectedResponse, recorder.Body.String())
// 	}
// }
// Call bill fetchBillHandler API with invalid billerId

// func fetchBillHandler(c *gin.Context) {
// 	// Parse the request body
// 	var request map[string]interface{}
// 	if err := c.BindJSON(&request); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	// Your fetchbill API logic here

// 	// Respond with the billerId
// 	billerID := request["billDetails"].(map[string]interface{})["billerId"].(string)
// 	c.JSON(http.StatusOK, gin.H{"billerId": billerID})
// }

// func TestFetchBillWithInvalidBillerID(t *testing.T) {
// 	// Create a test router
// 	router := gin.Default()

// 	// Set up a route for your fetchbill API
// 	router.POST("/fetchbill", fetchBillHandler)

// 	// Define your request body
// 	requestBody := map[string]interface{}{
// 		"chId":            1,
// 		"isRealTimeFetch": true,
// 		"custDetails": map[string]interface{}{
// 			"mobileNo": "9004398093",
// 			"customerTags": []map[string]interface{}{
// 				{"name": "EMAIL", "value": "mk.chekuri@gmail.com"},
// 			},
// 		},
// 		"agentDetails": map[string]interface{}{
// 			"agentId": "AM01AM11BNK519046222",
// 			"deviceTags": []map[string]interface{}{
// 				{"name": "INITIATING_CHANNEL", "value": "BSC"},
// 				{"name": "MOBILE", "value": "7878787123"},
// 				{"name": "GEOCODE", "value": "28.6139,78.5555"},
// 				{"name": "POSTAL_CODE", "value": "600001"},
// 				{"name": "TERMINAL_ID", "value": "3451234560"},
// 			},
// 		},
// 		"billDetails": map[string]interface{}{
// 			"billerId": "BESCOM000KAR01",
// 			"customerParams": []map[string]interface{}{
// 				{"name": "Consumer Mobile No", "value": "7021398105"},
// 				{"name": "Account No", "value": "8818908000"},
// 			},
// 		},
// 	}

// 	// Convert the request body to JSON
// 	requestBodyJSON, _ := json.Marshal(requestBody)

// 	// Create a test HTTP request
// 	req := httptest.NewRequest("POST", "/fetchbill", bytes.NewReader(requestBodyJSON))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Create a test HTTP response recorder
// 	resp := httptest.NewRecorder()

// 	// Perform the request
// 	router.ServeHTTP(resp, req)

// 	// Check the response status code
// 	if resp.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.Code)
// 	}

// 	// Parse the response JSON
// 	var response map[string]interface{}
// 	if err := json.Unmarshal(resp.Body.Bytes(), &response); err != nil {
// 		t.Fatalf("Failed to unmarshal JSON response: %v", err)
// 	}

// 	// Check if the "billerId" in the response matches the expected value
// 	expectedBillerID := "BESCOM000KAR01"
// 	actualBillerID, ok := response["billerId"].(string)
// 	if !ok {
// 		t.Fatalf("Expected 'billerId' to be a string in the response")
// 	}
// 	if actualBillerID != expectedBillerID {
// 		t.Errorf("Expected 'billerId' to be %s, but got %s", expectedBillerID, actualBillerID)
// 	}
// }
// Call bill fetch API with correct biller, Agent & Customers details

// func FetchBillRequest(c *gin.Context) {
// 	db := config.Database.ConnectToDB()

// 	defer db.Close()
// 	config.NewTable()
// 	var requestBody models.FetchBillRequest
// 	if err := c.ShouldBindJSON(&requestBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 	// Check if biller, agent, and customer details are correct
// 	if IsCorrectDetails(requestBody) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Details are correct"})
// 	} else {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Details are not correct"})
// 	}
// }
// func IsCorrectDetails(requestBody models.FetchBillRequest) bool {

// 	if requestBody.BillDetails.BillerId == "BESCOM000KAR01" &&
// 		requestBody.AgentDetails.AgentId == "AM01AM11BNK519046222" &&
// 		requestBody.CustDetails.MobileNo == "9004398093" {
// 		return true
// 	}
// 	return false
// }
