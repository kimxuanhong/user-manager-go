package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/user-manager-go/pkg/config"
	"log"
	"sync"
)

type HttpServer interface {
	Get(path string, handler HandlerFunc[any])
	Post(path string, handler HandlerFunc[any])
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

func (s *httpServer) Get(path string, handler HandlerFunc[any]) {
	s.GET(path, RouteHandler(s.deps, handler))
}

func (s *httpServer) Post(path string, handler HandlerFunc[any]) {
	s.POST(path, RouteHandler(s.deps, handler))
}

func (s *httpServer) Start(host string, port string) {
	err := s.Run(host + ":" + port)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}
