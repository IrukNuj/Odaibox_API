package model

import "time"

type Post struct {
	ID        int64 `gorm:"primary_key"`
	Body      string
	Reply     string `v-post:"max=0"`
	CreatedAt time.Time
}
