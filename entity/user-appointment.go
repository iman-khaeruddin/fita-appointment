package entity

import "time"

type UserAppointment struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	CoachID             uint      `gorm:"column:coach_id" json:"coachId"`
	UserID              uint      `gorm:"column:user_id" json:"userId"`
	UserAppointmentDate time.Time `gorm:"column:appointment_date" json:"userAppointmentDate"`
}
