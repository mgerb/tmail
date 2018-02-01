package main

import (
	"os"

	"github.com/asdine/storm"
	"github.com/mgerb/tmail/db"
	"github.com/mgerb/tmail/smtpserver"
	"github.com/mgerb/tmail/webserver"
	log "github.com/sirupsen/logrus"
)

//DB - database instance
var DB *storm.DB

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	db.Init()
	defer db.Conn.Close()

	// start webserver
	go webserver.Start()

	// start smtp server
	smtpserver.Start()
}
