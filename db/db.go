package db

import (
	"github.com/asdine/storm"
	log "github.com/sirupsen/logrus"
)

var Conn *storm.DB

func Init() {

	var err error
	Conn, err = storm.Open("mail.db")

	if err != nil {
		log.Fatal(err)
	}

}
