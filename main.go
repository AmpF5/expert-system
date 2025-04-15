package main

import "expert-system/ui"

func main() {
	app := ui.InitUI()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
