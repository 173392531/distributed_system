package discovery

import (
	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	client *api.Client
}

func NewConsulClient(address string) *ConsulClient {
	config := api.DefaultConfig()
	config.Address = address
	client, _ := api.NewClient(config)
	return &ConsulClient{client}
}

func (c *ConsulClient) Register(serviceName, serviceID, serviceAddress string) error {
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: serviceAddress,
	}
	return c.client.Agent().ServiceRegister(registration)
}

func (c *ConsulClient) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}

func (c *ConsulClient) GetService(serviceName string) ([]*api.ServiceEntry, error) {
	services, _, err := c.client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}
	return services, nil
}
