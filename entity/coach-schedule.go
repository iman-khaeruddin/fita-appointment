package entity

type CoachSchedule struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	CoachID        uint   `gorm:"column:coach_id" json:"coachId"`
	DayOfWeek      string `gorm:"column:day_of_week" json:"dayOfWeek"`
	AvailableAt    string `gorm:"column:available_at" json:"availableAt"`
	AvailableUntil string `gorm:"column:available_until" json:"availableUntil"`
}

func (CoachSchedule) TableName() string {
	return "coach_schedule"
}
