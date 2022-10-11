package main

import (
	"fmt"
	"net/http"

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

	router.Run("localhost:8080")
}
