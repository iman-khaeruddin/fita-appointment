package appointment

import "context"

type Controller struct {
	UseCase UseCaseInterface
}

type ControllerInterface interface {
	CreateAppointment(ctx context.Context, request CreateAppointment) ResponseMeta
	CoachDeclineAppointment(ctx context.Context, request CoachDeclineAppointment) ResponseMeta
	CoachRescheduleAppointment(ctx context.Context, request CoachRescheduleAppointment) ResponseMeta
	UserDeclineAppointment(ctx context.Context, request UserDeclineAppointment) ResponseMeta
}

func (ctrl Controller) CreateAppointment(ctx context.Context, request CreateAppointment) ResponseMeta {
	return ctrl.UseCase.CreateAppointment(ctx, request)
}

func (ctrl Controller) CoachDeclineAppointment(ctx context.Context, request CoachDeclineAppointment) ResponseMeta {
	return ctrl.UseCase.CoachDeclineAppointment(ctx, request)
}

func (ctrl Controller) CoachRescheduleAppointment(ctx context.Context, request CoachRescheduleAppointment) ResponseMeta {
	return ctrl.UseCase.CoachRescheduleAppointment(ctx, request)
}

func (ctrl Controller) UserDeclineAppointment(ctx context.Context, request UserDeclineAppointment) ResponseMeta {
	return ctrl.UseCase.UserDeclineAppointment(ctx, request)
}
