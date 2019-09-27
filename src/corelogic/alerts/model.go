package alerts

// Alerts represents
type Alert struct {
	ID    string `json:"id"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}
