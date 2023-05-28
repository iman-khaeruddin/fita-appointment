package appointment

import "time"

type CreateAppointment struct {
	UserID   int       `json:"userId"`
	CoachID  int       `json:"coachId"`
	Timezone int       `json:"timezone"`
	Date     time.Time `json:"date"`
}

type CoachDeclineAppointment struct {
	CoachID       int `json:"coachId"`
	AppointmentID int `json:"appointmentId"`
}

type CoachRescheduleAppointment struct {
	CoachID       int       `json:"coachId"`
	AppointmentID int       `json:"appointmentId"`
	NewDate       time.Time `json:"newDate"`
}

type UserDeclineAppointment struct {
	UserID        int `json:"coachId"`
	AppointmentID int `json:"appointmentId"`
}
