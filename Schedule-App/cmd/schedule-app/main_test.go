package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binFile  = "schedule"
  fileName string
)

func TestMain(m *testing.M) {
	fmt.Println("Building app...")

  var envFile = os.Getenv("APPS_FILENAME")
  if envFile != "" {
    fileName = envFile
  }

	if runtime.GOOS == "windows" {
		binFile += ".exe"
	}

	build := exec.Command("go", "build", "-o", binFile)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error running %s for build %s", binFile, err)
		os.Exit(1)
	}

	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binFile)
	os.Remove(fileName)
	os.Exit(result)

}

func TestSchedulerCLI(t *testing.T) {
	app1 := "test appointment 1"

	dir, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binFile)

	t.Run("AddAppointmentFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", app1)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

  app2 := "test appointment number 2"

  t.Run("AddAppointmentFromSTDIN", func(t *testing.T){
    cmd := exec.Command(cmdPath, "-add")
    in, err := cmd.StdinPipe()
     
    if err != nil {
      t.Fatal(err)
    }

    io.WriteString(in, app2)
    in.Close()

    if err := cmd.Run(); err != nil {
      t.Fatal(err)
    }
  })

	t.Run("ListAppointments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")

		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}

    expected := fmt.Sprintf(" 1: %s\n 2: %s", app1, app2)

		if expected != string(out) {
			t.Errorf("expected %q, got %q instead", expected, string(out))
		}
	})
}
