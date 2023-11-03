package routes

import (
	"github.com/Bek0sh/hotel-management-api-gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine, auth *handler.AuthHandler,
	room *handler.RoomHandler,
	roomType *handler.RoomTypeHandler,
	booking *handler.BookingHandler) {

	router := r.Group("hotel/v1/")

	router.POST("/auth/register", auth.Register)
	router.POST("/auth/sign-in", auth.SignIn)
	router.GET("/profile", auth.CheckUser(auth.Profile))

	router.POST("/room-type/create", auth.CheckUser(auth.CheckAdmin(roomType.CreateRoomType)))
	router.GET("/room-type/all", auth.CheckUser(roomType.GetAll))
	router.GET("/room-type", auth.CheckUser(roomType.GetByType))
	router.PATCH("/room-type/update", auth.CheckUser(auth.CheckAdmin(roomType.Update)))
	router.DELETE("/room-type/delete", auth.CheckUser(auth.CheckAdmin(roomType.Delete)))

	router.POST("/room/create", auth.CheckUser(auth.CheckAdmin(room.Create)))
	router.GET("/room/all", auth.CheckUser(room.GetAvailable))
	router.GET("/room", auth.CheckUser(room.GetByNum))
	router.PATCH("/room/update", auth.CheckUser(auth.CheckAdmin(room.Update)))
	router.DELETE("/room/delete", auth.CheckUser(auth.CheckAdmin(room.Delete)))

	router.POST("/booking/create", auth.CheckUser(booking.Create))
	router.GET("/booking/all", auth.CheckUser(auth.CheckAdmin(booking.GetAll)))
	router.GET("/booking", auth.CheckUser(booking.GetCustomers))
	router.PATCH("/booking/update", auth.CheckUser(booking.Update))
	router.DELETE("/booking/delete", auth.CheckUser(auth.CheckAdmin(booking.Delete)))
}
