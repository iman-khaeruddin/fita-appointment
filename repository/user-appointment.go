package repository

import "gorm.io/gorm"

type UserAppointment struct {
	db *gorm.DB
}

func NewUserAppointment(con *gorm.DB) UserAppointment {
	return UserAppointment{
		db: con,
	}
}

type UserAppointmentInterface interface {
}
