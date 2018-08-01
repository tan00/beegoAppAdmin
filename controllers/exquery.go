package controllers

import (
	. "../errcode"
	"../models"
	"github.com/astaxie/beego/orm"
	"strconv"
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

//根据 userapp 的id 获取userapp的其他信息
func (c *ExQueryController) AppDetail() {
	result := make(map[string]interface{})

	var (
		Appid          int = 0
		err            error
		uapp           models.UserApp
		sysapisNameStr string
		sysapis        []models.SysApi
	)

	Appid, err = strconv.Atoi(c.GetString("appid"))
	if err != nil {
		result["error_msg"] = "Invalid parameter"
		result["error_code"] = ErrInvalidPara
		goto respon
	}

	uapp, err = c.GetUserApp(Appid)
	if err != nil {
		result["error_msg"] = "Invalid Appid"
		result["error_code"] = ErrInvalidAppid
		goto respon
	}

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
	c.Data["json"] = result
	c.ServeJSON()
}