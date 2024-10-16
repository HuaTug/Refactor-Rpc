package common

import "fmt"

// ServiceUtil 提供服务相关的工具方法
type ServiceUtil struct{}

// ServiceKey 构造服务的唯一标识
func ServiceKey(serviceName string, version string) string {
	return fmt.Sprintf("%s-%s", serviceName, version)
}
