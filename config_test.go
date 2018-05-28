package client

import "testing"

func TestGetServiceID(t *testing.T) {
	serviceID := "payment-service"
	c := &Config{
		ServiceID: serviceID,
	}
	srvcID := c.GetServiceID()
	if srvcID != serviceID {
		t.Errorf("\texpected service ID %s, actual %s\n", serviceID, srvcID)
	}
}

func TestGetServiceName(t *testing.T) {
	serviceName := "transaction-service"
	c := &Config{
		ServiceName: serviceName,
	}
	srvcName := c.GetServiceName()
	if srvcName != serviceName {
		t.Errorf("\texpected service name %s, actual %s\n", serviceName, srvcName)
	}
}

func TestGetAddress(t *testing.T) {
	address := "127.0.0.1"
	c := &Config{
		Address: address,
	}
	addrs := c.GetAddress()
	if address != addrs {
		t.Errorf("\texpected service address %s, actual %s\n", address, addrs)
	}
}

func TestGetPort(t *testing.T) {
	port := 8080
	c := &Config{
		Port: port,
	}
	p := c.GetPort()
	if port != p {
		t.Errorf("\texpected port %d, but actual %d\n", port, p)
	}
}

func TestGetHealthCheckURL(t *testing.T) {
	healthcheckURL := "http://localhost:8080/healthcheck"
	c := &Config{
		HealthCheckURL: healthcheckURL,
	}
	hlthcheckURL := c.GetHealthCheckURL()
	if hlthcheckURL != healthcheckURL {
		t.Errorf("\texpected healthcheck URL %s, but actual %s\n", healthcheckURL, hlthcheckURL)
	}
}

func TestGetInterval(t *testing.T) {
	interval := "3s"
	c := &Config{
		Interval: interval,
	}
	intrvl := c.GetInterval()
	if intrvl != interval {
		t.Errorf("\texpected interval %s, but actual %s\n", interval, intrvl)
	}
}

func TestGetTimeOut(t *testing.T) {
	timeout := "5s"
	c := &Config{
		Timeout: timeout,
	}
	to := c.GetTimeOut()
	if to != timeout {
		t.Errorf("\texpected timeout %s, but actual %s\n", timeout, to)
	}
}

func TestNewConfiguration(t *testing.T) {
	serviceID := "service-001"
	serviceName := "payment-service"
	address := "localhost"
	port := 8080
	healthcheck := "localhost:8080/healthcheck"
	interval := "3s"
	timeout := "5s"
	c := NewConfiguration(serviceID, serviceName, address, port, healthcheck, interval, timeout)
	sid := c.GetServiceID()
	if sid != serviceID {
		t.Errorf("\texpected service ID %s, actual %s\n", serviceID, sid)
	}
	srvcName := c.GetServiceName()
	if serviceName != srvcName {
		t.Errorf("\texpected service name %s, actual %s\n", serviceName, srvcName)
	}
	addr := c.GetAddress()
	if addr != address {
		t.Errorf("\texpected service address %s, actual %s\n", address, addr)
	}
	p := c.GetPort()
	if port != p {
		t.Errorf("\texpected port %d, but actual %d\n", port, p)
	}
	healthcheckURL := c.GetHealthCheckURL()
	if healthcheck != healthcheckURL {
		t.Errorf("\texpected healthcheck URL %s, but actual %s\n", healthcheck, healthcheckURL)
	}
	intrvl := c.GetInterval()
	if interval != intrvl {
		t.Errorf("\texpected interval %s, but actual %s\n", interval, intrvl)
	}
	to := c.GetTimeOut()
	if timeout != to {
		t.Errorf("\texpected timeout %s, but actual %s\n", timeout, to)
	}
}
