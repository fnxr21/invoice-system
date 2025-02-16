package resultdto

type (
	SuccessResult struct {
		// Code int         `json:"code"`
		Data interface{} `json:"data"`
	}
	SuccessResultIndex struct {
		// Code int         `json:"code"`
		Data interface{} `json:"data"`
		Pagination interface{} `json:"paging"`
	}
	ErrorResult struct {
		// Status  int         `json:"code"`
		Errors interface{} `json:"errors"`
	}
)
