package mail

import (
	"github.com/mgerb/tmail/db"
	log "github.com/sirupsen/logrus"
)

type Mail struct {
	ID      int    `storm:"increment" json:"id"`
	From    string `json:"from"`
	To      string `storm:"index" json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (m *Mail) Save() {
	err := db.Conn.Save(m)

	if err != nil {
		log.Error(err)
	}
}
