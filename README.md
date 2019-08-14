## tollbooth_beego

[Beego](https://github.com/https://github.com/astaxie/beego) middleware for rate limiting HTTP requests.


## Five Minutes Tutorial

```
package main

import (
	"github.com/didip/tollbooth"
	"github.com/astaxie/beego"
	"github.com/didip/tollbooth/limiter"
	"github.com/tollbooth_beego"
)

var res *limiter.Limiter

func main() {

	res = tollbooth.NewLimiter(10, nil)

	res.SetMessage("已超过最大连接数，请稍后测试")

	beego.Router("/v1/aa",&Admin{},"get:Demo")

	beego.InsertFilter("/v1/*",beego.BeforeExec,tollbooth_beego.LimitHandler(res))

	beego.Run(":12345")


}



type Admin struct {

	beego.Controller
}

func (this *Admin) Demo()  {

	this.Ctx.Output.Body([]byte("testing"))
}

```
