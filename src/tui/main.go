package tui

func TUIMAIN() {
	initUI()
	// Manage Database Initialisation
	initDB()
	defer db.Close()

	if err := app.
		SetRoot(page, true).
		SetFocus(page).
		Run(); err != nil {
		panic(err)
	}
}
