package appointment

import (
	"github.com/gin-gonic/gin"
	"github.com/iman-khaeruddin/fita-appointment/repository"
	"gorm.io/gorm"
	"net/http"
)

type RequestHandler struct {
	db   *gorm.DB
	ctrl ControllerInterface
}

func NewRequestHandler(
	db *gorm.DB,
) RequestHandler {
	return RequestHandler{
		db: db,
	}
}

func (h RequestHandler) Handle(r *gin.Engine) {
	userAppointmentRepo := repository.NewUserAppointment(h.db)
	coachRepo := repository.NewCoach(h.db)
	coachScheduleRepo := repository.NewCoachSchedule(h.db)
	useCase := UseCase{
		userAppointmentRepo: userAppointmentRepo,
		coachRepo:           coachRepo,
		coachScheduleRepo:   coachScheduleRepo,
	}
	h.ctrl = Controller{
		UseCase: useCase,
	}

	r.POST("/create-appointment", h.createAppointment)
	r.POST("/coach-decline-appointment", h.coachDeclineAppointment)
	r.POST("/coach-reschedule-appointment", h.coachRescheduleAppointment)
	r.POST("/user-decline-appointment", h.userDeclineAppointment)
}

func (h RequestHandler) createAppointment(c *gin.Context) {
	var request CreateAppointment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res := h.ctrl.CreateAppointment(c.Request.Context(), request)
	c.JSON(http.StatusOK, res)
	return
}

func (h RequestHandler) coachDeclineAppointment(c *gin.Context) {
	var request CoachDeclineAppointment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res := h.ctrl.CoachDeclineAppointment(c.Request.Context(), request)
	c.JSON(http.StatusOK, res)
	return
}

func (h RequestHandler) coachRescheduleAppointment(c *gin.Context) {
	var request CoachRescheduleAppointment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res := h.ctrl.CoachRescheduleAppointment(c.Request.Context(), request)
	c.JSON(http.StatusOK, res)
	return
}

func (h RequestHandler) userDeclineAppointment(c *gin.Context) {
	var request UserDeclineAppointment
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	res := h.ctrl.UserDeclineAppointment(c.Request.Context(), request)
	c.JSON(http.StatusOK, res)
	return
}
