package resultdto

type (
	SuccessResult struct {
		// Code int         `json:"code"`
		Data interface{} `json:"data"`
	}
	ErrorResult struct {
		// Status  int         `json:"code"`
		Errors interface{} `json:"errors"`
	}
)
