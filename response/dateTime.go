package riminderResponse

// DateTime time representation as riminder response give.
type DateTime struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}
