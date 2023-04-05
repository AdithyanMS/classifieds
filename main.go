package main

import (
	"context"

	"github.com/AdithyanMS/classifieds/controllers"
	"github.com/AdithyanMS/classifieds/middlewares"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

func main() {
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// docker run --name redis-classifieds -p 6379:6379 -d redis
	Pong, err := rds.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Error initialising redis")
	}
	log.Infof("pong: %s", Pong)
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	unsecured := router.Group("/v1")
	{
		unsecured.POST("/login", controllers.Login)
		unsecured.POST("/signup", controllers.Signup)
		secured := unsecured.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Hey)
		}
	}
	return router
}
