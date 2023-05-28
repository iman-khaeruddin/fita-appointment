package appointment

import "github.com/iman-khaeruddin/fita-appointment/repository"

type UseCase struct {
	repo repository.UserAppointmentInterface
}

type UseCaseInterface interface {
}
