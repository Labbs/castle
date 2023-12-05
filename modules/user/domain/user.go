package domain

import "time"

type User struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"primaryKey" json:"email"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty" gorm:"default:''"`
	DarkMode string `json:"dark_mode" gorm:"default:'auto'"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}

func (User) TableName() string {
	return "user"
}

type PasswordChangeRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type DarkModeChangeRequest struct {
	DarkMode string `json:"dark_mode"`
}

type AvatarChangeRequest struct {
	Avatar string `json:"avatar"`
}
