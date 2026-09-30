package tui

func TUIMAIN() {
	initUI()
	// Manage Database Initialisation
	initDB()
	defer db.Close()

	// Handle UI Initialisation and Update Logic
	initUI()
	updateUI()
	if err := app.
		SetRoot(page, true).
		SetFocus(page).
		Run(); err != nil {
		panic(err)
	}
}
