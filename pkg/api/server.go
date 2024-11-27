package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/user-manager-go/pkg/config"
	"log"
	"sync"
)

type Route interface {
	RouteHandler(ctx *Context, whenDone Handler[any])
}

type HttpServer interface {
	Get(path string, route Route)
	Post(path string, route Route)
	Start(host string, port string)
}

type httpServer struct {
	*gin.Engine
	deps *config.Dependencies
}

var instanceHttpServer *httpServer
var httpServerOnce sync.Once

func NewHttpServer(deps *config.Dependencies) HttpServer {
	httpServerOnce.Do(func() {
		instanceHttpServer = &httpServer{
			Engine: gin.New(),
			deps:   deps,
		}
	})
	return instanceHttpServer
}

func (s *httpServer) Get(path string, route Route) {
	s.GET(path, RouteHandler(s.deps, route.RouteHandler))
}

func (s *httpServer) Post(path string, route Route) {
	s.POST(path, RouteHandler(s.deps, route.RouteHandler))
}

func (s *httpServer) Start(host string, port string) {
	err := s.Run(host + ":" + port)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}
