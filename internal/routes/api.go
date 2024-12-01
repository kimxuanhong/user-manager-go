package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimxuanhong/user-manager-go/internal/routes/route"
	"github.com/kimxuanhong/user-manager-go/pkg/app"
	"net/http"
)

func PartnerRoutes(route route.PartnerRoute) []app.RouteConfig {
	return []app.RouteConfig{
		{
			Path:    "/partner/:id",
			Method:  http.MethodPost,
			Handler: route.GetUserByPartnerId,
			Middleware: []gin.HandlerFunc{
				app.LogResponseMiddleware(),
			},
		},
		{
			Path:       "/partner/all",
			Method:     http.MethodPost,
			Handler:    route.GetAllUser,
			Middleware: []gin.HandlerFunc{},
		},
	}
}
