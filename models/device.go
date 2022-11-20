package models

type Device struct {
	Id           uint64 `json:"id"`
	LocationId   uint64 `json:"locationId"`
	DeviceTypeId uint64 `json:"DeviceTypeId"`
	Name         string `json:"name"`
}
