package entities

type AuditEntity struct {
	CreatedAt int64   `json:"createdAt" gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt int64   `json:"updatedAt" gorm:"autoUpdateTime:milli;column:updated_at"`
	CreatedBy *string `json:"createdBy" gorm:"column:created_by"`
	UpdatedBy *string `json:"updatedBy" gorm:"column:updated_by"`
}
