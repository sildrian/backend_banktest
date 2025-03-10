package controllers

import (
	"bytes"
    "banktest/app/customers/models"
	auth_model "banktest/app/auth/models"
	"banktest/app/customers/controllers"
    "encoding/json"
    "net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestFindCustomer(t *testing.T) {
	tests := []struct {
		name           string
		request        models.CustomerRequest
		expectedStatus int
		expectedResult []auth_model.User
	}{
		{
			name:       "successful find customer",
			expectedStatus: http.StatusOK,
			expectedResult: []auth_model.User{
				{
					ID: 1,
					Name: "tester",
					Username: "tester",
					Password: "tester123",
					Hp: "6238349940",
					Address: "Jl. jalan",
					Bank: auth_model.Bank{
						Name: "Superbank",
						Account: 2949003,
					},
					Pocket: auth_model.Pocket{
						Saldo: 20000000,
					},
					Term: auth_model.Term{
						PrincipalDeposit: 10000000,
						DepositInterestProfit: 125000,
						DepositInterestTax: 5000,
						TotalInvestment: 10120000,
					},
					TotalSaldo: 30120000,
				},
				{
					ID: 2,
					Name: "pentest",
					Username: "pentest",
					Password: "pentest123",
					Hp: "6238344893",
					Address: "Jl. dimana",
					Bank: auth_model.Bank{
						Name: "BNI",
						Account: 1239327290,
					},
					Pocket: auth_model.Pocket{
						Saldo: 10000000,
					},
					Term: auth_model.Term{
						PrincipalDeposit: 5000000,
						DepositInterestProfit: 12500,
						DepositInterestTax: 500,
						TotalInvestment: 5012000,
					},
					TotalSaldo: 15012000,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request body
			reqBody, err := json.Marshal(tt.request)
			assert.NoError(t, err, "Failed to marshal request body")

			// // Create test request
			req := httptest.NewRequest(http.MethodPost, "/get-customer", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			
			// Create response recorder
			w := httptest.NewRecorder()

			// Call the handler
			controllers.GetAllUserData(w, req)

			// Assert response status
			assert.Equal(t, tt.expectedStatus, w.Code, "Status code mismatch")

			// // Parse response body
			var response struct {
				Status  int           `json:"status"`
				Message string        `json:"message"`
				Data    []auth_model.User `json:"data"`
			}
			err = json.NewDecoder(w.Body).Decode(&response)
			assert.NoError(t, err, "Failed to decode response body")

			// // Assert response data
			assert.Equal(t, len(tt.expectedResult), len(response.Data), "Number of users mismatch")
			for i, expected := range tt.expectedResult {
				assert.Equal(t, expected.Name, response.Data[i].Name)
				assert.Equal(t, expected.Username, response.Data[i].Username)
			}
		})
	}
}
