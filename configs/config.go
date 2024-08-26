package configs

type ServiceConfig struct {
	Services map[string]string
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		Services: allServices,
	}
}

func (sc *ServiceConfig) GetHost(serviceName string) string {
	return allServices[serviceName]
}

var allServices = map[string]string{
	"auth": "http://auth-app:8082",
	"cart": "http://cart-app:8083",
}
