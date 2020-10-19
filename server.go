package urlzap

import "net/http"

// Handler returns HTTP handler which contains set redirects.
func Handler(conf Config) func(http.ResponseWriter, *http.Request) {
	mux := http.NewServeMux()
	if err := Read("", conf.URLs, HTTPMuxCallback(conf.HTTP.BasePath, mux)); err != nil {
		panic(err)
	}

	return mux.ServeHTTP
}

// Server HTTP server with redirect set-up.
type Server struct {
	handler func(http.ResponseWriter, *http.Request)
}

// NewServer returns an instance of Server.
func NewServer(config Config) *Server {
	return &Server{handler: Handler(config)}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler(w, r)
}
