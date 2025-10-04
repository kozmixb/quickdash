package stats

import (
	"runtime"

	"github.com/shirou/gopsutil/v3/disk"
)

const GIGABYTE float64 = 1024 * 1024 * 1024

type DiskInfo struct {
	Free    float64
	Percent float64
	Used    float64
	Total   float64
}

func getMountPath() string {
	if runtime.GOOS == "windows" {
		return "C:\\"
	}

	return "/"
}

func ReadDiskInfo() (DiskInfo, error) {
	mountPath := getMountPath()
	diskInfo, err := disk.Usage(mountPath)
	if err != nil {
		return DiskInfo{}, err
	}

	return DiskInfo{
		Percent: convertTo2Decimals(diskInfo.UsedPercent),
		Free:    convertTo2Decimals(float64(diskInfo.Free) / GIGABYTE),
		Used:    convertTo2Decimals(float64(diskInfo.Used) / GIGABYTE),
		Total:   convertTo2Decimals(float64(diskInfo.Total) / GIGABYTE),
	}, nil
}
