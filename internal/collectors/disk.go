package collectors

import (
	"github.com/StackExchange/wmi"
	"github.com/irshadmb/system-info/internal/models"
)

type DiskCollector struct{}

func NewDiskCollector() *DiskCollector {
	return &DiskCollector{}
}

func (c *DiskCollector) Collect() ([]models.DiskDrive, error) {
	var diskInfo []models.DiskDrive
	query := "SELECT Caption, Size, InterfaceType FROM Win32_DiskDrive"
	err := wmi.Query(query, &diskInfo)
	return diskInfo, err
}