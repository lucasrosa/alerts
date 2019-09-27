package alerts

type port struct {
	repo AlertsSecondaryPort
}

// NewAlertsService receives a Secondary Port of domain and insantiates a Primary Port
func NewAlertsService(repo AlertsSecondaryPort) AlertsPrimaryPort {
	return &port{
		repo,
	}
}

func (p *port) GetActive() ([]Alert, error) {
	alerts, err := p.repo.GetActive()
	return alerts, err
}
