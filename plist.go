package plist

import (
	"flag"
)

var P1 = flag.Int("port", 50051, "The server 1 port")
var P2 = flag.Int("port", 50052, "The server 1 port")
var P3 = flag.Int("port", 50053, "The server 1 port")
var P4 = flag.Int("port", 50054, "The server 1 port")
var P5 = flag.Int("port", 50055, "The server 1 port")

type PortList struct{}

func (c PortList) GetPortList() int {
	var p int = 50051
	return p
}
