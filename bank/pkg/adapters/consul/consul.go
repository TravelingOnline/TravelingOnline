package consul

import (
	"fmt"

	consulAPI "github.com/hashicorp/consul/api"
)

type Consul struct {
	Address string
}

func NewConsul(address string) *Consul {
	return &Consul{Address: address}
}

func (c *Consul) RegisterService(serviceName, serviceHostAddress, servicePrefixPath, serviceHTTPHealthPath string, serviceGRPCPort, serviceHTTPPort int) error {
	consulConfig := consulAPI.DefaultConfig()
	consulConfig.Address = c.Address
	consulClient, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		return err
	}

	// HTTP health check URL
	HTTPHealthURL := fmt.Sprintf("http://%s:%d/%s", serviceHostAddress, serviceHTTPPort, serviceHTTPHealthPath)

	// Register HTTP service
	httpServiceRegistration := &consulAPI.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-http", serviceName),
		Name:    fmt.Sprintf("%s-http", serviceName),
		Address: serviceHostAddress,
		Port:    serviceHTTPPort,
		Tags:    []string{"http"},
		Checks: []*consulAPI.AgentServiceCheck{
			{
				HTTP:     HTTPHealthURL,
				Interval: "10s",
				Timeout:  "1s",
			},
		},
	}

	// Register gRPC service
	grpcServiceRegistration := &consulAPI.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-grpc", serviceName),
		Name:    fmt.Sprintf("%s-grpc", serviceName),
		Address: serviceHostAddress,
		Port:    serviceGRPCPort,
		Tags:    []string{"grpc"},
		Checks: []*consulAPI.AgentServiceCheck{
			{
				GRPC:     fmt.Sprintf("%s:%d", serviceHostAddress, serviceGRPCPort),
				Interval: "10s",
				Timeout:  "1s",
			},
		},
	}

	// Register both services
	if err := consulClient.Agent().ServiceRegister(httpServiceRegistration); err != nil {
		return fmt.Errorf("failed to register HTTP service: %w", err)
	}

	if err := consulClient.Agent().ServiceRegister(grpcServiceRegistration); err != nil {
		return fmt.Errorf("failed to register gRPC service: %w", err)
	}

	return nil
}

func (c *Consul) DiscoverService(serviceName string) (port int, ip string, err error) {
	consulConfig := consulAPI.DefaultConfig()
	consulConfig.Address = c.Address
	consulClient, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		return 0, "", err
	}
	services, _, err := consulClient.Catalog().Service(serviceName, "", nil)
	if err != nil {
		return 0, "", err
	}
	if len(services) == 0 {
		return 0, "", fmt.Errorf("service %s not found", serviceName)
	}
	service := services[0]
	return service.ServicePort, service.ServiceAddress, nil
}
