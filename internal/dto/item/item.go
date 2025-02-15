package item

type (
	ItemRequest struct {
		Name string `json:"item_name" `
		Type string `json:"item_type" `
	}

	ItemResponse struct {
		Name string `json:"item_name"`
		Type string `json:"item_type"`
	}
)
