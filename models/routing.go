package models

import (
	"github.com/gin-gonic/gin"
)

func Routing() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/api/members", serveMembers)
	router.POST("/my-representatives", myRepresentatives)
	router.Static("/static", "./static")
	router.Run(":3000")
}
