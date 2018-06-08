package types

import "time"

// User ...
type User struct {
	UserID string
}

// DeviceExt ...
type DeviceExt struct {
	DeviceID string    `json:"deviceId"`
	UserID   string    `json:"userId"`
	Date     time.Time `json:"date"`
	ImageURL string    `json:"imageUrl"`
}
