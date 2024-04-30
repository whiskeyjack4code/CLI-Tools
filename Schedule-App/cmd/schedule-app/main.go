package main

import (
	"flag"
	"fmt"
	"os"

	scheduleapp "github.com.whiskeyjack4code/CLI-Tools/Schedule-App"
)

const fileName = ".apps.json"

func main() {

  appointment := flag.String("appointment", "", "Appointment to Set in Scheduler")
  list := flag.Bool("list", false, "List all Appointments")
  attend := flag.Int("attend", 0, "Appointment attended")

  flag.Parse()

	l := &scheduleapp.AppList{}

	if err := l.RetrieveApp(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, app := range *l {
      if !app.Attended {
        fmt.Println(app)
      }
		}
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
