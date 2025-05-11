package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/DimDimi4-and-fixiki/auth-back/internal/metrics"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func PrometheusHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		// Use request path instead of route pattern for gRPC-Gateway routes
		path := r.Method + " " + r.URL.Path
		if chiCtx := chi.RouteContext(r.Context()); chiCtx != nil && chiCtx.RoutePattern() != "/*" {
			path = chiCtx.RoutePattern()
		}

		metrics.HTTPRequestsTotal.WithLabelValues(
			r.Method,
			path,
			http.StatusText(ww.Status()),
		).Inc()

		metrics.HTTPRequestDuration.WithLabelValues(
			r.Method,
			path,
		).Observe(time.Since(start).Seconds())
	})
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, req)

		st, _ := status.FromError(err)
		metrics.GRPCRequestsTotal.WithLabelValues(
			info.FullMethod,
			st.Code().String(),
		).Inc()

		metrics.GRPCRequestDuration.WithLabelValues(
			info.FullMethod,
		).Observe(time.Since(start).Seconds())

		return resp, err
	}
}
