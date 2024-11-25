package collectors

import (
	"github.com/StackExchange/wmi"
	"github.com/irshadmb/system-info/internal/models"
)

type CPUCollector struct{}

func NewCPUCollector() *CPUCollector {
	return &CPUCollector{}
}

func (c *CPUCollector) Collect() ([]models.Processor, error) {
	var cpuInfo []models.Processor
	query := "SELECT Name, NumberOfCores, NumberOfLogicalProcessors, MaxClockSpeed, LoadPercentage FROM Win32_Processor"
	err := wmi.Query(query, &cpuInfo)
	return cpuInfo, err
}