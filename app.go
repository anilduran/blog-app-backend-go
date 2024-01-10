package main

import (
	"context"

	"example.com/blog-app-backend-go/db"
	"example.com/blog-app-backend-go/routes"
	"example.com/blog-app-backend-go/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("unable to load SDK config")
	}

	client := s3.NewFromConfig(cfg)

	utils.InitializePresigner(client)

	err = godotenv.Load()

	if err != nil {
		panic("failed to load .env file")
	}

	err = db.InitializeDB()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")

}
