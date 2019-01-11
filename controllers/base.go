package controllers

import (
	//	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()

	if p.controllerName == "admin" && (p.actionName != "login" && p.actionName != "reg") {
		if p.GetSession("user") == nil {
			p.History("未登录", "/admin/login")
			p.Ctx.WriteString(p.controllerName + "===" + p.actionName)
		}
	}

}

func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}

//获取用户IP地址
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}


