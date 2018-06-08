package types

import "github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types/db"

// User ...
type User struct {
	UserID string
}

// DeviceExt ...
type DeviceExt struct {
	db.Device
	Image string
}
