package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `binding:"required"`
	Email     string `binding:"required"`
	Password  string
}
type Place struct {
	ID          uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `binding:"required"`
	Address     string `binding:"required"`
	ZipCode     string `binding:"required"`
	City      	string `binding:"required"`
	Code        string
	Mail        string
	Latitude    float32 `binding:"required"`
	Longitude   float32 `binding:"required"`
	MaxDistance float32 `binding:"required"`
}
type Picture struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `binding:"required"`
	MimeType  string `binding:"required"`
	Data      []byte `binding:"required"`
}
type Campaign struct {
	ID          uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StartDate   *time.Time `binding:"required"`
	EndDate     *time.Time `binding:"required"`
	Name        string     `binding:"required"`
	Description string
	PictureID   uint64
}
type Gift struct {
	ID          uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `binding:"required"`
	CampaignID  uint64 `binding:"required"`
	PictureID   uint64
	Description string
}
type InstantWinnerGame struct {
	ID          uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string     `binding:"required"`
	PlayTime    *time.Time `binding:"required" json:"-"` //in seconds, - will hide PlayTime in response
	StartDate   *time.Time `binding:"required"`
	EndDate     *time.Time `binding:"required"`
	GiftID      uint64     `binding:"required"`
	CampaignID  uint64     `binding:"required"`
	PlaceID     uint64     `binding:"required"`
	Place       Place
	Description string
	Won         bool
}

type RandomDrawGame struct {
	ID          uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `binding:"required"`
	Description string
	GiftID      uint64     `binding:"required"`
	StartDate   *time.Time `binding:"required"`
	EndDate     *time.Time `binding:"required"`
}
type InstantWinnerPlayer struct {
	ID                  uint64
	CreatedAt           time.Time
	UpdatedAt           time.Time
	InstantWinnerGameID uint64 `binding:"required"`
	InstantWinnerGame   *InstantWinnerGame
	IPAddress           string
	Fingerprint         string `binding:"required"`
	Email               string
	Time                *time.Time `binding:"required"`
	Result              bool
}
type RandomDrawPlayer struct {
	ID               uint64
	RandomDrawGameID uint64 `binding:"required"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	IPAddress        string
	Email            string     `binding:"required"`
	Time             *time.Time `binding:"required"`
	Won              bool
}

func InitModels(router *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Place{})
	db.AutoMigrate(&Picture{})
	db.AutoMigrate(&Campaign{})
	db.AutoMigrate(&Gift{})
	db.AutoMigrate(&InstantWinnerGame{})
	db.AutoMigrate(&RandomDrawGame{})
	db.AutoMigrate(&InstantWinnerPlayer{})
	db.AutoMigrate(&RandomDrawPlayer{})
}
