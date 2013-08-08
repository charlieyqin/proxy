package throttle

import (
    "net"
    "log"
)

type Conn struct {
	net.Conn
	sema chan struct{}
}

func (c Conn) Close() error {
	var v struct{}
	c.sema <- v
    err := c.Conn.Close()
    if err == nil {
        log.Println(" - ", c.Conn)
    } else {
        log.Println("Close error: ", err)
    }
	return err
}

type Listener struct {
	net.Listener
	sema chan struct{}
}

func (l Listener) Accept() (c net.Conn, err error) {
	<-l.sema
	c, err = l.Listener.Accept()
	lc := Conn{c, l.sema}
    log.Println(" + ", lc.Conn)
	return lc, err
}

func NewListener(addr string, maxConn uint64) net.Listener {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	tl := Listener{l, make(chan struct{}, maxConn)}
	for i := uint64(0); i < maxConn; i++ {
		var v struct{}
		tl.sema <- v
	}
	return tl
}

