package core_middleware_logger

import (
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// AccessLogMiddleware — поля одной записи access-лога (структурированный лог запроса).
type AccessLogMiddleware struct {
	Router   string
	Method   string
	Path     string
	Status   int
	Duration time.Duration
	Result   string
	Error    string
}

type responseRecorder struct {
	http.ResponseWriter
	status int
	result any
	err    error
}

type loggableResponseWriter interface {
	SetResult(any) any
	SetError(err error)
}

func SetError(w http.ResponseWriter, err error) {
	if rec, ok := w.(*responseRecorder); ok {
		rec.err = err
	}
}

func SetResult(w http.ResponseWriter, result any) {
	if rec, ok := w.(*responseRecorder); ok {
		rec.result = result
	}
}

func (rw *responseRecorder) SetResult(w http.ResponseWriter, result any) any {
	if lw, ok := w.(loggableResponseWriter); ok {
		lw.SetResult(result)
	}
	return result
}
func (rw *responseRecorder) SetError(w http.ResponseWriter, err error) {
	if lw, ok := w.(loggableResponseWriter); ok {
		lw.SetError(err)
	}
	rw.err = err
}

func (rw *responseRecorder) WriteHeader(code int) {
	if rw.status == 0 {
		rw.status = code
	}
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseRecorder) Write(b []byte) (int, error) {
	if rw.status == 0 {
		rw.status = http.StatusOK
	}
	return rw.ResponseWriter.Write(b)
}

func AccessLog(logger *zap.Logger, routerName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rec := &responseRecorder{ResponseWriter: w, status: 0, result: nil, err: nil}

			next.ServeHTTP(rec, r)

			status := rec.status
			if status == 0 {
				status = http.StatusOK
			}

			entry := AccessLogMiddleware{
				Router:   routerName,
				Method:   r.Method,
				Path:     r.URL.Path,
				Status:   status,
				Duration: time.Since(start),
				Result:   fmt.Sprintf("%v", rec.result),
				Error:    "",
			}
			if status >= http.StatusInternalServerError {
				entry.Result = "error"
				entry.Error = rec.err.Error()
			}

			logger.Info("http_access",
				zap.String("router", entry.Router),
				zap.String("method", entry.Method),
				zap.String("path", entry.Path),
				zap.Int("status", entry.Status),
				zap.Duration("duration", entry.Duration),
				zap.String("result", entry.Result),
				zap.String("error", entry.Error),
			)
		})
	}
}
