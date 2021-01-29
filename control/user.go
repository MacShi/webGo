package control

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"webGo/model"
	"webGo/service"
	"webGo/utils"
)
var (
	cookieNameForSessionID = "sessionId"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})

)

func Register(ctx iris.Context)  {
	loginInfo := model.UserLogin{}
	err:=ctx.ReadJSON(&loginInfo)
	if nil!=err{
		ctx.JSON(utils.ResultUtil.Failure(err.Error()))
	}else {
		a,err := service.Register(loginInfo)
		if nil!=err{
			ctx.JSON(utils.ResultUtil.Failure(err.Error()))
		}else {
			ctx.JSON(utils.ResultUtil.Success(a))
		}

	}

}


func Login(ctx iris.Context)  {
	session := sess.Start(ctx)
	session.Set("authenticated",true)
	loginInfo := model.UserLogin{}
	err:=ctx.ReadJSON(&loginInfo)
	if err!=nil{
		ctx.JSON(utils.ResultUtil.Failure(err.Error()))
	}else {
		aa,err:=service.UserLogin(loginInfo)
		if nil!=err{
			ctx.JSON(utils.ResultUtil.Failure(err.Error()))
		}else {
			ctx.JSON(utils.ResultUtil.Success(aa))
		}
	}
}

func LogOut(ctx iris.Context)  {
	session :=sess.Start(ctx)
	session.Set("authenticated",false)
	session.Delete("authenticated")
	session.Destroy()
	ctx.HTML("Logout")

}
