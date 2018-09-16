package vlog

import "log"

const (
	L_I = "info"
	L_D = "debug"
	L_E = "error"
)

var Loglevel string = "debug"

func SetLogLevel(level string)  {

	switch level {
	case L_I,L_D,L_E:
		Loglevel = level
		log.Println("set log level : ",level)
	default:
		log.Println("use default log info")
		Loglevel = "info"
	}

}

func Info(v ...interface{})  {

	if Loglevel != L_E {
		log.Println("INFO: ", v)
	}
}

func Debug( v ...interface{})  {
	if Loglevel == L_D{
		log.Println("DEBUG: ",v)
	}

}

func Error(v ...interface{})  {
		log.Fatalln("ERROR: ",v)
}