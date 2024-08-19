package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"net"
	"os/exec"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			logrus.Warn("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		logrus.Infoln("receive from client:", recvStr)
		result := execShell(recvStr)
		write, err := conn.Write(result)
		if err != nil {
			logrus.Warn("write to client failed, err:", err)
			break
		}
		if write != len(result) {
			logrus.Warnln("write to client failed, write length:", write)
			break
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10086")
	if err != nil {
		logrus.Panicln(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			logrus.Warnln("accept failed,err:", err)
			continue
		}
		process(conn)
	}
}

func execShell(command string) []byte {
	c := exec.Command("bash", "-c", command)
	output, err := c.CombinedOutput()
	if err != nil {
		logrus.Fatalln("execute failed,err:", err)
		return nil
	}
	return output
}
