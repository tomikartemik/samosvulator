package model

type User struct {
	ID       int      `gorm:"primaryKey;uniqueIndex;autoIncrement;primaryKey" json:"id"`
	Username string   `gorm:"uniqueIndex;not null" json:"username"`
	Password string   `gorm:"not null" json:"password"`
	Name     string   `gorm:"not null" json:"name"`
	Surname  string   `gorm:"not null" json:"surname"`
	Company  string   `gorm:"not null" json:"company"`
	Section  string   `gorm:"not null" json:"section"`
	JobTitle string   `gorm:"not null" json:"job_title"`
	Records  []Record `gorm:"foreignKey:UserID" json:"records"`
}

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInOutput struct {
	Token string     `json:"token"`
	User  UserOutput `json:"user"`
}

type ChangePasswordInput struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_pass"`
}

type UserOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Company  string `json:"company"`
	Section  string `json:"section"`
	JobTitle string `json:"job_title"`
}
