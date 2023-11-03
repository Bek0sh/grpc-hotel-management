package handler

import (
	"github.com/Bek0sh/hotel-management-api-gateway/internal/api_clients"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoomTypeHandler struct {
	cl  *api_clients.RoomTypeClient
	log *logging.Logger
}

func NewRoomTypeHandler(cl *api_clients.RoomTypeClient, log *logging.Logger) *RoomTypeHandler {
	return &RoomTypeHandler{cl: cl, log: log}
}

func (h *RoomTypeHandler) CreateRoomType(ctx *gin.Context) {
	var input models.RoomType

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

	id, err := h.cl.CreateRoomType(&input)
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

func (h *RoomTypeHandler) GetByType(ctx *gin.Context) {
	req := ctx.Query("type")

	res, err := h.cl.GetRoomTypeByType(req)
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

func (h *RoomTypeHandler) Update(ctx *gin.Context) {
	var input models.RoomType

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

	err := h.cl.UpdateRoomType(&input)
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

func (h *RoomTypeHandler) GetAll(ctx *gin.Context) {
	res, err := h.cl.GetAllRoomTypes()
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

func (h *RoomTypeHandler) Delete(ctx *gin.Context) {
	req := ctx.Query("type")

	err := h.cl.DeleteRoomType(req)
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
