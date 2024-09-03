package register

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type consul struct {
	client *api.Client
}

var Reg Register

// 确保consul结构体实现了Register接口
var _ Register = (*consul)(nil)

func Init(addr string) error {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	c, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	Reg = &consul{c}
	return nil
}

// RegisterServer 将grpc服务注册到consul
func (c *consul) RegisterServer(serviceName string, ip string, port int, tags []string) error {
	// init check
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", ip, port),
		Timeout:                        "10s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "20s",
	}

	// register
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port),
		Name:    serviceName,
		Tags:    tags,
		Address: ip,
		Port:    port,
		Check:   check,
	}
	return c.client.Agent().ServiceRegister(srv)
}

// ListService 服务发现
func (c *consul) ListService(serviceName string) (map[string]*api.AgentService, error) {
	return c.client.Agent().ServicesWithFilter(fmt.Sprintf("Service=='%s'", serviceName))
}

// Deregister 注销服务
func (c *consul) Deregister(serviceID string) error {
	return c.client.Agent().ServiceDeregister(serviceID)
}
