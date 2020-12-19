package urlzap

import (
	"context"
	"net/http"
)

// Handler returns HTTP handler which contains set redirects.
func Handler(ctx context.Context, conf Config) func(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	if err := Read(ctx, "", conf.URLs, HTTPMuxCallback(conf.HTTP.BasePath, mux)); err != nil {
		panic(err)
	}

	return mux.ServeHTTP
}

// Server HTTP server with redirect set-up.
type Server struct {
	handler func(http.ResponseWriter, *http.Request)
}

// NewServer returns an instance of Server.
func NewServer(ctx context.Context, config Config) *Server {
	return &Server{handler: Handler(ctx, config)}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler(w, r)
}
