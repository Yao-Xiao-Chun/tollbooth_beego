package tollbooth_beego

import (
	"github.com/astaxie/beego/context"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth"
	"github.com/astaxie/beego"
)

func LimitHandler(lmt *limiter.Limiter ) beego.FilterFunc {

	return func(ctx *context.Context) {

		httpError := tollbooth.LimitByRequest(lmt, ctx.ResponseWriter, ctx.Request)

		if httpError != nil {

			ctx.Output.Body([]byte(httpError.Message))

			return
		}

	}
}
