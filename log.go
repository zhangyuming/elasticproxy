package main

import "github.com/labstack/gommon/log"

const (
	L_I = "info"
	L_D = "debug"
	L_E = "error"
)

type Log struct {
	level string
}

func GetLog(level string) Log  {
	if level == L_E || level == L_D || level == L_I {
		return Log{
			level:level,
		}
	}else{
		log.Warn("unsuport log level : " , level)
		return Log{
			level:"info",
		}
	}
}

func (l *Log) Info(v ...interface{})  {

	if l.level != L_E {
		log.Print(v...)
	}
}

func (l *Log) Debug( v ...interface{})  {
	if l.level == L_D{
		log.Print(v...)
	}

}

func (l *Log) Error(v ...interface{})  {
		log.Error(v...)
}