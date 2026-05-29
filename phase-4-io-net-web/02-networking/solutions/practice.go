package solutions

import (
	"net"
)


func StartEchoTCPServer(addr string, stopChan chan struct{}) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	go func() {
		<-stopChan
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-stopChan:
				return nil
			default:
				return err
			}
		}
		go func(c net.Conn) {
			defer c.Close()
			buf := make([]byte, 1024)
			n, err := c.Read(buf)
			if err != nil {
				return
			}
			msg := string(buf[:n])
			c.Write([]byte("ECHO: " + msg))
		}(conn)
	}
}

func SendTCPMessage(addr, message string) (string, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(message)); err != nil {
		return "", err
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}
