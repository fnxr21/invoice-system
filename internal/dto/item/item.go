package item

type (
	ItemRequest struct {
		Name string `json:"item_name" validate:"required"`
		Type string `json:"item_type" validate:"required"`
	}

	ItemResponse struct {
		Name string `json:"item_name"`
		Type string `json:"item_type"`
	}
)
