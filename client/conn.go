package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func tcpConnect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	_, err = conn.Write([]byte("echo"))
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	return nil
}
