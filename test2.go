package main

import (
	"log"
	"time"

	"HuaTug.com/rpc-core/balance"
	"HuaTug.com/rpc-core/discovery"
	registry "HuaTug.com/rpc-core/register"
)

// SimpleLoadBalance is a simple implementation of LoadBalance
type SimpleLoadBalance struct{}

func (lb *SimpleLoadBalance) ChooseOne(services []balance.ServiceInfo) *balance.ServiceInfo {
	if len(services) == 0 {
		return nil
	}
	return &services[0] // Simply return the first service
}

func main() {
	endpoints := []string{"localhost:2379"} // Etcd address
	loadBalance := &SimpleLoadBalance{}

	// Create registry service instance
	registryService, err := registry.NewEtcdRegistryService(endpoints)
	if err != nil {
		log.Fatalf("Failed to create EtcdRegistryService: %v", err)
	}

	// Create discovery service instance
	discoveryService, err := discovery.NewEtcdDiscoveryService(endpoints, loadBalance)
	if err != nil {
		log.Fatalf("Failed to create EtcdDiscoveryService: %v", err)
	}

	// Register service
	serviceInfo := registry.ServiceInfo{
		ServiceName: "TestService",
		Address:     "127.0.0.1",
		Port:        8080,
	}

	err = registryService.Register(serviceInfo)
	if err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	// Wait a moment to ensure the service is registered successfully
	time.Sleep(1 * time.Second)

	// Service discovery
	discoveredService, err := discoveryService.Discovery("TestService")
	if err != nil {
		log.Fatalf("Failed to discover service: %v", err)
	}else{
		log.Println("Service discovery successful")
	}

	if discoveredService == nil {
		log.Fatal("Expected to find the registered service, but got nil")
	}else{
		log.Println("Service discovery successful")
	}

	log.Printf("Discovered service: %+v\n", discoveredService)

	// Unregister service
	err = registryService.Unregister(serviceInfo)
	if err != nil {
		log.Fatalf("Failed to unregister service: %v", err)
	}else{
		log.Println("Service unregistered successfully")
	}

	// Clean up resources
	err = registryService.Destroy()
	if err != nil {
		log.Fatalf("Failed to destroy registry service: %v", err)
	}else{
		log.Println("Registry service destroyed successfully")
	}
}
