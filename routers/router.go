package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	//外部查询路由
	beego.Router("/exquery/appDetail", &controllers.ExQueryController{},"Get:AppDetail")

	//管理系统路由
	beego.Router("/", &controllers.AdminController{}, "*:Main")
	beego.Router("/admin/appDetail", &controllers.AdminController{},"Get:AppDetail" )
	beego.AutoRouter(&controllers.AdminController{})

	//fmt.Println(beego.URLFor("AdminController.AppDetail"))
}
