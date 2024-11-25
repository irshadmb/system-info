package collectors

import (
	"github.com/StackExchange/wmi"
	"github.com/irshadmb/system-info/internal/models"
)

type OSCollector struct{}

func NewOSCollector() *OSCollector {
	return &OSCollector{}
}

func (c *OSCollector) Collect() ([]models.OperatingSystem, error) {
	var osInfo []models.OperatingSystem
	query := "SELECT Caption, Version, BuildNumber, OSArchitecture, TotalVisibleMemorySize, FreePhysicalMemory, TotalVirtualMemorySize, FreeVirtualMemory FROM Win32_OperatingSystem"
	err := wmi.Query(query, &osInfo)
	return osInfo, err
}