package registry

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// ServiceInfo 定义服务信息结构体
type ServiceInfo struct {
	ServiceName string
	Address     string
	Port        int
}

// RegistryService 接口定义
type RegistryService interface {
	Register(serviceInfo ServiceInfo) error
	Unregister(serviceInfo ServiceInfo) error
	Destroy() error
}

// EtcdRegistryService 实现 RegistryService 接口
type EtcdRegistryService struct {
	cli        *clientv3.Client
	leaseID    clientv3.LeaseID
	leaseGrant clientv3.LeaseID
}

// NewEtcdRegistryService 创建 EtcdRegistryService 实例
func NewEtcdRegistryService(endpoints []string) (*EtcdRegistryService, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Failed to connect to Etcd: %v", err)
		return nil, err
	}

	return &EtcdRegistryService{cli: cli}, nil
}

// Register 注册服务
func (e *EtcdRegistryService) Register(serviceInfo ServiceInfo) error {
	key := fmt.Sprintf("/demo_rpc/%s", serviceInfo.ServiceName)
	value := fmt.Sprintf("%s:%d", serviceInfo.Address, serviceInfo.Port)
	log.Println(key, ",", value)
	leaseResp, err := e.cli.Grant(context.Background(), 5) // 租约 TTL 设置为 5 秒
	if err != nil {
		return fmt.Errorf("failed to create lease: %v", err)
	}
	_, err = e.cli.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return fmt.Errorf("failed to register service: %v", err)
	}

	e.leaseID = leaseResp.ID // 保存租约 ID
	return nil
}

// Unregister 注销服务
func (e *EtcdRegistryService) Unregister(serviceInfo ServiceInfo) error {
	key := fmt.Sprintf("/demo_rpc/%s", serviceInfo.ServiceName)

	_, err := e.cli.Delete(context.Background(), key)
	return err
}

// Destroy 关闭连接（如果需要）
func (e *EtcdRegistryService) Destroy() error {
	return e.cli.Close() // 显式关闭 Etcd 客户端连接
}
