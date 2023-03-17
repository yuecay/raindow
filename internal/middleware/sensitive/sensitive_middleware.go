package sensitive

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	wordsChan chan string
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func SensitiveMiddlewaare(r *ghttp.Request) {

	r.Middleware.Next()
}
