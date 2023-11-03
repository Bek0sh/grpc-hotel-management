package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *AuthHandler) CheckUser(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("access_token")
		if err != nil {
			h.log.Errorf("failed to find access_token, error: %v", err)
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  "fail",
					"message": "failed to find access_token",
				},
			)
			return
		}

		if token == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  "fail",
					"message": "you are not logged in",
				},
			)
			return
		}

		id, role, err := h.cl.ValidateAccessToken(token)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusNotAcceptable,
				gin.H{
					"status":  "fail",
					"message": err.Error(),
				},
			)
			return
		}
		ctx.Set("user_id", id)
		ctx.Set("user_role", role)

		next(ctx)
	}
}

func (h *AuthHandler) CheckAdmin(next gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role := ctx.MustGet("user_role")

		if strings.ToLower(role.(string)) != "admin" {
			ctx.AbortWithStatusJSON(
				http.StatusNotAcceptable,
				gin.H{
					"status":  "fail",
					"message": "this action is only for admins",
				},
			)
			return
		}

		next(ctx)
	}
}
