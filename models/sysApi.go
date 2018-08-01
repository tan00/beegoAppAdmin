package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
)

//Appinfo 所有系统API
type SysApi struct {
	Id       int
	Name     string `orm:"unique;size(256)" form:"Name"      valid:"Required"`
	Describe string `orm:"size(256)" form:"Describe"  valid:"Required"`
}

var (
	urlMap =  map[int]string{
		1:"https://aip.baidubce.com/rest/2.0/ocr/v1/idcard",
		2:"https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard",
		3:"https://aip.baidubce.com/rest/2.0/ocr/v1/general",
	}
)

func SysApiUrl(id int)string{
	return urlMap[id]
}

func init(){
	orm.RegisterModel(new(SysApi))
	beego.AddFuncMap("SysApiUrl",SysApiUrl)
}

func (u *SysApi) TableName() string {
	return "tb_SysAPI"
}

//添加sysApi
func AddApi(app *SysApi) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(app)
	return id, err
}



//删除sysApi
func DelApiById(Id int) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&SysApi{Id: Id})
	return status, err
}

func GetAllApiNames() (apiNames []string, err error) {
	o := orm.NewOrm()
	// 可以直接使用对象作为表名
	api := new(SysApi)
	qs := o.QueryTable(api) // 返回 QuerySeter

	var apis []SysApi
	_, queryerr := qs.All(&apis)
	if queryerr != nil {
		log.Printf("query sysapi error", queryerr)
		return nil, queryerr
	} else {
		for _, api := range apis {
			apiNames = append(apiNames, api.Name)
		}
	}
	return apiNames, err
}

func GetSysApiByName(appName string) (api SysApi, err error) {
	api = SysApi{Name: appName}
	o := orm.NewOrm()
	err = o.Read(&api, "Name")
	return api, err
}


func GetApiById(id int)(api SysApi , err error){
	api = SysApi{Id: id}
	o := orm.NewOrm()
	err = o.Read(&api, "Id")
	return api , err
}


func GetAllApiByIds(ids string)(apis []SysApi , err error){

	apiids := strings.Split(ids,",")
	for _,v := range apiids {
		id , _ := strconv.Atoi(v)

		api , err := GetApiById(id)
		if err == nil{
			apis = append(apis, api)
		}
	}
	return  apis , err
}


