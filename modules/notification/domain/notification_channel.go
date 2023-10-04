package domain

import (
	"encoding/json"
	"time"
)

type NotificationChannel struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	ChannelType     string          `json:"channel_type"`
	SlackChannel    SlackChannel    `json:"slack_channel,omitempty"`
	EmailChannel    EmailChannel    `json:"email_channel,omitempty"`
	DiscordChannel  DiscordChannel  `json:"discord_channel,omitempty"`
	TelegramChannel TelegramChannel `json:"telegram_channel,omitempty"`

	Enabled bool `json:"enabled"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"-"`
}

type SlackChannel struct {
	WebhookUrl string `json:"webhook_url"`
}

type EmailChannel struct {
	Email string `json:"email"`
}

type DiscordChannel struct {
	WebhookUrl string `json:"webhook_url"`
}

type TelegramChannel struct {
	WebhookUrl string `json:"webhook_url"`
}

func (sla *SlackChannel) Scan(value interface{}) error {
	var skills SlackChannel
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla *EmailChannel) Scan(value interface{}) error {
	var skills EmailChannel
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla *DiscordChannel) Scan(value interface{}) error {
	var skills DiscordChannel
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla *TelegramChannel) Scan(value interface{}) error {
	var skills TelegramChannel
	err := json.Unmarshal([]byte(value.(string)), &skills)
	if err != nil {
		return err
	}
	*sla = skills
	return nil
}

func (sla SlackChannel) Value() (string, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

func (sla EmailChannel) Value() (string, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

func (sla DiscordChannel) Value() (string, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

func (sla TelegramChannel) Value() (string, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}
