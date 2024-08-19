package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sirupsen/logrus"
)

func main() {
	//fmt.Println("please enter the host and port to connect(ip:host):")
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		logrus.Fatal(err)
	}
}
