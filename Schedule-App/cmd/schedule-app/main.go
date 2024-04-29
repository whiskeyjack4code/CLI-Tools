package main

import (
	"fmt"
	"os"
	"strings"

	scheduleapp "github.com.whiskeyjack4code/CLI-Tools/Schedule-App"
)

const fileName = ".apps.json"

func main() {

	list := &scheduleapp.AppList{}

	if err := list.RetrieveApp(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, v := range *list {
			fmt.Println(v.Name)
		}
	default:
		app := strings.Join(os.Args[1:], " ")
		list.AddAppointment(app)

		if err := list.SaveApp(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
