package tui

func TUIMAIN() {
	initUI()

	if err := app.
		SetRoot(page, true).
		SetFocus(page).
		Run(); err != nil {
		panic(err)
	}
}
