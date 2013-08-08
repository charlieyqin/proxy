package throttle

import (
    "net"
)

type Conn struct {
	net.Conn
    Throttler
}

func (c Conn) Close() error {
    _ = c.Throttler.Release()
    err := c.Conn.Close()
	return err
}

type Listener struct {
	net.Listener
    Throttler
}

func (l Listener) Accept() (c net.Conn, err error) {
    err = l.Throttler.Acquire()
    if err != nil {
        panic(err)
    }

	c, err = l.Listener.Accept()
	lc := Conn{c, l.Throttler}
	return lc, err
}

func NewListener(addr string, maxConn uint64) net.Listener {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	tl := Listener{l, NewCountingThrottler(maxConn)}
	return tl
}

