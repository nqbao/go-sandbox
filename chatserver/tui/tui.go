package tui

import (
	"fmt"
	"io"

	"github.com/marcusolsson/tui-go"
	"github.com/nqbao/learn-go/chatserver/client"
)

func StartUi(c client.ChatClient) {
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
		c.SendMessage(msg)
	})

	go func() {
		for {
			select {
			case err := <-c.Error():

				if err == io.EOF {
					ui.Update(func() {
						chatView.AddMessage("Connection closed connection from server.")
					})
				} else {
					panic(err)
				}
			case msg := <-c.Incoming():
				// we need to make the change via ui update to make sure the ui is repaint correctly
				ui.Update(func() {
					chatView.AddMessage(fmt.Sprintf("%v: %v", msg.Name, msg.Message))
				})
			}
		}
	}()

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
