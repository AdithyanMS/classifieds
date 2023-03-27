package main

import (
	"context"
	"fmt"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func hey(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`hey there guys`))
}

func main() {
	r := mux.NewRouter()
	serverPortAddr := "8080"
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
	r.HandleFunc("/hey", hey)
	fmt.Printf("starting server on %s", serverPortAddr)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPortAddr), r)
	fmt.Println("server stopped")

}
