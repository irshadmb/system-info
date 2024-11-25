package models

type SystemInfo struct {
	OperatingSystem   []OperatingSystem   `json:"os"`
	Processors        []Processor         `json:"processors"`
	DiskDrives        []DiskDrive        `json:"disk_drives"`
	NetworkAdapters   []NetworkAdapter    `json:"network_adapters"`
}

type OperatingSystem struct {
	Caption                string  `json:"caption"`
	Version                string  `json:"version"`
	BuildNumber            string  `json:"build_number"`
	OSArchitecture         string  `json:"architecture"`
	TotalVisibleMemorySize uint64  `json:"total_memory"`
	FreePhysicalMemory     uint64  `json:"free_memory"`
	TotalVirtualMemorySize uint64  `json:"total_virtual_memory"`
	FreeVirtualMemory      uint64  `json:"free_virtual_memory"`
}

type Processor struct {
	Name                    string  `json:"name"`
	NumberOfCores           uint32  `json:"cores"`
	NumberOfLogicalProcessors uint32  `json:"logical_processors"`
	MaxClockSpeed           uint32  `json:"max_clock_speed"`
	LoadPercentage          uint16  `json:"load_percentage"`
}

type DiskDrive struct {
	Caption       string  `json:"caption"`
	Size          uint64  `json:"size"`
	InterfaceType string  `json:"interface_type"`
}

type NetworkAdapter struct {
	Name            string  `json:"name"`
	AdapterType     string  `json:"adapter_type"`
	MACAddress      string  `json:"mac_address"`
	Speed           uint64  `json:"speed"`
	PhysicalAdapter bool    `json:"is_physical"`
}