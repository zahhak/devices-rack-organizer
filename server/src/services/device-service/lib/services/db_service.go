package services

import "github.com/zahhak/devices-rack-organizer/server/src/services/device-service/lib/types/db"

// DBService ...
type DBService struct {
}

// GetAllDevices returns all devices stored in the database.
func (service *DBService) GetAllDevices() ([]db.Device, error) {
	return []db.Device{}, nil
}
