package balance

// ServiceInfo 定义服务信息结构体
type ServiceInfo struct {
	ServiceName string
	Address     string
	Port        int
}

// LoadBalance 定义负载均衡接口
type LoadBalance interface {
	// ChooseOne 从服务列表中选择一个服务实例
	ChooseOne(services []ServiceInfo) *ServiceInfo
}
