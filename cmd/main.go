package main

import (
	"log"
	"net"

	"github.com/duchoang206h/telebot-storage/bot"
	"github.com/duchoang206h/telebot-storage/config"
	"github.com/duchoang206h/telebot-storage/handler"
	"github.com/duchoang206h/telebot-storage/proto"
	"github.com/duchoang206h/telebot-storage/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
)

func main() {
	app := fiber.New()
	bot.InitBot()
	app.Use(logger.New())
	router.SetupRoute(app)
	restPort, grpcPort := config.Config("APP_PORT"), config.Config("GRPC_PORT")
	go func ()  {
		log.Fatal(app.Listen(restPort))
	}()
	//rpc
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	proto.RegisterFileStorageServer(server, &handler.GRPCHandler{})
	log.Printf("gRPC server listening on port %s \n", grpcPort)
	log.Fatal(server.Serve(lis))
}
