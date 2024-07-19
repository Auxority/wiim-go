package network

import (
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", fmt.Errorf("failed to get interface addresses: %w", err)
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no IP address found")
}

func Scan(checkFunc func(string) bool) string {
	localIP, err := GetLocalIP()
	if err != nil {
		log.Error().Err(err).Msg("Error getting local IP")
		return ""
	}

	ipRange := localIP[:strings.LastIndex(localIP, ".")+1]

	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	var ipAddress string

	for i := 1; i <= 254; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()

			currentIP := fmt.Sprintf("%s%d", ipRange, i)
			log.Debug().Str("ipAddress", currentIP).Msg("Scanning IP")

			if checkFunc(currentIP) {
				mutex.Lock()
				ipAddress = currentIP
				mutex.Unlock()
			}
		}(i)
	}

	waitGroup.Wait()

	return ipAddress
}
