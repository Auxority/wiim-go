package device

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Auxority/wiim-go/network"
)

const (
	protocol       = "http"
	port           = 49152
	timeout        = 2 * time.Second
	xmlSnippet     = "<?xml"
	descriptionXML = "description.xml"
)

var (
	ErrNotFound = errors.New("WiiM device not found")
)

func Find() string {
	ip := os.Getenv("WIIM_IP")
	if ip == "" {
		ip = network.Scan(isWiiMDevice)
	}

	return ip
}

func isWiiMDevice(ip string) bool {
	url := protocol + "://" + ip + ":" + strconv.Itoa(port) + "/" + descriptionXML
	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return false
		}

		if strings.HasPrefix(string(body), xmlSnippet) {
			return true
		}
	}

	return false
}
