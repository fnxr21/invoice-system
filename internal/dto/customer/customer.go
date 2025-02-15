package customerdto

type (
	CustomerRequest struct {
		Name    string `json:"customer_name" validate:"required"`
		Address string `json:"customer_address" validate:"required"`
	}

	CustomerResponse struct {
		Name    string `json:"customer_name"`
		Address string `json:"customer_address"`
	}
)
