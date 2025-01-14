//go:build darwin
// +build darwin

package tcp

import (
	"github.com/sqkam/xray-core/common/errors"
	"github.com/sqkam/xray-core/common/net"
	"github.com/sqkam/xray-core/transport/internet"
	"github.com/sqkam/xray-core/transport/internet/stat"
)

// GetOriginalDestination from tcp conn
func GetOriginalDestination(conn stat.Connection) (net.Destination, error) {
	la := conn.LocalAddr()
	ra := conn.RemoteAddr()
	ip, port, err := internet.OriginalDst(la, ra)
	if err != nil {
		return net.Destination{}, errors.New("failed to get destination").Base(err)
	}
	dest := net.TCPDestination(net.IPAddress(ip), net.Port(port))
	if !dest.IsValid() {
		return net.Destination{}, errors.New("failed to parse destination.")
	}
	return dest, nil
}
