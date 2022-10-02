package controllers

// CustomersResponse returns just the data we want to return.
type CustomersResponse struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
}
