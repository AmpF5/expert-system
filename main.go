package main

func main() {
	app := InitUI()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
