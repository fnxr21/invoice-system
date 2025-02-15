package invoiceindexingdto

type (
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
	PagingInvoiceIndexing struct {
		CurrentPage uint ``
	}
)
