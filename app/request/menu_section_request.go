package request

type MenuSectionRequest struct {
	Name      string `json:"name" validate:"required"`
	Ord       int    `json:"ord" validate:"required"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
