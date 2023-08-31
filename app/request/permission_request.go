package request

type PermissionRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
}
