package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binFile  = "schedule"
	fileName = ".apps.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building app...")

	if runtime.GOOS == "windows" {
		fileName += ".exe"
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

	t.Run("AddAppointment", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-appointment", app1)

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

		expected := app1 + "\n"

		if expected != string(out) {
		  t.Errorf("expected %q, got %q instead", expected, string(out))
		}
	})
}
