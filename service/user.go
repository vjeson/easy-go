package service

import (
	. "demo/db"
)

type User struct {
	Id         int    `gorm:"column:id"json:"id"`
	Username   string `gorm:"column:username"json:"username"`
	Password   string `gorm:"column:password"json:"password"`
}

func (user *User) ChekcLogin() *User {
	u := new(User)
	Db.Where(user).Take(u)
	return u
}

func (user *User) Insert() (id int, err error) {
	result := Db.Create(&user)
	id = user.Id
	if result != nil {
		err = result.Error
		return
	}
	return
}

func (user *User) Users() (users []User, err error) {
	if err = Db.Find(&users).Error; err != nil {
		return
	}
	return
}

func (user *User) Update(id int64) (updateUser User, err error) {

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

func (user *User) Destroy(id int64) (Result User, err error) {

	if err = Db.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = Db.Delete(&user).Error; err != nil {
		return
	}
	Result = *user
	return
}

