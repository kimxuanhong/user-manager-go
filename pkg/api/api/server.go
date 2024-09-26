package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

type HandleRouter func(group *gin.RouterGroup)

type Server interface {
	Router(relativePath string, routerFunc HandleRouter)
	Start()
}

type server struct {
	*gin.Engine
	Cfg *Config `inject:""`
}

var instanceServer *server
var serverOnce sync.Once

func NewServer() Server {
	serverOnce.Do(func() {
		instanceServer = &server{
			Engine: gin.New(),
		}
	})
	return instanceServer
}

func (r *server) Router(relativePath string, routerFunc HandleRouter) {
	group := r.Group(relativePath)
	routerFunc(group)
}

func (r *server) Start() {
	// Khởi động server trên cổng được cấu hình
	err := r.Run(fmt.Sprintf("%s:%d", r.Cfg.Server.Host, r.Cfg.Server.Port))
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
