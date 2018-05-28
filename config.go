package client

// ConfigurationInfo - Configuration info of a client
type ConfigurationInfo interface {
	GetServiceID() string
	GetServiceName() string
	GetAddress() string
	GetPort() int
	GetHealthCheckURL() string
	GetInterval() string
	GetTimeOut() string
}

// Config - Consul client configuration
type Config struct {
	ServiceID      string
	ServiceName    string
	Address        string
	Port           int
	HealthCheckURL string
	Interval       string
	Timeout        string
}

// NewConfiguration - Instance new configuration
func NewConfiguration(serviceID string, serviceName string, address string, port int,
	healthcheckURL, interval, timeout string) *Config {
	return &Config{
		ServiceID:      serviceID,
		ServiceName:    serviceName,
		Address:        address,
		Port:           port,
		HealthCheckURL: healthcheckURL,
		Interval:       interval,
		Timeout:        timeout,
	}
}

// GetServiceID - get service ID information
func (c *Config) GetServiceID() string { return c.ServiceID }

// GetServiceName - get service name information
func (c *Config) GetServiceName() string { return c.ServiceName }

// GetAddress - get service address information
func (c *Config) GetAddress() string { return c.Address }

// GetPort - get service port information
func (c *Config) GetPort() int { return c.Port }

// GetHealthCheckURL - get service healthcheck URL information
func (c *Config) GetHealthCheckURL() string { return c.HealthCheckURL }

// GetInterval - get service interval information
func (c *Config) GetInterval() string { return c.Interval }

// GetTimeOut - get service timeout information
func (c *Config) GetTimeOut() string { return c.Timeout }
