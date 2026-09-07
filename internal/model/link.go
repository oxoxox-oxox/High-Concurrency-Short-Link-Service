package model

import (
	"shortlink/pkg/base62"
	"time"
)

type Link struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	ShortCode string `gorm:"type:varchar(10);uniqueIndex"`
	LongURL   string `gorm:"type:varchar(2048);not null"`
	CreatedAt time.Time
}

func GenerateShortLink(longURL string) (string, error) {
	link := Link{LongURL: longURL}

	if err := DB.Create(&link).Error; err != nil {
		return "", err
	}

	shortCode := base62.Encode(link.ID)

	if err := DB.Model(&link).Update("short_code", shortCode).Error; err != nil {
		return "", err
	}

	return shortCode, nil
}

func GetLongURL(shortCode string) (string, error) {
	var link Link

	if err := DB.Where("short_code = ?", shortCode).First(&link).Error; err != nil {
		return "", err
	}
	return link.LongURL, nil
}
