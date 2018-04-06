package log

import (
	"fmt"
	"time"
)

// Logger is struct of log
type Logger struct {
	category string

	pool []map[string]interface{}

	flush time.Duration
}

// New instantiate Loggger
func New() *Logger {
	return &Logger{
		flush: flushInterval,
	}
}

var (
	logger = New()

	flushInterval = time.Second * 10 //日志刷新间隔
)

func init() {
	go func() {
		ticker := time.NewTicker(logger.flush)

		for {
			select {
			case <-ticker.C:
				logger.Flush()
			}
		}

	}()
}

func (l *Logger) Write(data map[string]interface{}) {
	l.pool = append(l.pool, data)
}

func Write(data map[string]interface{}) {
	logger.Write(data)
}

func (l *Logger) write() {
	fmt.Println("start write log...")
	for i, d := range l.pool {
		fmt.Printf("write index:%d with data: %+v \n", i, d)
	}
	fmt.Println("end write log...")
}

// Flush write log immediately
func (l *Logger) Flush() {
	l.write()
	l.pool = []map[string]interface{}{}
}
