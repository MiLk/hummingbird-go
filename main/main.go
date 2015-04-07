package main

import (
	"fmt"
	"os"

	"github.com/howeyc/gopass"
	"github.com/jawher/mow.cli"
	"github.com/milk/hummingbird"
)

func printErrors(errs []error) {
	for _, err := range errs {
		fmt.Println(err)
	}
}

func main() {
	app := cli.App("hummingbird", "Hummingbird Go Client")
	username := app.StringArg("USERNAME", "", "Hummingbird username")
	app.Action = func() {
		fmt.Printf("Please enter your password:")
		password := gopass.GetPasswd()
		api := hummingbird.NewAPI()
		if errs, _ := api.UserAuthenticate(*username, "", string(password)); len(errs) != 0 {
			printErrors(errs)
			return
		}

		errs, user := api.UserInformation(*username)
		if len(errs) != 0 {
			printErrors(errs)
			return
		}
		fmt.Printf("user: %+v", user)
	}

	app.Run(os.Args)
}
