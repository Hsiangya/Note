package register

import "github.com/hashicorp/consul/api"

type Register interface {
	// RegisterServer 注册
	RegisterServer(serviceName string, ip string, port int, tags []string) error

	// ListService 发现
	ListService(serviceName string) (map[string]*api.AgentService, error)

	// Deregister 注销
	Deregister(serviceID string) error
}
