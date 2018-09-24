package tui

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
	"github.com/nqbao/learn-go/chatserver/client"
)

func StartUi(c *client.ChatClient) {
	// loginView := NewLoginView()
	chatView := NewChatView()

	chatView.OnMessage(func(msg string) {
		c.Send(fmt.Sprintf("%v\n", msg))
	})

	ui, err := tui.New(chatView)
	if err != nil {
		panic(err)
	}

	quit := func() { ui.Quit() }

	ui.SetKeybinding("Esc", quit)
	ui.SetKeybinding("Ctrl+c", quit)

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
