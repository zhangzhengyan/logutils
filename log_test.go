package logutils

import (
	"testing"
)

func Test(t *testing.T) {
	var LogPath = "/app/log"
	var LogLevel = "info"
	lv, std := ParseLevel(LogLevel)
	if std {
		LogPath = ""
	}
	InitLog(LogPath, lv)
}

