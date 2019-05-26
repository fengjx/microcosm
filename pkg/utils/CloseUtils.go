package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
)

func Close(c io.ReadCloser) {
	err := c.Close()
	if err != nil {
		log.Error("close error", c, err)
	}
}
