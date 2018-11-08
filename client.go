package gedis

import (
	"net"
	"fmt"
	"github.com/fatih/pool"
	"bufio"
)

type client struct {
	Host string
	Port int
	InitialCap int
	MaxCap int
	P pool.Pool
}

func Dial(host string, port int, args ...int) (*client, error) {
	initialCap, maxCap := 5, 30
	switch {
	case len(args) >= 2:
		initialCap, maxCap = args[0], args[1]
	case len(args) == 1:
		initialCap = args[0]
	}

	addr := fmt.Sprintf("%s:%d", host, port)

	factory := func() (net.Conn, error) { return net.Dial("tcp", addr) }
	p, err := pool.NewChannelPool(initialCap, maxCap, factory)
	if err != nil {
		return nil, nil
	}

	c := &client{
		host,
		port,
		initialCap,
		maxCap,
		p,
	}
	return c, nil
}

func (c *client) Do(messages ...string) (string, error) {
	conn, err := c.P.Get()
	if err != nil {
		return "", err
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	err = send(writer, messages...)
	if err != nil {
		return "", err
	}

	message, err := recv(reader)
	if err != nil {
		return "", err
	}
	return message, nil
}

func (c *client) Close() {
	c.P.Close()
}