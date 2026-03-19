package models

import (
	"time"
)

type Song struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SingerId    uint      `json:"singer_id"`
	Name        string    `gorm:"size:45" json:"name"`
	Introduction string  `gorm:"size:255" json:"introduction"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime  time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"update_time"`
	Pic         string    `gorm:"size:255" json:"pic"`
	Lyric       string    `gorm:"type:text" json:"lyric"`
	Url         string    `gorm:"size:255" json:"url"`
}

func (Song) TableName() string {
	return "song"
}
