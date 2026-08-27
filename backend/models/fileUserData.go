package models

import (
	"time"
)

type UserInfo struct {
    Filename    string
    Size        int64
    Path        string // or blob_id
    CreatedAt   time.Time
}