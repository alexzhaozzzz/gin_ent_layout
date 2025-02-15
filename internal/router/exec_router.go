package router

import (
	"back-end/merchant/pkg/bootstrap"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(rg *gin.RouterGroup, app *bootstrap.App), 0)
	routerCheckRole   = make([]func(rg *gin.RouterGroup, app *bootstrap.App, gHanFun ...gin.HandlerFunc), 0)
)

func ExecRouter(r *gin.Engine, app *bootstrap.App, gHanFun ...gin.HandlerFunc) *gin.Engine {
	// 无需认证的路由
	noCheckRoleRouter(r, app)
	// 需要认证的路由
	checkRoleRouter(r, app, gHanFun...)

	return r
}

// 无需认证的路由示例
func noCheckRoleRouter(r *gin.Engine, app *bootstrap.App) {
	// 可根据业务需求来设置接口版本
	rg := r.Group("/merchant/api/v1")
	for _, f := range routerNoCheckRole {
		f(rg, app)
	}
}

// 需要认证的路由示例
func checkRoleRouter(r *gin.Engine, app *bootstrap.App, gHanFun ...gin.HandlerFunc) {
	// 可根据业务需求来设置接口版本
	rg := r.Group("/merchant/api/v1")
	for _, f := range routerCheckRole {
		f(rg, app, gHanFun...)
	}
}
