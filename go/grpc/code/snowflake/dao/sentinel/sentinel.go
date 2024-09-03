package sentinel

import (
	"github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"log"
	"snowflake/dao/config"
)

func initSentinel() {
	err := api.InitDefault() // 使用默认配置初始化 Sentinel
	if err != nil {
		log.Fatalf("Sentinel initialization failed: %v", err)
	}

	// 配置 Sentinel 流量控制规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:         config.Conf.Name, // 使用配置中的服务名作为资源名
			ControlBehavior:  flow.Reject,      // 当 QPS 超限时，直接拒绝请求
			StatIntervalInMs: 1000,             // 统计间隔为 1000ms
		},
	})
	if err != nil {
		log.Fatalf("Failed to load Sentinel rules: %v", err)
	}
}
