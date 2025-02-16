package invoicedto

type (
	Item struct {
		ID        uint    `json:"id"`
		Name      string  `json:"items_name"`
		Quantity  float64 `json:"quantity"`
		UnitPrice float64 `json:"unit_price"`
	}
	InvoiceRequest struct {
		IssueDate  string `json:"issue_date"`
		DueDate    string `json:"due_date"`
		Subject    string `json:"subject"`
		CustomerID uint   `json:"customer_id"`
		Items      []Item `json:"items"`
	}

	Customer struct {
		ID      uint   `json:"id"`
		Name    string `json:"customer_name"`
		Address string `json:"customer_address"`
	}

	InvoiceResponse struct {
		InvoiceID uint     `json:"invoice_id"`
		IssueDate string   `json:"issue_date"`
		DueDate   string   `json:"Due_date"`
		Subject   string   `json:"subject"`
		Customer  Customer `json:"customer"`
		Items     []Item   `json:"items"`
	}
	InvoiceIndexing struct {
		InvoiceID    uint   `json:"invoice_id"`
		IssueDate    string `json:"issue_date"`
		DueDate      string `json:"due_date"`
		Subject      string `json:"subject"`
		CustomerName string `json:"customer_name"`
		TotalItems   int    `json:"total_items"`
		Status       string `json:"status"`
		Page         int    `json:"page"`
		Size         int    `json:"size" validate:""`
	}
	InvoiceIndexingResponse struct {
		InvoiceID    uint   `json:"invoice_id"`
		IssueDate    string `json:"issue_date"`
		DueDate      string `json:"due_date"`
		Subject      string `json:"subject"`
		CustomerName string `json:"customer_name"`
		TotalItems   int    `json:"total_items"`
		Status       string `json:"status"`
	}
	PagingInvoiceIndexing struct {
		CurrentPage int `json:"current_page"`
		// TotalPage   int `json:"total_page"`
		Size        int `json:"size"`
	}
)
