package models

type User struct {
	UserId          int64
	UserName        string `gorm:"type:varchar(32);unique"`
	Password        string `gorm:"type:varchar(32);not null"`
	Token           string `gorm:"type:varchar(32);not null"`
	FollowCount     uint   `gorm:"default:0;"`
	FollowerCount   uint   `gorm:"default:0;"`
	IsFollow        bool   `gorm:"default:false;"`
	Avatar          string `gorm:"type:varchar(255);"`
	BackGroundImage string `gorm:"type:varchar(255);"`
	Signature       string `gorm:"type:varchar(255);"`
	TotalFavorited  string `gorm:"type:varchar(255);"`
	WorkCount       uint   `gorm:"default:0;"`
	FavoriteCount   uint   `gorm:"default:0;"`
}

type RegisterForm struct {
	UserName string `gorm:"varchar(32);unique;not null" json:"username" binding:"required"`
	Password string `gorm:"varchar(32);not null" json:"password" binding:"required"`
}

type LoginForm struct {
	UserName string `gorm:"varchar(32);unique;not null" json:"username" binding:"required"`
	Password string `gorm:"varchar(32);not null" json:"password" binding:"required"`
}
