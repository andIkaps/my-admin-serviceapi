package request

type MenuRequest struct {
	Name          string `json:"name" validate:"required"`
	Url           string `json:"url" validate:"required"`
	Icon          string `json:"icon"`
	Ord           int    `json:"ord" validate:"required"`
	ParentID      string `json:"parent_id"`
	MenuSectionID string `json:"menu_section_id"`
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
}
