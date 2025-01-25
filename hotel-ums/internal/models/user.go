package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// User Model
type User struct {
	ID                     int                    `json:"id" gorm:"primaryKe;autoIncrement"`
	PhotoPath              string                 `json:"photo_path" gorm:"type:varchar(255)"`
	Username               string                 `json:"username" gorm:"type:varchar(50);uniqueIndex" validate:"required" form:"username"`
	Password               string                 `json:"password,omitempty" gorm:"type:varchar(255)" validate:"required"`
	Email                  string                 `json:"email" gorm:"type:varchar(50);uniqueIndex" validate:"required,email" form:"email"`
	Role                   string                 `json:"role" gorm:"type:user_role;default:guest"`
	FullName               string                 `json:"full_name" type:"varchar(50)" form:"full_name"`
	Phone                  string                 `json:"phone" gorm:"type:varchar(20)" form:"phone"`
	Address                string                 `json:"address" gorm:"type:text" form:"address"`
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

// EmailVerificationToken Model
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

// UserSession Model
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

// Request
type ResendEmailVerificationRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (I *ResendEmailVerificationRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (I *LoginRequest) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

// Response
type LoginResponse struct {
	UserID       int    `json:"user_id"`
	FullName     string `json:"full_name"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}
