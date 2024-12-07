package views

import (
	"crypto/rsa"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"client/interfaces"
	"client/models"
)

type ChatView struct {
	messages      []string
	chatViewModel *ChatViewModel
	width         int
	height        int
}

func InitChatView(user *models.User, connections *map[string]interfaces.SocketConnection,
	paymentProcessor interfaces.PaymentProcessor, certificateService interfaces.CertificateService,
	privateKey *rsa.PrivateKey) ChatView {
	var err error
	encryptionCertificate, err := certificateService.LoadCertificate(&user.Certificate)
	if err != nil {
		panic("failed to load encryption certificate")
	}

	chatView := ChatView{
		chatViewModel: &ChatViewModel{
			user:                  user,
			nodePayments:          make(map[string]string),
			encryptionCertificate: encryptionCertificate,
			CertifciateService:    certificateService,
			NodeConnections:       connections,
			PaymentProcessor:      paymentProcessor,
			CertificatePrivateKey: privateKey,
		},
	}

	return chatView
}

func (c ChatView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			fmt.Println(c.chatViewModel.input)
			go c.chatViewModel.ProcessInput()
		case "backspace":
			if len(c.chatViewModel.input) > 0 {
				c.chatViewModel.input = c.chatViewModel.input[:len(c.chatViewModel.input)-1]
			}
		case "ctrl+c", "esc":
			return c, tea.Quit
		default:
			c.chatViewModel.input += msg.String()
		}

	case tea.WindowSizeMsg:
		c.width = msg.Width
		c.height = msg.Height

	}

	return c, nil
}

func (c ChatView) Init() tea.Cmd {
	c.chatViewModel.init()
	return nil
}

func (c ChatView) View() string {
	s := "Connection established with:" + c.chatViewModel.user.Name + " \n\n"

	for _, msg := range c.messages {
		s += msg + "\n\n"
	}
	s += "Input: " + c.chatViewModel.input + "\n\n"
	return s
}
