package handler

import (
	"github.com/Bek0sh/hotel-management-api-gateway/internal/api_clients"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookingHandler struct {
	cl  *api_clients.BookingClient
	log *logging.Logger
}

func NewBookingHandler(cl *api_clients.BookingClient, log *logging.Logger) *BookingHandler {
	return &BookingHandler{cl: cl, log: log}
}

func (h *BookingHandler) Create(ctx *gin.Context) {
	var input models.Booking

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.log.Errorf("failed to bind json in create booking, error: %v", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}
	userId := ctx.MustGet("user_id")
	input.CustomerId = userId.(string)
	id, price, err := h.cl.CreateBooking(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":      "success",
			"id":          id,
			"total_price": price,
		},
	)
}
func (h *BookingHandler) Update(ctx *gin.Context) {
	var input models.Booking

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.log.Errorf("failed to bind json in update booking, error: %v", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	msg, err := h.cl.UpdateBooking(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": msg,
	})
}
func (h *BookingHandler) Delete(ctx *gin.Context) {
	req := ctx.Query("id")

	err := h.cl.DeleteBooking(req)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}
	ctx.JSON(
		http.StatusOK,
		"success",
	)
}
func (h *BookingHandler) GetAll(ctx *gin.Context) {
	res, err := h.cl.GetAllBookings()
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":     "success",
			"room_types": res,
		},
	)
}
func (h *BookingHandler) GetCustomers(ctx *gin.Context) {
	id := ctx.MustGet("user_id")
	res, err := h.cl.GetCustomersBookings(id.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":     "success",
			"room_types": res,
		},
	)
}
