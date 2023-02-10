package main

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	if err := app.server.ListenAndServe(); err != nil {
		panic(err)
	}
}
