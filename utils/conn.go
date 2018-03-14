package utils

import (
	"net"
	"time"
	"sync"
	"context"
)

/*
	We provide this dialer wraper ReadTimeout & WriteTimeout attributes into connection object.
	This timeout is for individual buffer I/O operation like other language (python, perl... etc),
	so don't bother with SetDeadline or stupid nethttp.Client timeout.
*/


var connPool sync.Pool

type netConn net.Conn
type netDialer net.Dialer

type Dialer struct {
	*net.Dialer
    ReadTimeout    time.Duration
    WriteTimeout   time.Duration
}

func NewDialer(connTimeout, readTimeout, writeTimeout time.Duration) *Dialer {
	d := &net.Dialer{
		DualStack: false,
		Timeout: connTimeout,
	}
	return &Dialer{d, readTimeout, writeTimeout}
}

func (d *Dialer) Dial(network, addr string) (net.Conn, error) {
    c, err := d.Dialer.Dial(network, addr)
    if err != nil {
        return nil, err
    }
    conn := NewConn(c)
    conn.readTimeout = d.ReadTimeout
    conn.writeTimeout = d.WriteTimeout
    return conn, nil
}

func (d *Dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
    c, err := d.Dialer.DialContext(ctx, network, addr)
    if err != nil {
        return nil, err
    }
    conn := NewConn(c)
    conn.readTimeout = d.ReadTimeout
    conn.writeTimeout = d.WriteTimeout
    return conn, nil
}

type Conn struct {
    netConn
    readTimeout  time.Duration
    writeTimeout time.Duration
    timeoutFunc  func() bool
}

func NewConn(c netConn) *Conn {
    conn, ok := c.(*Conn)
    if ok {
        return conn
    }
    conn, ok = connPool.Get().(*Conn)
    if !ok {
        conn = new(Conn)
    }
    conn.netConn = c
    return conn
}

func (c *Conn) SetReadTimeout(d time.Duration) {
    if c.readTimeout > 0 {
        c.netConn.SetReadDeadline(time.Time{})
    }
    c.readTimeout = d
}

func (c *Conn) SetWriteTimeout(d time.Duration) {
    if c.writeTimeout > 0 {
        c.netConn.SetWriteDeadline(time.Time{})
    }
    c.writeTimeout = d
}

func (c Conn) Read(buf []byte) (n int, err error) {
    if c.readTimeout > 0 {
        c.SetReadDeadline(time.Now().Add(c.readTimeout))
    }
    n, err = c.netConn.Read(buf)
    return
}


func (c Conn) Write(buf []byte) (n int, err error) {
    if c.writeTimeout > 0 {
        c.SetWriteDeadline(time.Now().Add(c.writeTimeout))
    }
    n, err = c.netConn.Write(buf)
    return
}

func (c Conn) Close() (err error) {
	if c.netConn == nil {
		return nil
	}
    err = c.netConn.Close()
    connPool.Put(c)
	c.netConn = nil
    c.readTimeout = 0
    c.writeTimeout = 0
    return
}

func IsTimeoutError(err error) bool {
    e, ok := err.(net.Error)
    if ok {
        return e.Timeout()
    }
    return false
}
