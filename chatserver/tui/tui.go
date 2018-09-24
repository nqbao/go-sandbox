package tui

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
	"github.com/nqbao/learn-go/chatserver/client"
)

func StartUi(c *client.ChatClient) {
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
		ui.SetWidget(chatView)
	})

	chatView.OnMessage(func(msg string) {
		c.Send(fmt.Sprintf("%v\n", msg))
	})

	go func() {
		for msg := range c.Incoming {
			// we need to make the change via ui update to make sure the ui is repaint correctly
			ui.Update(func() {
				chatView.AddMessage(msg)
			})
		}
	}()

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
