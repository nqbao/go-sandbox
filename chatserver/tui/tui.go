package tui

import (
	"github.com/marcusolsson/tui-go"
	"github.com/nqbao/learn-go/chatserver/client"
)

func StartUi(c *client.TcpChatClient) {
	loginView := NewLoginView()
	chatView := NewChatView()

	ui, err := tui.New(loginView)
	if err != nil {
		panic(err)
	}

	quit := func() { ui.Quit() }

	ui.SetKeybinding("Esc", quit)
	ui.SetKeybinding("Ctrl+c", quit)

	loginView.OnLogin(func(username string) {
		c.SetName(username)
		ui.SetWidget(chatView)
	})

	chatView.OnSubmit(func(msg string) {
		c.Send(msg)
	})

	go func() {
		for msg := range c.Incoming {
			// we need to make the change via ui update to make sure the ui is repaint correctly
			ui.Update(func() {
				chatView.AddMessage(msg.Name, msg.Message)
			})
		}
	}()

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
