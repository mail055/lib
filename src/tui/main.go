package tui

func TUIMAIN() {
	// Manage Database Initialisation
	initDB()
	defer db.Close()

	// Handle UI Initialisation and Update Logic
	initUI()
	updateUI()

	// Start the application
	if err := app.
		SetRoot(page, true).
		SetFocus(page).
		Run(); err != nil {
		panic(err)
	}
}
