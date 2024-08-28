package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"example/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/gin-contrib/cors"
)

var (
	r = gin.Default()
)

func StartApp() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config := cors.DefaultConfig()

	// config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	
	r.Use(cors.New(config))
	// r.Use(AuthMiddleware())
	
	model.DBConnect()
	Router()
	port := os.Getenv("PORT")
	// r.SetTrustedProxies([]string{"http://localhost:3000"})
	r.Run(port)

	fmt.Println("Port number: ", port)

	// Signal handling for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
			<-quit
			fmt.Println("Shutting down...")
			model.DB.Close()
			os.Exit(0)
	}()
	
}
