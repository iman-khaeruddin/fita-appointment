//go:generate mockery --all --inpackage --case snake

package appointment

import (
	"fmt"
	"github.com/iman-khaeruddin/fita-appointment/constant"
	"github.com/iman-khaeruddin/fita-appointment/entity"
	"github.com/iman-khaeruddin/fita-appointment/repository"
	"golang.org/x/net/context"
	"time"
)

type UseCase struct {
	userAppointmentRepo repository.UserAppointmentInterface
	coachRepo           repository.CoachInterface
	coachScheduleRepo   repository.CoachScheduleInterface
}

type UseCaseInterface interface {
	CreateAppointment(ctx context.Context, request CreateAppointment) ResponseMeta
	CoachDeclineAppointment(ctx context.Context, request CoachDeclineAppointment) ResponseMeta
	CoachRescheduleAppointment(ctx context.Context, request CoachRescheduleAppointment) ResponseMeta
	UserDeclineAppointment(ctx context.Context, request UserDeclineAppointment) ResponseMeta
}

func (u UseCase) CreateAppointment(ctx context.Context, request CreateAppointment) ResponseMeta {
	date, err := time.Parse("2006-01-02T15:04:05Z0700", request.Date)
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      "invalid date format",
		}
	}

	coach, err := u.coachRepo.FindByID(ctx, request.CoachID)
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		}
	}

	dateUTC := date.In(time.UTC)
	coachDate := dateUTC.Add(time.Duration(coach.Timezone) * time.Hour)
	fmt.Println("payload date :", request.Date)
	fmt.Println("date UTC :", dateUTC)
	fmt.Println("coach datetime :", coachDate)

	err = u.coachScheduleRepo.FindAvailableCoach(ctx, request.CoachID, coachDate)
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		}
	}

	err = u.userAppointmentRepo.FindByCoachIDAndAppointmentDate(ctx, request.CoachID, dateUTC)
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		}
	}

	userAppointment := &entity.UserAppointment{
		CoachID:             request.CoachID,
		UserID:              request.UserID,
		UserAppointmentDate: dateUTC,
		Status:              constant.WAITING,
	}
	userAppointment, err = u.userAppointmentRepo.Save(ctx, userAppointment)
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		}
	}

	return ResponseMeta{
		Success:      true,
		MessageTitle: "success",
		Message:      "success",
	}
}

func (u UseCase) CoachDeclineAppointment(ctx context.Context, request CoachDeclineAppointment) ResponseMeta {
	var ua *entity.UserAppointment
	ua = &entity.UserAppointment{
		ID:      request.AppointmentID,
		CoachID: request.CoachID,
		Status:  constant.DECLINE,
	}
	u.userAppointmentRepo.UpdateSelectedFields(ctx, ua, "Status", "CoachID")

	return ResponseMeta{
		Success:      true,
		MessageTitle: "success",
		Message:      "success",
	}
}

func (u UseCase) CoachRescheduleAppointment(ctx context.Context, request CoachRescheduleAppointment) ResponseMeta {
	date, err := time.Parse("2006-01-02T15:04:05Z0700", request.NewDate)
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      "invalid date format",
		}
	}

	dateUTC := date.In(time.UTC)

	var ua *entity.UserAppointment
	ua = &entity.UserAppointment{
		ID:                  request.AppointmentID,
		Status:              constant.WAITING,
		UserAppointmentDate: dateUTC,
	}
	_, err = u.userAppointmentRepo.UpdateSelectedFields(ctx, ua, "Status", "UserAppointmentDate")
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		}
	}

	return ResponseMeta{
		Success:      true,
		MessageTitle: "success",
		Message:      "success",
	}
}

func (u UseCase) UserDeclineAppointment(ctx context.Context, request UserDeclineAppointment) ResponseMeta {
	var ua *entity.UserAppointment
	ua = &entity.UserAppointment{
		ID:     request.AppointmentID,
		UserID: request.UserID,
		Status: constant.DECLINE,
	}
	_, err := u.userAppointmentRepo.UpdateSelectedFields(ctx, ua, "Status", "UserID")
	if err != nil {
		return ResponseMeta{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		}
	}

	return ResponseMeta{
		Success:      true,
		MessageTitle: "success",
		Message:      "success",
	}
}
