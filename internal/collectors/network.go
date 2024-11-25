package collectors

import (
	"github.com/StackExchange/wmi"
	"github.com/irshadmb/system-info/internal/models"
)

type NetworkCollector struct{}

func NewNetworkCollector() *NetworkCollector {
	return &NetworkCollector{}
}

func (c *NetworkCollector) Collect() ([]models.NetworkAdapter, error) {
	var networkInfo []models.NetworkAdapter
	query := "SELECT Name, AdapterType, MACAddress, Speed, PhysicalAdapter FROM Win32_NetworkAdapter WHERE PhysicalAdapter=true"
	err := wmi.Query(query, &networkInfo)
	return networkInfo, err
}