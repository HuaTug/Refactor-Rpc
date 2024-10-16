package balance

import (
	"sync"
)

// FullRoundBalance 实现 LoadBalance 接口
type FullRoundBalance struct {
	index int
	mu    sync.Mutex // 用于同步访问
}

// NewFullRoundBalance 创建 FullRoundBalance 实例
func NewFullRoundBalance() *FullRoundBalance {
	return &FullRoundBalance{}
}

// ChooseOne 选择一个服务实例
func (f *FullRoundBalance) ChooseOne(services []ServiceInfo) *ServiceInfo {
	f.mu.Lock() // 加锁以防止并发访问
	defer f.mu.Unlock()

	if len(services) == 0 {
		return nil // 如果没有服务实例，返回 nil
	}

	if f.index >= len(services) {
		f.index = 0 // 重置索引
	}

	service := services[f.index]
	f.index++ // 更新索引

	return &service // 返回选择的服务实例
}
