package models

import (
	"errors"
	"log"
	"time"

	. "../lib"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
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
	UserApps      string    `orm:"type(text);" valid:"Required"`  //存放多个 UserApp的Id
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

func GetUserById(id int) (user User) {
	user = User{Id: id}
	o := orm.NewOrm()
	o.Read(&user, "Id")
	return user
}

func AddUserAppsID(username string, ID int ) {
	user := User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "username")
	user.UserApps += strconv.Itoa(ID)
	user.UserApps += ","
	o.Update(&user)
	return
}
