package repository

import (
	"context"
	"fmt"
	"github.com/iman-khaeruddin/fita-appointment/entity"
	"gorm.io/gorm"
	"time"
)

type UserAppointment struct {
	db *gorm.DB
}

func NewUserAppointment(con *gorm.DB) UserAppointment {
	return UserAppointment{
		db: con,
	}
}

type UserAppointmentInterface interface {
	Save(ctx context.Context, userAppointment *entity.UserAppointment) (*entity.UserAppointment, error)
	FindByCoachIDAndAppointmentDate(ctx context.Context, coachID uint, appointmentDate time.Time) error
	UpdateSelectedFields(ctx context.Context, userAppointment *entity.UserAppointment, fields ...string) (*entity.UserAppointment, error)
}

func (repo UserAppointment) Save(ctx context.Context, userAppointment *entity.UserAppointment) (*entity.UserAppointment, error) {
	err := repo.db.WithContext(ctx).Model(&entity.UserAppointment{}).Create(userAppointment).Error
	return userAppointment, err
}

func (repo UserAppointment) FindByCoachIDAndAppointmentDate(ctx context.Context, coachID uint, appointmentDate time.Time) error {
	var userAppointment *entity.UserAppointment
	err := repo.db.WithContext(ctx).Model(&entity.UserAppointment{}).
		Where("coach_id = ?", coachID).
		Where("appointment_date = ?", appointmentDate).
		Find(&userAppointment).
		Error

	if err != nil {
		return err
	}

	if userAppointment.ID == 0 {
		return nil
	} else {
		return fmt.Errorf("no available coach")
	}
}

func (repo UserAppointment) UpdateSelectedFields(ctx context.Context, userAppointment *entity.UserAppointment, fields ...string) (*entity.UserAppointment, error) {
	err := repo.db.WithContext(ctx).Model(userAppointment).Select(fields).Updates(*userAppointment).Error
	return userAppointment, err
}
