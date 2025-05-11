package middleware

import "net/http"

type WrapResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func NewWrapResponseWriter(w http.ResponseWriter, protoMajor int) *WrapResponseWriter {
	return &WrapResponseWriter{ResponseWriter: w, status: http.StatusOK}
}

func (w *WrapResponseWriter) Status() int {
	return w.status
}

func (w *WrapResponseWriter) WriteHeader(code int) {
	if !w.wroteHeader {
		w.status = code
		w.ResponseWriter.WriteHeader(code)
		w.wroteHeader = true
	}
}

func (w *WrapResponseWriter) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}
