package types

import (
	"time"
)

// User ...
type User struct {
	UserID string
}

// DeviceExt ...
type DeviceExt struct {
	DeviceID   string    `json:"deviceId"`
	UserID     string    `json:"userId"`
	Date       time.Time `json:"date"`
	ImageURL   string    `json:"imageUrl"`
	HistoryURL string    `json:"historyUrl"`
}

// DeviceHistoryRecord ...
type DeviceHistoryRecord struct {
	OldOwner string    `json:"oldOwner"`
	NewOwner string    `json:"newOwner"`
	Date     time.Time `json:"date"`
}

// DeviceState ...
type DeviceState struct {
	DeviceState string `json:"deviceState"`
}
