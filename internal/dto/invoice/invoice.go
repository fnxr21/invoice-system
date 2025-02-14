package invoicedto

type (
	Item struct {
		ID        uint    `json:"id"`
		Name      string  `json:"items_name"`
		Quantity  float64 `json:"quantity"`
		UnitPrice float64 `json:"unit_price"`
	}

	Customer struct {
		ID      uint   `json:"id"`
		Name    string `json:"customer_name"`
		Address string `json:"customer_address"`
	}

	InvoiceRequest struct {
		IssueDate  string `json:"issue_date"`
		DueDate    string `json:"Due_date"`
		Subject    string `json:"subject"`
		CustomerID uint   `json:"customer_id"`
		Items      []Item `json:"items"`
	}

	InvoiceResponse struct {
		InvoiceID uint     `json:"invoice_id"`
		IssueDate string   `json:"issue_date"`
		DueDate   string   `json:"Due_date"`
		Subject   string   `json:"subject"`
		Customer  Customer `json:"customer"`
		Items     []Item   `json:"items"`
	}
)
