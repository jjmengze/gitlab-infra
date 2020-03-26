package hook

import "net/http"

// Server implements http.Handler. It validates incoming gitlab webhooks and
// then dispatches them to the appropriate plugins.
type Server struct{}

// ServeHTTP validates an incoming webhook and puts it into the event channel.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}