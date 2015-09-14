package logger


type BasicLogger interface{
  	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}