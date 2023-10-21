package domain

import "time"

type User struct {
	Username string `gorm:"primaryKey" json:"username"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty" gorm:"default:''"`
	DarkMode string `json:"dark_mode" gorm:"default:'auto'"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}

type UsernameChangeRequest struct {
	CurrentUsername string `json:"current_username"`
	NewUsername     string `json:"new_username"`
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
