package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type E struct {
	Events string
}

func start() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/getSystemInfo", getSystemInfo)
	router.GET("/getInterfaceInfo/:name", getInteraceInfo)
	router.POST("/getAllActiveIps", ipScan)
	router.Run("localhost:8080")
}

func getSystemInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetSystemInfo())
}

func getInteraceInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetInterfaceInfo(c.Param("name")))
}

func ipScan(c *gin.Context) {

	body := IpscanPost{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, SendSuccess(GetAllActiveIPs(body.IfaceName, body.IfaceIP, body.Timeout, body.TimeBetweenPackets)))

}
