package types

import (
	"time"
)

// User ...
type User struct {
	UserID string
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
