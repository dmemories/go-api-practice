package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type TempData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tempDataList = []TempData{
	{ID: 1, Name: "A"},
	{ID: 2, Name: "B"},
	{ID: 3, Name: "C"},
	{ID: 4, Name: "D"},
}

func getDataList(c *gin.Context) {
	fmt.Println(tempDataList)
	c.IndentedJSON(http.StatusOK, tempDataList)
}

func createData(c *gin.Context) {
	var newData TempData

	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	tempDataList = append(tempDataList, newData)
	c.IndentedJSON(http.StatusOK, newData)
}

func main() {
	router := gin.Default()

	router.GET("/data", getDataList)
	router.POST("/data", createData)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting Server")

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
