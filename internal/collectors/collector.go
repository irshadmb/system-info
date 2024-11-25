package collectors

import (
	"github.com/irshadmb/system-info/internal/models"
)

type SystemCollector struct {
	osCollector      *OSCollector
	cpuCollector     *CPUCollector
	diskCollector    *DiskCollector
	networkCollector *NetworkCollector
}

func NewSystemCollector() *SystemCollector {
	return &SystemCollector{
		osCollector:      NewOSCollector(),
		cpuCollector:     NewCPUCollector(),
		diskCollector:    NewDiskCollector(),
		networkCollector: NewNetworkCollector(),
	}
}

func (sc *SystemCollector) CollectAll() (*models.SystemInfo, error) {
	sysInfo := &models.SystemInfo{}
	var err error

	// Collect OS info
	sysInfo.OperatingSystem, err = sc.osCollector.Collect()
	if err != nil {
		return nil, err
	}

	// Collect CPU info
	sysInfo.Processors, err = sc.cpuCollector.Collect()
	if err != nil {
		return nil, err
	}

	// Collect disk info
	sysInfo.DiskDrives, err = sc.diskCollector.Collect()
	if err != nil {
		return nil, err
	}

	// Collect network info
	sysInfo.NetworkAdapters, err = sc.networkCollector.Collect()
	if err != nil {
		return nil, err
	}

	return sysInfo, nil
}