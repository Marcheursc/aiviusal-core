package domain

// DeviceInfo 设备信息
type DeviceInfo struct {
	DeviceID   string `json:"device_id"`
	DeviceName string `json:"device_name"`
	IPAddress  string `json:"ip_address"`
	Port       int    `json:"port"`
}