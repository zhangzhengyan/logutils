package logutils

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func NewHook(logName string) (log.Hook, error) {
	writer, err := rotatelogs.New(
		logName+".%Y%m%d",
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithLinkName(logName),
		rotatelogs.WithRotationCount(7),
	)
	if err != nil {
		return nil, err
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{
		TimestampFormat:"2006-01-02 15:04:05",  //时间格式
		FullTimestamp:true,
	})
	return lfsHook, nil
}

func ParseLevel(lvl string) (log.Level, bool) {
	s := strings.ToLower(lvl)
	var lv log.Level
	if strings.EqualFold(s, "panic") {
		lv = log.PanicLevel
	} else if strings.EqualFold(s, "fatal") {
		lv = log.FatalLevel
	} else if strings.EqualFold(s, "error") {
		lv = log.ErrorLevel
	} else if strings.EqualFold(s, "warn") || strings.EqualFold(s, "warning") {
		lv = log.WarnLevel
	} else if strings.EqualFold(s, "info") {
		lv = log.InfoLevel
	} else if strings.EqualFold(s, "debug") {
		lv = log.DebugLevel
	} else {
		lv = log.TraceLevel
	}
	if strings.Contains(s, "std") {
		return lv, true
	} else {
		return lv, false
	}
}

func InitLog(logFileName string, level log.Level) {
	hook, _ := NewHook(logFileName)
	log.SetLevel(level)
	log.SetReportCaller(true)
	if logFileName == "" {
		log.SetOutput(os.Stdout)
	} else {
		if hook != nil {
			log.AddHook(hook)
		} else {
			log.SetOutput(os.Stdout)
		}
	}
}
