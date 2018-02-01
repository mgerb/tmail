package smtpserver

import (
	"bytes"
	"net"
	"net/mail"

	mymail "github.com/mgerb/tmail/mail"
	"github.com/mhale/smtpd"
	log "github.com/sirupsen/logrus"
)

func Start() {
	log.Error(smtpd.ListenAndServe("0.0.0.0:25", mailHandler, "tmail", ""))
}

func mailHandler(origin net.Addr, from string, to []string, data []byte) {
	msg, err := mail.ReadMessage(bytes.NewReader(data))

	if err != nil {
		log.Error(err)
		return
	}

	subject := msg.Header.Get("Subject")

	m := &mymail.Mail{
		Content: string(data),
		From:    from,
		To:      to[0],
		Subject: subject,
	}

	m.Save()
}
