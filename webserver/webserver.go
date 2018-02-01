package webserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mgerb/tmail/db"
	"github.com/mgerb/tmail/mail"
)

func Start() {
	r := gin.Default()
	r.GET("/api/mail", mailHander)
	r.Run("0.0.0.0:8090")
}

func mailHander(c *gin.Context) {

	to := c.Query("to")
	var mail []mail.Mail
	log.Println(to)

	if to != "" {
		db.Conn.Find("To", to, &mail)
	} else {
		db.Conn.All(&mail)
	}

	c.JSON(200, mail)
}
