package models

import (
	"errors"
	"log"
	"time"

	"strconv"

	. "../lib"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

//用户表
type User struct {
	Id            int
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(50);MinSize(6)"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
	Company       string    `orm:"size(256)" form:"Email" valid:"Email"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add"`
	UserApps      string    `orm:"type(text);" valid:"Required"` //存放多个 UserApp的Id
	Level         int8      `orm:"default(0)"`                   //用户权限划分 普通用户0  超级管理员1
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "tb_User"
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

func (u *User) IsSupperUser() bool {
	return u.Level == 1
}

//验证用户信息
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

//添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Email = u.Email

	id, err := o.Insert(user)
	return id, err
}

func DelUserById(Id int) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}

//GetUserById  GetUserById
func GetUserById(id int) (user User) {
	user = User{Id: id}
	o := orm.NewOrm()
	o.Read(&user, "Id")
	return user
}

//AddUserAppsID add UserApps field
func AddUserAppsID(username string, ID int) {
	user := User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "username")
	user.UserApps += strconv.Itoa(ID)
	user.UserApps += ","
	o.Update(&user)
	return
}

//ModifyAppAuth 修改 UserApp Authorize字段  ,可禁用 启用 UserApp
func ModifyAppAuth(ID int, auth int8) {
	uapp := UserApp{Id: ID}
	o := orm.NewOrm()
	o.Read(&uapp, "Id")
	uapp.Authorize = auth
	o.Update(&uapp)
	return
}
