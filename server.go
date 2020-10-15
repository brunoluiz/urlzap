package urlzap

import "net/http"

// HTTPHandler returns HTTP handler which contains set redirects.
func HTTPHandler(conf Config) func(http.ResponseWriter, *http.Request) {
	mux := http.NewServeMux()
	if err := Read("", conf.URLs, HTTPRedirectCallback(conf.HTTP.Path, mux)); err != nil {
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
	return &Server{handler: HTTPHandler(config)}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler(w, r)
}
