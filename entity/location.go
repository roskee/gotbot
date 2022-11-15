package entity

// Location represents a point on the map.
type Location struct {
	// Longitude is longitude as defined by sender.
	//
	// It is a required field
	Longitude float64 `json:"longitude,omitempty"`
	// Latitude is latitude as defined by sender.
	//
	// It is a required field
	Latitude float64 `json:"latitude,omitempty"`
	// HorizontalAccuracy is the radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// LivePeriod is time relative to the message sending date,
	// during which the location can be updated; in seconds.
	// For active live locations only.
	LivePeriod int64 `json:"live_period,omitempty"`
	// Heading is  The direction in which user is moving, in degrees; 1-360.
	// For active live locations only.
	Heading int64 `json:"heading,omitempty"`
	// ProximityAlertRadius is the maximum distance for proximity alerts
	// about approaching another chat member, in meters.
	// For sent live locations only.
	ProximityAlertRadius int64 `json:"proximity_alert_radius,omitempty"`
}
