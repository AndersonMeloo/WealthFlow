package models

import "time"

type ProcessedEvent struct {
	ID            uint      `gorm:"primaryKey"`
	EventID       string    `gorm:"column:event_id;uniqueIndex;not null"`
	CardID        string    `gorm:"column:card_id;not null"`
	CustomerEmail string    `gorm:"column:customer_email;not null"`
	Timestamp     time.Time `gorm:"column:timestamp;not null"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
}

func (ProcessedEvent) TableName() string {
	return "processed_events"
}
