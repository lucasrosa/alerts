package alerts

// Alerts represents
type Alert struct {
	AlertID string `dynamo:"alertid" json:"id"`
	Start   int    `dynamo:"start" json:"start"`
	End     int    `dynamo:"end" json:"end"`
}
