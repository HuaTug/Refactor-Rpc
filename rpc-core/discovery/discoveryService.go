package discovery

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"HuaTug.com/rpc-core/balance" // Adjust this import according to your project structure
	clientv3 "go.etcd.io/etcd/client/v3"
)

// DiscoveryService interface definition
type DiscoveryService interface {
	Discovery(serviceName string) (*balance.ServiceInfo, error)
}

// EtcdDiscoveryService implements DiscoveryService interface
type EtcdDiscoveryService struct {
	cli         *clientv3.Client
	loadBalance balance.LoadBalance // Assuming LoadBalance is a defined interface
}

// NewEtcdDiscoveryService creates an instance of EtcdDiscoveryService
func NewEtcdDiscoveryService(endpoints []string, loadBalance balance.LoadBalance) (*EtcdDiscoveryService, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Failed to connect to Etcd: %v", err)
		return nil, err
	}

	return &EtcdDiscoveryService{cli: cli, loadBalance: loadBalance}, nil
}

// Discovery method for service discovery
func (e *EtcdDiscoveryService) Discovery(serviceName string) (*balance.ServiceInfo, error) {
	key := fmt.Sprintf("/demo_rpc/%s", serviceName)

	// Get all service instances for the given service name
	resp, err := e.cli.Get(context.Background(), key)
	if err != nil {
		return nil, fmt.Errorf("failed to query service instances: %w", err)
	}

	var serviceInstances []balance.ServiceInfo

	for _, kv := range resp.Kvs {
		valueStr := strings.TrimSpace(string(kv.Value))
		log.Println("Raw value:", valueStr)

		if len(valueStr) == 0 {
			log.Println("Value is empty")
			continue
		}

		// 使用 strings.Split 分割字符串
		parts := strings.Split(valueStr, ":")
		if len(parts) != 2 {
			log.Printf("Unexpected format for service instance data: %s", valueStr)
			continue
		}

		address := parts[0]
		portStr := parts[1]

		// 解析端口为整数
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.Printf("Failed to parse port: %v", err)
			continue
		}

		serviceInstances = append(serviceInstances, balance.ServiceInfo{
			ServiceName: serviceName,
			Address:     address,
			Port:        port,
		})
	}

	if len(serviceInstances) == 0 {
		return nil, fmt.Errorf("no instances found for service: %s", serviceName)
	}

	// Use load balancing to choose one service instance
	return e.loadBalance.ChooseOne(serviceInstances), nil
}
