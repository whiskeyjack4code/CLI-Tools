package main

import (
	"flag"
	"fmt"
	"os"

	scheduleapp "github.com.whiskeyjack4code/CLI-Tools/Schedule-App"
)

var fileName string = ".apps.json"

func main() {

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool developed by WhiskeyJack4Code:\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Copyright 2024 David McMahon")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage:")
		flag.PrintDefaults()
	}

	appointment := flag.String("appointment", "", "Appointment to Set in Scheduler")
	list := flag.Bool("list", false, "List all Appointments")
	attend := flag.Int("attend", 0, "Appointment attended")

	flag.Parse()

	l := &scheduleapp.AppList{}
  
  envFile := os.Getenv("APPS_FILENAME")
  if envFile != "" {
    fileName = envFile
  }
  
	if err := l.RetrieveApp(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {

	case *list:
    fmt.Print(l)

	case *attend > 0:
		if err := l.SetVisitedByID(*attend); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.SaveApp(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *appointment != "":
		l.AddAppointment(*appointment)

		if err := l.SaveApp(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	default:
		fmt.Fprintln(os.Stderr, "invalid option selected")
		os.Exit(1)

	}
}
