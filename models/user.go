package models

type User struct {
	ID       uint   `gorm:"primary_key"`
	Email    string `gorm:"type:varchar(100);uniqueIndex"`
	Name     string `gorm:"type:varchar(100)"`
	Password string `json:"Password"`
	Role     string `gorm:"type:varchar(10)"`
	Verify   string `gorm:"type:varchar(10)"`
}
