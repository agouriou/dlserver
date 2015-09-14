package logger


// AggregateLogger allows to aggregate several loggers together
// when a call to an AggregateLogger is done, this call is made
// on each of the sub-loggers
type AggregateLogger struct {
	loggers []BasicLogger
}


func NewAggregateLogger(loggers... BasicLogger) (aggLogger *AggregateLogger){
	aggLogger = new(AggregateLogger)
	for _, logger := range loggers {
		aggLogger.loggers = append(aggLogger.loggers, logger)
	}
	return aggLogger
}



func (logger *AggregateLogger) Print(v ...interface{}){
	logger.applyToEachLogger(func(l BasicLogger){
		l.Print(v...)
	})
}

func (aggLogger *AggregateLogger) Printf(format string, v ...interface{}){
	aggLogger.applyToEachLogger(func(l BasicLogger){
		l.Printf(format, v...)
	})
}

func (aggLogger *AggregateLogger) Println(v ...interface{}){
	aggLogger.applyToEachLogger(func(l BasicLogger){
		l.Println(v...)
	})
}

type logFunc func (l BasicLogger)

// applyToEachLogger applies funcToApply to each of the sub-loggers of aggLogger
func (aggLogger *AggregateLogger) applyToEachLogger(funcToApply logFunc){
	for _, logger := range aggLogger.loggers {
		funcToApply(logger)
	}
}