package handler

import (
	"github.com/Bek0sh/hotel-management-api-gateway/internal/api_clients"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/models"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	cl  *api_clients.AuthClient
	log *logging.Logger
}

func NewAuthHandler(cl *api_clients.AuthClient, log *logging.Logger) *AuthHandler {
	return &AuthHandler{cl: cl, log: log}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var userInput models.RegisterUser

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		h.log.Errorf("failed to bind json in registration, error: %v", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	res, err := h.cl.RegisterUser(&userInput)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
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
			"id":     res,
		},
	)
}

func (h *AuthHandler) SignIn(ctx *gin.Context) {
	var userInput models.SignIn

	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		h.log.Errorf("failed to bind json in signIn, error: %v", err)

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	res, err := h.cl.SignInUser(&userInput)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "fail",
				"error":  err.Error(),
			},
		)
		return
	}

	ctx.SetCookie("access_token", res, 20, "/", "/", false, true)

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":       "success",
			"access_token": res,
		},
	)
}

func (h *AuthHandler) Profile(ctx *gin.Context) {
	id := ctx.MustGet("user_id")

	res, err := h.cl.GetUserProfile(id.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
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
			"user":   res,
		},
	)
}
