package main
import (
	"net/http"
	_ "fmt"
	_ "io/ioutil"
	"log"
	"github.com/agouriou/dlserver/logger"
	"os"
	"flag"
	"fmt"
)

var mainLogger *logger.AggregateLogger

// Serve files from contents/
func rootHandler(contentDir *string, w http.ResponseWriter, r *http.Request) {
	filename := *contentDir + "/" + r.URL.Path[1:]
	mainLogger.Printf("Serving file [%s] to [%s]", filename, r.RemoteAddr)
	http.ServeFile(w, r, filename)
}

func main() {
	// handle args
	contentDir := flag.String("contentDir", "/tmp/dlserver-content", "base dir for content to be served")
	serverPort := flag.Int("port", 8080, "server's port")
	flag.Parse()

	// setup loggers
	stdoutLogger := log.New(os.Stdout, "", log.Lshortfile)
	httpLogger := logger.NewHttpLogger("http://localhost:9003/")
	mainLogger = logger.NewAggregateLogger(stdoutLogger, httpLogger)

	// ready to serve
	log.Printf("To get options, rerun with the '-h' flag")
	log.Printf("Server starting and listening on port %d...", *serverPort)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		rootHandler(contentDir, w, r)
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *serverPort), nil); err != nil {
		log.Fatal(err)
	}
}