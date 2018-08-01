package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

//UserApp 用户API

type UserApp struct {
	Id     int
	Name   string `orm:"size(256)"   valid:"Required"`                 //应用名称 可以重复
	Commet string `orm:"size(256)"  form:"Describe"  valid:"Required"` //用户对应用的描述
	Owner  string `orm:"size(256)"   valid:"Required"`                 //app用户
	Appkey     string    `orm:"size(256)"`
	SecretKey  string    `orm:"size(256)"`
	Createtime time.Time `orm:"type(datetime);auto_now_add"`
	Authorize  int8      `orm:"size(1)" valid:"Required"`     //-1禁用  0审核中  1 启用
	SysApis    string    `orm:"type(text);" valid:"Required"` //存放多个 SysApi 的Id
}

var (
	statMap = map[int8]string{
		-1: "已禁用",
		0:  "审核中",
		1:  "已启用",
	}
)

func AuthString(status int8) string {
	return statMap[status]
}

func init() {
	orm.RegisterModel(new(UserApp))

	beego.AddFuncMap("AuthString", AuthString)
}

func (u *UserApp) TableName() string {
	return "tb_UserApp"
}

//添加app
func AddUserApp(app *UserApp) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(app)
	return id, err
}

//删除app
func DelUserAppById(Id int) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&UserApp{Id: Id})
	return status, err
}

// GetAllUerAppbyUserName
func GetAllUerAppbyUserName(user string) (userApps []UserApp, err error) {
	o := orm.NewOrm()
	_t := new(UserApp)
	qs := o.QueryTable(&_t)

	userApps = make([]UserApp, 0)
	qs.Filter("Owner__exact", user) // WHERE Owner = $user
	_, err = qs.All(&userApps)
	return userApps, err
}
