package net

import (
	n "net"
	"sync/atomic"
)

type Conn n.TCPConn
type ServerState uint32
type NetPort int

type Host struct {
	Hostname 		string
	IpAddr 			n.IPAddr
	Port 			NetPort
}

type NetworkUsage struct {
	TotalBytesSent		uint64
	TotalBytesReceived 	uint64
}

type Server struct {
	Host Host
	Weight uint8
	MaxConns uint16
	State ServerState
	NetUsage *NetworkUsage
}

type ServerTcpConn struct {
	s *Server
	Conn *Conn
}

func (ser *ServerTcpConn) SendMsg(msg []byte) {
	 ser.Conn.Write(msg)
	atomic.AddUint64(&ser.s.NetUsage.TotalBytesSent, uint64((len(msg))))
}

func RecvMsg() {

}