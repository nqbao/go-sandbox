package tui

import (
	"fmt"
	"strings"

	"github.com/marcusolsson/tui-go"
	"github.com/nqbao/learn-go/chatserver/client"
)

func StartUi(c *client.ChatClient) {
	// ref: https://github.com/marcusolsson/tui-go/blob/master/example/chat/main.go
	history := tui.NewVBox()

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	input.OnSubmit(func(e *tui.Entry) {
		if e.Text() != "" {
			c.Send(fmt.Sprintf("%v\n", e.Text()))
			e.SetText("")
		}
	})

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	box := tui.NewVBox(
		historyBox,
		inputBox,
	)

	ui, err := tui.New(box)
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
				history.Append(
					tui.NewHBox(
						tui.NewLabel(strings.TrimSpace(msg)),
					),
				)
			})
		}
	}()

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
