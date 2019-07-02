package entity

import "time"

type Post struct {
	ID        int64 `gorm:"primary_key"`
	Body      string
	Reply     string
	CreatedAt time.Time
}
