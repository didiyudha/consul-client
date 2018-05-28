package client

// Registrator - service registration abstraction
type Registrator interface {
	Register() error
}

type registrator struct {
	ConsulClient *ConsulClient
}

// NewRegistrator - instance new registrator
func NewRegistrator(conf ConfigurationInfo) (Registrator, error) {
	consulClient, err := NewConsulClient(conf)
	if err != nil {
		return nil, err
	}
	r := &registrator{
		ConsulClient: consulClient,
	}
	return r, nil
}

// Register service to consul
func (r *registrator) Register() error {
	return r.ConsulClient.Agent.ServiceRegister(r.ConsulClient.Registration)
}
