package tui

import tui "github.com/marcusolsson/tui-go"

type LoginView struct {
	tui.Box
	frame *tui.Box
}

func NewLoginView() *LoginView {
	// https://github.com/marcusolsson/tui-go/blob/master/example/login/main.go
	user := tui.NewEntry()
	user.SetFocused(true)

	view := &LoginView{}
	view.frame = tui.NewVBox(
		user,
	)
	view.frame.SetBorder(true)
	view.Append(view.frame)

	return view
}
