package goport

import (
	"math/rand"
	"net"
	"time"
)

func GetFreePort() int {
	return GetFreeRangePort(5000, 49151)
}

func GetFreeRangePort(startPort int, endPort int) int {

	unavailables := make(map[int]interface{})

	var resultPort int = -1

	for {
		iPort := randInt(startPort, endPort)

		if _, ok := unavailables[iPort]; ok {
			continue
		}

		if len(unavailables) > endPort-startPort+1 {
			break
		}

		if CheckFree(iPort) {
			resultPort = iPort
			break
		}

		unavailables[iPort] = nil
	}

	return resultPort
}

func CheckFree(port int) bool {

	var isFree = true

	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: port,
	}

	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if err == nil {
		isFree = false
		conn.Close()
	}

	return isFree
}

func randInt(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Intn(max-min)
}
