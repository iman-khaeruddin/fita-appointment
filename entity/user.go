package entity

type User struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
}

func (User) TableName() string {
	return "user"
}
