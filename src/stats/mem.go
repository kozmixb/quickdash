package stats

import (
	"github.com/shirou/gopsutil/v3/mem"
)

type MemInfo struct {
	Available float64
	Free      float64
	Total     float64
	Percent   float64
}

func ReadMemInfo() (MemInfo, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return MemInfo{}, err
	}

	return MemInfo{
		Available: convertTo2Decimals(float64(v.Available) / GIGABYTE),
		Free:      convertTo2Decimals(float64(v.Free) / GIGABYTE),
		Total:     convertTo2Decimals(float64(v.Total) / GIGABYTE),
		Percent:   convertTo2Decimals(float64(v.UsedPercent)),
	}, nil
}
