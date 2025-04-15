package main

import "expert-system/ui"

func main() {
	app := ui.InitMainWindow()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
