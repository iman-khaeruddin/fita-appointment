package entity

import "time"

type CoachSchedule struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	CoachID        uint      `gorm:"column:coach_id" json:"coachId"`
	Timezone       int       `gorm:"column:timezone" json:"timezone"`
	DayOfWeek      string    `gorm:"column:day_of_week" json:"dayOfWeek"`
	AvailableAt    time.Time `gorm:"column:available_at" json:"availableAt"`
	AvailableUntil time.Time `gorm:"column:available_until" json:"availableUntil"`
}
