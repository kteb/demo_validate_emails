package main

import (
	"log"

	"github.com/kteb/demo_validate_emails/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
