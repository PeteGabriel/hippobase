package main

func main() {

	app := Application{}

	err := app.serve()
	if err != nil {
		panic(err)
	}
}
