package alerts

// AlertsPrimaryPort is the entrypoint for the alerts Package
type AlertsPrimaryPort interface {
	GetActive() ([]Alert, error)
}

// AlertsSecondaryPort is the way the business rules communicate to the external world
type AlertsSecondaryPort interface {
	GetActive() ([]Alert, error)
}
