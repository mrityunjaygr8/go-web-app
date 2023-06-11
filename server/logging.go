package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type customLogEntry struct {
	logFormatter customLogFormatter
	fields       logrus.Fields
}

func (lf customLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	lf.fields["status"] = status
	lf.fields["bytes"] = bytes

	lf.fields["elapsed"] = fmt.Sprintf("%d ns", elapsed.Microseconds())
	if extra != nil {
		lf.fields["extra"] = extra
	}
	lf.fields["response_header"] = header

	lf.logFormatter.logger.WithFields(lf.fields).Println()

}
func (lf customLogEntry) Panic(v interface{}, stack []byte) {
	lf.fields["extra"] = v
	lf.logFormatter.logger.WithFields(lf.fields).Println(string(stack))
}

type customLogFormatter struct {
	logger *logrus.Logger
}

func (mlf customLogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	mle := customLogEntry{
		logFormatter: mlf,
	}
	reqID := middleware.GetReqID(r.Context())
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	mle.fields = logrus.Fields{
		"reqID":          reqID,
		"host":           r.Host,
		"URI":            r.RequestURI,
		"remote":         r.RemoteAddr,
		"proto":          r.Proto,
		"scheme":         scheme,
		"method":         r.Method,
		"request_header": r.Header,
	}

	return mle
}
