package client

import (
	consulapi "github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	Conf         *consulapi.Config
	Client       *consulapi.Client
	Agent        *consulapi.Agent
	Registration *consulapi.AgentServiceRegistration
	ServiceCheck *consulapi.AgentServiceCheck
}

// SetDefaultConfig - set default config of Consul
func (cc *ConsulClient) SetDefaultConfig(conf *consulapi.Config) {
	cc.Conf = conf
}

// SetConsulClient - set consul client
func (cc ConsulClient) SetConsulClient(client *consulapi.Client) {
	cc.Client = client
}

// SetAgent - set consul agent
func (cc *ConsulClient) SetAgent(agent *consulapi.Agent) {
	cc.Agent = agent
}

// SetRegistration - set consul registration
func (cc *ConsulClient) SetRegistration(registration *consulapi.AgentServiceRegistration) {
	cc.Registration = registration
}

// NewDefaultConfig - instance new default config
func NewDefaultConfig() *consulapi.Config {
	return consulapi.DefaultConfig()
}

// NewClient - instance new client
func NewClient(config *consulapi.Config) (*consulapi.Client, error) {
	return consulapi.NewClient(config)
}

// NewRegitration - instance new consul registration
func NewRegitration() *consulapi.AgentServiceRegistration {
	return new(consulapi.AgentServiceRegistration)
}

// NewServiceCheck - instance new consul agent service check
func NewServiceCheck() *consulapi.AgentServiceCheck {
	return new(consulapi.AgentServiceCheck)
}

// NewConsulClient - instance new consul client
func NewConsulClient(configurationInfo ConfigurationInfo) (*ConsulClient, error) {
	config := NewDefaultConfig()
	client, err := NewClient(config)
	serviceCheck := NewServiceCheck()
	if err != nil {
		return nil, err
	}
	registration := NewRegitration()
	registration.ID = configurationInfo.GetServiceID()
	registration.Name = configurationInfo.GetServiceName()
	registration.Address = configurationInfo.GetAddress()
	registration.Port = configurationInfo.GetPort()
	registration.Check = serviceCheck
	registration.Check.HTTP = configurationInfo.GetHealthCheckURL()
	registration.Check.Interval = configurationInfo.GetInterval()
	registration.Check.Timeout = configurationInfo.GetTimeOut()
	consulClient := &ConsulClient{
		Conf:         config,
		Client:       client,
		Registration: registration,
		Agent:        client.Agent(),
	}
	return consulClient, nil
}
