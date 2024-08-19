package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)
type tcpClient struct {
	textinput textinput.Model
	err       error
	banner    string
	connected bool
}

func initialModel() tcpClient {
	ti := textinput.New()
	ti.Placeholder = "localhost:8080"
	ti.Focus()
	ti.CharLimit = 1000
	return tcpClient{
		textinput: ti,
		err:       nil,
		banner:    "please enter the server(ip:port)",
		connected: false,
	}
}

func (tc tcpClient) Init() tea.Cmd {
	return textinput.Blink
}

func (tc tcpClient) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return tc, tea.Quit
		case tea.KeyEnter:
			tcpConnect(msg.String())
			if tc.connected {
				tc.banner = "test"
				tc.textinput.Placeholder = "please enter the command"
			}
		}
	case errMsg:
		tc.err = msg
		return tc, nil
	}
	tc.textinput, cmd = tc.textinput.Update(msg)
	return tc, cmd
}

func (tc tcpClient) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n%s", tc.banner,
		tc.textinput.View(),
		"(esc to quit)") + "\n"
}
