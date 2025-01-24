package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID                     int                    `json:"id" gorm:"primaryKe;autoIncrement"`
	PhotoPath              string                 `json:"photo_path" gorm:"type:varchar(255)"`
	Username               string                 `json:"username" gorm:"type:varchar(50);uniqueIndex" validate:"required"`
	Password               string                 `json:"password,omitempty" gorm:"type:varchar(255)" validate:"required"`
	Email                  string                 `json:"email" gorm:"type:varchar(50);uniqueIndex" validate:"required,email"`
	Role                   string                 `json:"role" gorm:"type:user_role;default:guest"`
	FullName               string                 `json:"full_name" type:"varchar(50)"`
	Phone                  string                 `json:"phone" gorm:"type:varchar(20)"`
	Address                string                 `json:"address" gorm:"type:text"`
	EmailVerificationToken EmailVerificationToken `json:"-" gorm:"foreignKey:UserID;references:ID"`
	IsVerified             bool                   `json:"is_verified" gorm:"default:false"`
	CreatedAt              time.Time              `json:"-" gorm:"autoCreateTime"`
	UpdateAt               time.Time              `json:"-" gorm:"autoCreateTime,autoUpdateTime"`
}

func (*User) TableName() string {
	return "users"
}

func (u *User) Validate() error {
	v := validator.New()
	return v.Struct(u)
}

type EmailVerificationToken struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"user_id" gorm:"type:int"`
	Token     string    `json:"token,omitempty" gorm:"type:varchar(255)"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
}

func (*EmailVerificationToken) TableName() string {
	return "email_verification_tokens"
}

type UserSession struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       int       `json:"user_id" gorm:"type:int"`
	Token        string    `json:"token" gorm:"type:varchar(255)"`
	RefreshToken string    `json:"refresh_token" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime"`
	UpdateAt     time.Time `json:"-" gorm:"autoCreateTime;autoUpdateTime"`
}

func (*UserSession) TableName() string {
	return "user_sessions"
}
