package models

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type RoomType struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"type:varchar(50);unique;not null" json:"name" validate:"required"`
	Description   string    `gorm:"type:text;not null" json:"description" validate:"required"`
	PricePerNight float64   `gorm:"type:decimal(10,2);not null" json:"price_per_night" validate:"required"`
	Capacity      int       `gorm:"type:int;not null" json:"capacity" validate:"required"`
	TotalRooms    int       `gorm:"type:int;not null" json:"total_rooms" validate:"required"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"-"`

	// Relationships
	RoomFeatures []RoomFeature `gorm:"foreignKey:RoomTypeID" json:"room_features"`
	RoomPhotos   []RoomPhoto   `gorm:"foreignKey:RoomTypeID" json:"room_photos"`
}

func (I *RoomType) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type RoomFeature struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomTypeID uint      `gorm:"type:int;not null" json:"room_type_id" validate:"required"`
	Features   string    `gorm:"type:varchar(50);not null" json:"features" validate:"required"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"-"`

	// Relationship
	RoomType RoomType `gorm:"foreignKey:RoomTypeID" json:"room_type"`
}

func (I *RoomFeature) Validate() error {
	v := validator.New()
	return v.Struct(I)
}

type RoomPhoto struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomTypeID uint      `gorm:"type:int;not null" json:"room_type_id"`
	FilePath   string    `gorm:"type:varchar(255);not null" json:"file_path"`
	IsPrimary  bool      `gorm:"type:boolean;not null;default:false" json:"is_primary"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"-"`

	// Relationship
	RoomType RoomType `gorm:"foreignKey:RoomTypeID" json:"room_type"`
}

type Room struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomNumber  string    `gorm:"type:varchar(10);unique;not null" json:"room_number" validate:"required"`
	RoomTypeID  uint      `gorm:"type:int;not null" json:"room_type_id" validate:"required"`
	Description string    `gorm:"type:text" json:"description" validate:"required"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"-"`

	// Relationship
	RoomType RoomType `gorm:"foreignKey:RoomTypeID" json:"room_type"`
}

func (I *Room) Validate() error {
	v := validator.New()
	return v.Struct(I)
}
