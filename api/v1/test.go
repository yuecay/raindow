package v1

import "github.com/gogf/gf/v2/frame/g"

type TestReq struct {
	g.Meta `path:"/test" tags:"test" method:"get" summary:"test one goframe project"`
}
type TestRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
