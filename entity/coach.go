package entity

type Coach struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Timezone int64  `gorm:"column:timezone" json:"timezone"`
}

func (Coach) TableName() string {
	return "coach"
}
