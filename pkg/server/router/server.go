package router

import (
	v1 "github.com/clarechu/infrapulse/pkg/server/router/v1"
	"github.com/emicklei/go-restful/v3"
)

type Server struct {
	RestfulCont *restful.Container
}

// NewServer initializes and configures a kubelet.Server object to handle HTTP requests.
func NewServer() Server {
	server := Server{
		RestfulCont: restful.NewContainer(),
	}

	server.RestfulCont.Add(DefaultHandlers())
	return server
}

// DefaultHandlers registers the default set of supported HTTP request
// patterns with the restful Container.
func DefaultHandlers() *restful.WebService {
	ws := new(restful.WebService)
	ws.Route(
		ws.GET("/healthz").To(v1.Health).
			Doc("健康检查").
			Operation("health"))
	return ws
}
