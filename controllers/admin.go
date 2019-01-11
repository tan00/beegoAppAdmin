package controllers

import (
	"fmt"
	//fmt"
	//	"fmt"
	//"strconv"
	"strings"
	"time"

	"strconv"

	"../lib"
	"../models"
	"../util"
)

type AdminController struct {
	baseController
}

//配置信息
func (c *AdminController) Config() {
	_t := c.GetSession("user")
	usr := _t.(models.User)
	if !usr.IsSupperUser() {
		c.Ctx.WriteString("当前用户非超级管理员")
	} else {
		c.Approve()
	}
}

//后台用户登录
func (c *AdminController) Login() {

	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{Username: username}
		c.o.Read(&user, "username")

		if user.Password == "" {
			c.History("账号不存在", "")
		}

		if util.Md5(password) != strings.Trim(user.Password, " ") {
			c.History("密码错误", "")
		}

		user.Lastlogintime = time.Now()
		if _, err := c.o.Update(&user); err != nil {
			c.History("登录异常", "")
		} else {
			c.SetSession("user", user)
			c.History("登录成功", "/admin/main.html")
		}
	}
	c.TplName = c.controllerName + "/login.html"
}

//new user register
func (c *AdminController) Reg() {

	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		email := c.GetString("email")
		company := c.GetString("company")

		var newuser models.User
		newuser.Username = username
		newuser.Password = util.Md5(password)
		newuser.Email = email
		newuser.Company = company
		newuser.Createtime = time.Now()

		// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
		if created, _, err := c.o.ReadOrCreate(&newuser, "Username"); err == nil {
			if created {
				c.History("注册成功", "/admin/main.html")
			} else {
				c.History("用户已存在", "/admin/main.html")
			}
		} else {
			c.History("注册失败", "/admin/main.html")
		}
	} else {
		c.TplName = c.controllerName + "/reg.html"
	}
}

func (c *AdminController) Logout() {
	c.DestroySession()
	c.History("退出登录", "/admin/login.html")
}

//单页
func (c *AdminController) About() {
	c.Ctx.WriteString("About")
}

//添加api
func (c *AdminController) Appadd() {

	if c.Ctx.Request.Method == "POST" {
		var userapp models.UserApp
		//get sysapid
		sysNames, _ := models.GetAllApiNames()

		//遍历所有系统api的Name , 如果和用户请求的名称相同， 则添加到 UserApp的 SysApis中
		for _, sysApiName := range sysNames {
			val := c.GetString(sysApiName)
			if val == "on" {
				sysapi, _ := models.GetSysApiByName(sysApiName)
				userapp.SysApis += strconv.Itoa(sysapi.Id)
				userapp.SysApis += ","
			}
		}

		//get username
		userapp.Owner = c.GetSession("user").(models.User).Username
		//get userapp comment
		userapp.Commet = c.GetString("Commet")
		//get userapi name
		userapp.Name = c.GetString("Name")
		//generate key
		userapp.Appkey = lib.GenRandKey(16)
		userapp.SecretKey = lib.GenRandKey(16)

		//  add userApp
		userappID, _ := models.AddUserApp(&userapp)
		//add User.UserApps
		models.AddUserAppsID(userapp.Owner, int(userappID))

		c.History("申请已提交", "/admin/appadd.html")
	} else {
		names, _ := models.GetAllApiNames()
		c.Data["SysApis"] = names
		c.TplName = c.controllerName + "/appadd.html"
	}
}

//列出用户的APP信息
func (c *AdminController) List() {
	user := c.GetSession("user").(models.User).Username
	apps, _ := models.GetAllUerAppbyUserName(user)
	c.Data["Apps"] = apps
	c.SetSession("Apps", apps)
	c.TplName = c.controllerName + "/list.html"
}

//APP详细信息
func (c *AdminController) AppDetail() {
	//user := c.GetSession("user").(models.User).Username
	//apps , _ := models.GetAllUerAppbyUserName(user)
	//c.Data["Apps"] = apps

	Appid, err := strconv.Atoi(c.GetString("appid"))
	if err != nil {
		c.History("appid错误", "")
	}

	Apps := c.GetSession("Apps").([]models.UserApp)
	var app models.UserApp
	for _, app = range Apps {
		if app.Id == Appid {
			c.Data["UserApp"] = app
			break
		}
	}

	sysapis, _ := models.GetAllApiByIds(app.SysApis)
	c.Data["SysApis"] = sysapis

	c.TplName = c.controllerName + "/appDetail.html"
}

//主页
func (c *AdminController) Main() {
	session := c.GetSession("user")
	if session != nil {
		user := session.(models.User)
		c.Data["Username"] = user.Username
	}
	c.TplName = c.controllerName + "/main.html"
}

//Approve 修改userApp授权
func (c *AdminController) Approve() {

	if c.Ctx.Request.Method == "POST" {
		c.History("申请已提交", "/admin/appadd.html")
	} else {
		UserApps, queryErr := models.GetAllUerApp()
		if queryErr != nil {
			c.History(fmt.Sprint(queryErr), "")
		}
		c.Data["UserApps"] = UserApps
		c.TplName = c.controllerName + "/Approve.html"
	}
}

//上传接口
// func (c *AdminController) Upload() {
// 	f, h, err := c.GetFile("uploadname")
// 	result := make(map[string]interface{})
// 	img := ""
// 	if err == nil {
// 		exStrArr := strings.Split(h.Filename, ".")
// 		exStr := strings.ToLower(exStrArr[len(exStrArr)-1])
// 		if exStr != "jpg" && exStr != "png" && exStr != "gif" {
// 			result["code"] = 1
// 			result["message"] = "上传只能.jpg 或者png格式"
// 		}
// 		img = "/static/upload/" + util.UniqueId() + "." + exStr
// 		c.SaveToFile("upFilename", img) // 保存位置在 static/upload, 没有文件夹要先创建
// 		result["code"] = 0
// 		result["message"] = img
// 	} else {
// 		result["code"] = 2
// 		result["message"] = "上传异常" + err.Error()
// 	}
// 	defer f.Close()
// 	c.Data["json"] = result
// 	c.ServeJSON()
// }
