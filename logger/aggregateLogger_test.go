package logger_test
import (
	"testing"
	"github.com/agouriou/dlserver/logger"
	"log"
	"bytes"
	"strings"
)


func TestPrint_OneLogger(t *testing.T){
	var out bytes.Buffer
	baseLogger := log.New(&out, "", log.Lshortfile)
	mainLog := logger.NewAggregateLogger(baseLogger)
	mainLog.Printf("toto %s", "ta")
	if s := out.String(); strings.Index(s, "toto ta") == -1 {
		t.Errorf("Expected out to be [%s] but was [%s]", "toto", s)
	}
}


func TestPrintln_TwoLogger(t *testing.T){
	var out bytes.Buffer
	var out2 bytes.Buffer
	baseLogger := log.New(&out, "", log.Lshortfile)
	baseLogger2 := log.New(&out2, "", log.Lshortfile)
	mainLog := logger.NewAggregateLogger(baseLogger, baseLogger2)
	mainLog.Printf("toto %s", "ta")
	if s, s2 := out.String(), out2.String(); strings.Index(s, "toto ta") == -1 || strings.Index(s2, "toto ta") == -1 {
		t.Errorf("Expected out to be [%s] but was [%s]", "toto", s)
	}
}