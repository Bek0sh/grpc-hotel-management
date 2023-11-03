package handler

import (
	"github.com/Bek0sh/hotel-management-api-gateway/internal/api_clients"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoomHandler struct {
	cl  *api_clients.RoomClient
	log *logging.Logger
}

func NewRoomHandler(cl *api_clients.RoomClient, log *logging.Logger) *RoomHandler {
	return &RoomHandler{cl: cl, log: log}
}

func (h *RoomHandler) Create(ctx *gin.Context) {
	var input models.Room

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.log.Errorf("failed to bind json in create room, error: %v", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	id, err := h.cl.CreateRoom(&input)
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
			"status": "success",
			"id":     id,
		},
	)
}

func (h *RoomHandler) Update(ctx *gin.Context) {
	var input models.Room

	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.log.Errorf("failed to bind json in update room, error: %v", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	err := h.cl.UpdateRoom(&input)
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

	ctx.JSON(http.StatusOK, "success")
}

func (h *RoomHandler) Delete(ctx *gin.Context) {
	req := ctx.Param("room_number")
	num, _ := strconv.Atoi(req)

	err := h.cl.DeleteRoom(int32(num))
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

func (h *RoomHandler) GetAvailable(ctx *gin.Context) {
	res, err := h.cl.GetAvailableRooms()
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

func (h *RoomHandler) GetByNum(ctx *gin.Context) {
	req := ctx.Param("room_number")
	num, _ := strconv.Atoi(req)

	res, err := h.cl.GetRoomByNumber(int32(num))
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
			"status":    "success",
			"room_type": res,
		},
	)
}
