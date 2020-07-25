package tpm

import (
	"fmt"
	"time"

	"github.com/google/go-tspi/tpmclient"
)

const host = "192.168.122.1"
const port = "12041"

// we're connecting to bridge, so 10ms is a reasonable timeout value
const timeout = 10 * time.Millisecond

// Extend extends the TPM log with the provided string. Returns any error.
func Extend(description string) error {
	hostPort := fmt.Sprintf("%s:%s", host, port)
	//connection := tpmclient.New("192.168.122.1:12041", timeout)
	connection := tpmclient.New(hostPort, timeout)
	err := connection.Extend(15, 0x1000, nil, description)
	return err
}
