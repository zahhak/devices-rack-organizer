package db

import (
	"time"
)

// Device ...
type Device struct {
	DeviceID string
	UserID   string
	Date     time.Time
}
