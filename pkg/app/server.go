package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/user-manager-go/pkg/dependencies"
	"log"
	"net/http"
	"sync"
)

type RouteConfig struct {
	Path       string
	Method     string
	Handler    HandlerFunc[any]
	Middleware []gin.HandlerFunc
}

type HttpServer interface {
	Middleware(middlewareFunc gin.HandlerFunc)
	Get(path string, handler HandlerFunc[any])
	Post(path string, handler HandlerFunc[any])
	Put(path string, handler HandlerFunc[any])
	Delete(path string, handler HandlerFunc[any])
	Routes(routes []RouteConfig)
	Start(host string, port string)
}

type httpServer struct {
	*gin.Engine
	*dependencies.Dependency
}

var instanceHttpServer *httpServer
var httpServerOnce sync.Once

func NewHttpServer(deps *dependencies.Dependency) HttpServer {
	httpServerOnce.Do(func() {
		instanceHttpServer = &httpServer{
			Engine:     gin.New(),
			Dependency: deps,
		}
	})
	return instanceHttpServer
}

func (s *httpServer) Middleware(middlewareFunc gin.HandlerFunc) {
	s.Use(middlewareFunc)
}

func (s *httpServer) Get(path string, handler HandlerFunc[any]) {
	s.GET(path, RouteHandler(s.Dependency, handler))
}

func (s *httpServer) Post(path string, handler HandlerFunc[any]) {
	s.POST(path, RouteHandler(s.Dependency, handler))
}

func (s *httpServer) Put(path string, handler HandlerFunc[any]) {
	s.PUT(path, RouteHandler(s.Dependency, handler))
}

func (s *httpServer) Delete(path string, handler HandlerFunc[any]) {
	s.DELETE(path, RouteHandler(s.Dependency, handler))
}

func (s *httpServer) Routes(routes []RouteConfig) {
	for _, r := range routes {
		group := s.Group(r.Path)
		group.Use(r.Middleware...)
		switch r.Method {
		case http.MethodGet:
			group.GET("", RouteHandler(s.Dependency, r.Handler))
		case http.MethodPost:
			group.POST("", RouteHandler(s.Dependency, r.Handler))
		case http.MethodPut:
			group.PUT("", RouteHandler(s.Dependency, r.Handler))
		case http.MethodDelete:
			group.DELETE("", RouteHandler(s.Dependency, r.Handler))
		default:
			panic("Unsupported HTTP method: " + r.Method)
		}
	}
}

func (s *httpServer) Start(host string, port string) {
	err := s.Run(host + ":" + port)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}
