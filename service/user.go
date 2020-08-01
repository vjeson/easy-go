package service

import (
	. "demo/db"
	"demo/entity"
)

type Users entity.Users

func (Users) TableName() string  {
	return "sys_user"
}

func (user *Users) ChekcLogin() *Users {
	u := new(Users)
	Db.Where(user).Take(u)
	return u
}

func (user *Users) Insert() (id int, err error) {
	result := Db.Create(&user)
	id = user.Id
	if result != nil {
		err = result.Error
		return
	}
	return
}

func (user *Users) Users() (users []Users, err error) {
	if err = Db.Find(&users).Error; err != nil {
		return
	}
	return
}

func (user *Users) Update(id int64) (updateUser Users, err error) {

	if err = Db.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = Db.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}
	return
}

func (user *Users) Destroy(id int64) (Result Users, err error) {

	if err = Db.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = Db.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

