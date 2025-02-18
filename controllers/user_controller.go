package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-app/services"
	"strconv"
)

func GetUser(c *gin.Context) {
	const op = "GetUser"

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUser(c, uint(id))
	if err != nil {
		log.Printf("%v: services.GetUser error: %v", op, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
