//go:generate mockery --all
//go:generate mockery --all --inpackage --case snake

package repository

import (
	"context"
	"fmt"
	"github.com/iman-khaeruddin/fita-appointment/entity"
	"gorm.io/gorm"
	"time"
)

type CoachSchedule struct {
	db *gorm.DB
}

func NewCoachSchedule(con *gorm.DB) CoachSchedule {
	return CoachSchedule{
		db: con,
	}
}

type CoachScheduleInterface interface {
	FindAvailableCoach(ctx context.Context, coachID uint, time time.Time) error
}

func (repo CoachSchedule) FindAvailableCoach(ctx context.Context, coachID uint, time time.Time) error {
	var coachSchedule entity.CoachSchedule
	err := repo.db.WithContext(ctx).
		Model(&entity.CoachSchedule{}).
		Where("coach_id = ?", coachID).
		Where("day_of_week = ?", time.Weekday().String()).
		Where("available_at <= ?", time.Format("15:04:05")).
		Where("available_until > ?", time.Format("15:04:05")).
		Find(&coachSchedule).
		Error

	if err != nil {
		return err
	}

	if coachSchedule.ID == 0 {
		return fmt.Errorf("no available coach")
	}
	return nil
}
