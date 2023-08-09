package model

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string
}

type User struct {
	ID       int64  `gorm:"primary_key" json:"-"`
	Username string `json:"name"`
	Password string `json:"password"`
	Email    string
}
