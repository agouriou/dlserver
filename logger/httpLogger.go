package logger
import (
	"net/http"
	"fmt"
	"log"
	"strings"
)


type HttpLogger struct {
	url string
	bodyType string
}

func NewHttpLogger(url string) (httpLogger *HttpLogger){
	httpLogger = &HttpLogger{url: url, bodyType: "text/plain"}
	return httpLogger
}

func HandleError(resp *http.Response, err error){
	if err != nil {
		log.Printf("Error on http call %s\n", err)
	}
}

func (logger *HttpLogger) GetPostAction() (PostFunction func(s string)){
	return func(s string) {
		HandleError(http.Post(logger.url, logger.bodyType, strings.NewReader(s)))
	}
}

func (logger *HttpLogger) Print(v ...interface{}){
	logger.GetPostAction()(fmt.Sprint(v...))
}

func (logger *HttpLogger) Printf(format string, v ...interface{}){
	logger.GetPostAction()(fmt.Sprintf(format, v...))
}

func (logger *HttpLogger) Println(v ...interface{}){
	logger.GetPostAction()(fmt.Sprintln(v...))
}