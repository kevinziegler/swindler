package connections

type ConnectionProfile struct {
	Name string
	Adapter string
	Configuration map[string]interface{}
}

type Repository interface {
	Create(*ConnectionProfile)
}
