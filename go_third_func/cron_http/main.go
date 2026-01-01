package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var cr *CrontabManager

func main() {
	cr = NewCronManager()
	cr.Serve()

	r := gin.Default()
	r.POST("/cron/add", AddCron)
	r.POST("/cron/delete", DelCron)
	r.Run(":8080")
}

func AddCron(c *gin.Context) {
	var req struct {
		Spec   string                 `json:"spec" binding:"required"`
		Name   string                 `json:"name" binding:"required"`
		Action string                 `json:"action" binding:"required"`
		Params map[string]interface{} `json:"params"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := cr.AddCronJob(req.Spec, req.Name, req.Action, req.Params); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func DelCron(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cr.DelCronJob(req.Name); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
