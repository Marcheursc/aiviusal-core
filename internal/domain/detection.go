package domain

import "time"

// ObjectDetection 目标检测信息
type ObjectDetection struct {
	ID           string    `json:"id"`
	TimeStamp    time.Time `json:"time_stamp"`
	DeviceID     string    `json:"device_id"`
	ObjectType   string    `json:"object_type"` // PERSON, VEHICLE, ANIMAL
	Confidence   float64   `json:"confidence"`
	BoundingBox  Box       `json:"bounding_box"`
	TrackingID   string    `json:"tracking_id,omitempty"`
	SnapshotURL  string    `json:"snapshot_url,omitempty"`
}

// Box 边界框
type Box struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}