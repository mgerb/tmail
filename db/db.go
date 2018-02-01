package db

import (
	"log"

	"github.com/asdine/storm"
)

var Conn *storm.DB

func Init() {

	var err error
	Conn, err = storm.Open("mail.db")

	if err != nil {
		log.Fatal(err)
	}

}
