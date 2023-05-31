package appointment

type CreateAppointment struct {
	UserID  uint   `json:"userId" binding:"required"`
	CoachID uint   `json:"coachId" binding:"required"`
	Date    string `json:"date"`
}

type CoachDeclineAppointment struct {
	CoachID       uint `json:"coachId"`
	AppointmentID uint `json:"appointmentId"`
}

type CoachRescheduleAppointment struct {
	CoachID       uint   `json:"coachId"`
	AppointmentID uint   `json:"appointmentId"`
	NewDate       string `json:"newDate"`
}

type UserDeclineAppointment struct {
	UserID        uint `json:"coachId"`
	AppointmentID uint `json:"appointmentId"`
}

type ResponseMeta struct {
	Success      bool   `json:"success"`
	MessageTitle string `json:"messageTitle"`
	Message      string `json:"message"`
}
