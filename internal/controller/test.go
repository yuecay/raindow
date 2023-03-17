package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "rainbow/api/v1"
)

var (
	Test = cTest{}
)

type cTest struct{}

func (c *cTest) Hello(ctx context.Context, req *v1.TestReq) (res *v1.TestRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("你好呀 go！")
	return
}
