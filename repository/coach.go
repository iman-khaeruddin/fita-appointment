package repository

import (
	"context"
	"github.com/iman-khaeruddin/fita-appointment/entity"
	"gorm.io/gorm"
)

type Coach struct {
	db *gorm.DB
}

func NewCoach(con *gorm.DB) Coach {
	return Coach{
		db: con,
	}
}

type CoachInterface interface {
	FindByID(ctx context.Context, coachID uint) (entity.Coach, error)
}

func (repo Coach) FindByID(ctx context.Context, coachID uint) (entity.Coach, error) {
	var coach entity.Coach
	err := repo.db.WithContext(ctx).
		First(&coach, coachID).
		Error
	return coach, err
}
