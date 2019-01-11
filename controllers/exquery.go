package controllers

import (
	"strconv"

	. "../errcode"
	"../models"
	"github.com/astaxie/beego/orm"
)

type ExQueryController struct {
	baseController
}

//根据 userapp 的id 获取userapp的其他信息
func (c *ExQueryController) GetUserApp(id int) (uapp models.UserApp, err error) {
	o := orm.NewOrm()
	uapp = models.UserApp{Id: id}
	err = o.Read(&uapp, "Id")
	return uapp, err
}

//AppDetail 根据 userapp 的id 获取userapp的其他信息
func (c *ExQueryController) AppDetail() {
	result := make(map[string]interface{})
	errorBlock := make(map[string]interface{})

	var (
		Appid          int = 0
		err            error
		uapp           models.UserApp
		sysapisNameStr string
		sysapis        []models.SysApi
	)

	Appid, err = strconv.Atoi(c.GetString("appid"))
	if err != nil {
		errorBlock["error_msg"] = "Invalid parameter"
		errorBlock["error_code"] = ErrInvalidPara
		goto respon
	}

	uapp, err = c.GetUserApp(Appid)
	if err != nil {
		errorBlock["error_msg"] = "Invalid Appid"
		errorBlock["error_code"] = ErrInvalidAppid
		goto respon
	}

	errorBlock["error_msg"] = ""
	errorBlock["error_code"] = ErrNormal

	result["error"] = errorBlock
	result["AppID"] = uapp.Id
	result["Name"] = uapp.Name
	result["Appkey"] = uapp.Appkey
	result["Commet"] = uapp.Commet
	result["Createtime"] = uapp.Createtime
	result["Authorize"] = uapp.Authorize

	sysapis, err = models.GetAllApiByIds(uapp.SysApis)
	for _, v := range sysapis {
		sysapisNameStr += v.Name
		sysapisNameStr += ","
	}
	result["SysAPIS"] = sysapisNameStr

respon:
	c.Data["json"] = result
	c.ServeJSON()
}

//AppDetailFake 去掉数据库操作 查看性能指标
func (c *ExQueryController) AppDetailFake() {
	result := make(map[string]interface{})

	var (
		Appid          int = 0
		err            error
		uapp           models.UserApp
		sysapisNameStr string
		//sysapis        []models.SysApi
	)

	Appid, err = strconv.Atoi(c.GetString("appid"))
	if err != nil {
		result["error_msg"] = "Invalid parameter"
		result["error_code"] = ErrInvalidPara
		goto respon
	}

	// uapp, err = c.GetUserApp(Appid)
	// if err != nil {
	// 	result["error_msg"] = "Invalid Appid"
	// 	result["error_code"] = ErrInvalidAppid
	// 	goto respon
	// }

	uapp.Id = Appid
	result["error_code"] = ErrNormal
	result["appid"] = uapp.Id
	result["Name"] = uapp.Name
	result["Appkey"] = uapp.Appkey
	result["Commet"] = uapp.Commet
	result["Createtime"] = uapp.Createtime
	result["Authorize"] = uapp.Authorize

	// sysapis, err = models.GetAllApiByIds(uapp.SysApis)
	// for _, v := range sysapis {
	// 	sysapisNameStr += v.Name
	// 	sysapisNameStr += ","
	// }
	result["SysApis"] = sysapisNameStr

respon:
	c.Data["json"] = result
	c.ServeJSON()
}

//AppDetailDBonly 去掉网络连接， 仅和db交互 查看性能指标
func AppDetailDBonly() {
	result := make(map[string]interface{})

	var (
		Appid          int
		err            error
		uapp           models.UserApp
		sysapisNameStr string
		sysapis        []models.SysApi
		controler      ExQueryController
	)

	Appid = 1

	uapp, err = controler.GetUserApp(Appid)
	if err != nil {
		result["error_msg"] = "Invalid Appid"
		result["error_code"] = ErrInvalidAppid
		goto respon
	}

	result["error_msg"] = ""
	result["error_code"] = ErrNormal
	result["appid"] = uapp.Id
	result["Name"] = uapp.Name
	result["Appkey"] = uapp.Appkey
	result["Commet"] = uapp.Commet
	result["Createtime"] = uapp.Createtime
	result["Authorize"] = uapp.Authorize

	sysapis, err = models.GetAllApiByIds(uapp.SysApis)
	for _, v := range sysapis {
		sysapisNameStr += v.Name
		sysapisNameStr += ","
	}
	result["SysApis"] = sysapisNameStr

respon:
	//controler.Data["json"] = result
	//c.ServeJSON()
}
