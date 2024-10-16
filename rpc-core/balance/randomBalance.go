package balance

import (
	"math/rand"
	"sync"
	"time"
)

// RandomBalance 实现 LoadBalance 接口
type RandomBalance struct {
	mu sync.Mutex // 用于同步访问
}

// NewRandomBalance 创建 RandomBalance 实例
func NewRandomBalance() *RandomBalance {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	return &RandomBalance{}
}

// ChooseOne 随机选择一个服务实例
func (r *RandomBalance) ChooseOne(services []ServiceInfo) *ServiceInfo {
	r.mu.Lock() // 加锁以防止并发访问
	defer r.mu.Unlock()

	if len(services) == 0 {
		return nil // 如果没有服务实例，返回 nil
	}

	index := rand.Intn(len(services)) // 随机选择一个索引
	return &services[index]           // 返回选择的服务实例
}
