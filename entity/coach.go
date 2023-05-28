package entity

type Coach struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}
