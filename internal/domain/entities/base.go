package entities

type BaseEntity struct {
	Id string `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	AuditEntity
}
