package cmd

import (
	"context"
	"rainbow/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"rainbow/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
			})
			s.Group("/test", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.SensitiveMiddlewaare)
				group.Bind(
					controller.Test,
				)
			})
			s.Run()
			return nil
		},
	}
)
