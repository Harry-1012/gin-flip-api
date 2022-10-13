package response

import "github.com/haoleiqin/gin-flip-api/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
