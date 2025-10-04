package stats

import (
	"fmt"
	"net"
	"time"

	"github.com/shirou/gopsutil/v3/host"
)

type HostInfo struct {
	IP       string
	Hostname string
	Arch     string
	Uptime   string
	OS       string
	Platform string
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

func ReadHostInfo() (HostInfo, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return HostInfo{}, err
	}

	uptimeDuration := time.Duration(hostInfo.Uptime) * time.Second

	return HostInfo{
		IP:       getLocalIP(),
		Hostname: hostInfo.Hostname,
		Arch:     hostInfo.KernelArch,
		OS:       hostInfo.OS,
		Platform: hostInfo.Platform,
		Uptime:   fmt.Sprintf("%v", uptimeDuration),
	}, nil
}
