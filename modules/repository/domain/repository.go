package domain

import "time"

type Repository struct {
	Id               string `gorm:"primaryKey" json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	Type             string `json:"type,omitempty"`
	Url              string `json:"url,omitempty"`
	SSHKey           string `json:"ssh_key,omitempty"`
	SSHKeyPassphrase string `json:"ssh_key_passphrase,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

func (r *Repository) TableName() string {
	return "repository"
}
