package repository

import "gorm.io/gorm"

type UserSchedule struct {
	db *gorm.DB
}

func NewUserSchedule(con *gorm.DB) UserSchedule {
	return UserSchedule{
		db: con,
	}
}

type UserScheduleInterface interface {
}
