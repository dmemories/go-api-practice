package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type TempData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tempDataList = []TempData{}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group(os.Getenv("PROJECT_API_URL"))
	g.POST("/data", h.CreateDataList)
	g.GET("/data", h.GetDataList)
	g.GET("/test", h.Test)
}

func (h *Handler) CreateDataList(c *gin.Context) {
	var newData TempData

	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	tempDataList = append(tempDataList, newData)
	c.JSON(http.StatusOK, newData)
}

func (h *Handler) GetDataList(c *gin.Context) {
	c.JSON(http.StatusOK, tempDataList)
}

func (h *Handler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Test": "Test",
	})
}
