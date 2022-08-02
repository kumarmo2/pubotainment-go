package misc

import "net"

func GetLocalIp() (string, error) {
	conn, err := net.Dial("udp", "1.2.3.4:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	addr := conn.LocalAddr()
	return addr.String(), nil
}
