package models

import "time"

type User struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Username    string `gorm:"varchar(300)" json:"username"`
	Email string `gorm:"varchar(300)" json:"email"`
	Password    string `gorm:"varchar(300)" json:"password"`
	CreatedAt time.Time `gorm:"type:time" json:"createdat"`
 	UpdatedAt time.Time `gorm:"type:time" json:"updatedat"`
}
