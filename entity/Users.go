package entity

type Users struct {
	Id         int    `gorm:"column:id"json:"id"`
	Username   string `gorm:"column:user_name"json:"username"`
	Password   string `gorm:"column:password"json:"password"`
}
