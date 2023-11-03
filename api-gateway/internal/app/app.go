package app

import (
	"github.com/Bek0sh/hotel-management-api-gateway/internal/api_clients"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/config"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/handler"
	"github.com/Bek0sh/hotel-management-api-gateway/internal/routes"
	"github.com/Bek0sh/hotel-management-api-gateway/pkg/logging"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Start() {
	logging.Init()
	log := logging.GetLogger()

	cfg := config.GetConfig()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	connUserSrv, err := grpc.Dial(cfg.AuthServicePort, opts...)
	if err != nil {
		log.Fatal("failed to connect to auth service")
	}

	connBookingSrv, err := grpc.Dial(cfg.BookingServicePort, opts...)
	if err != nil {
		log.Fatal("failed to connect to booking service")
	}

	authClient := api_clients.NewAuthClient(connUserSrv, log)
	roomTypeClient := api_clients.NewRoomTypeClient(connBookingSrv, log)
	roomClient := api_clients.NewRoomClient(connBookingSrv, log)
	bookingClient := api_clients.NewBookingClient(connBookingSrv, log)

	authHandler := handler.NewAuthHandler(authClient, log)
	roomTypeHandler := handler.NewRoomTypeHandler(roomTypeClient, log)
	roomHandler := handler.NewRoomHandler(roomClient, log)
	bookingHandler := handler.NewBookingHandler(bookingClient, log)

	r := gin.Default()

	routes.Router(r, authHandler, roomHandler, roomTypeHandler, bookingHandler)

	if err = r.Run(cfg.Run.Port); err != nil {
		log.Fatal(err)
	}
}
