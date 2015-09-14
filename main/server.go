package main
import (
	"net/http"
	_ "fmt"
	_ "io/ioutil"
	"log"
	"github.com/agouriou/dlserver/logger"
	"os"
)

var mainLogger *logger.AggregateLogger

// Serve files from contents/
func rootHandler(w http.ResponseWriter, r *http.Request) {
	filename := "contents/" + r.URL.Path[1:];
	mainLogger.Printf("Serving file [%s] to [%s]", filename, r.RemoteAddr);
	http.ServeFile(w, r, filename);
}

func main() {
	stdoutLogger := log.New(os.Stdout, "", log.Lshortfile)
	httpLogger := logger.NewHttpLogger("http://localhost:9003/")
	mainLogger = logger.NewAggregateLogger(stdoutLogger, httpLogger)
	log.Printf("Server starting...");
	http.HandleFunc("/", rootHandler);
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}