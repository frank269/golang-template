package entities

type User struct {
	BaseEntity
	Username string  `json:"username" gorm:"column:username"`
	Password *string `json:"password" gorm:"column:password"`
}

func (t *User) TableName() string {
	return "users"
}
