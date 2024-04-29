package scheduleapp_test

import (
	"os"
	"testing"

	scheduleapp "github.com.whiskeyjack4code/CLI-Tools/Schedule-App"
)

func TestAddAppointment(t *testing.T) {
	l := scheduleapp.AppList{}

	task := "Doctor Visit"
	l.AddAppointment(task)

	if l[0].Name != task {
		t.Errorf("expected %s, got %s", task, l[0].Name)
	}
}

func TestSetVisitedByID(t *testing.T) {
	l := scheduleapp.AppList{}

	task := "Doctor Visit"
	l.AddAppointment(task)

	if l[0].Name != task {
		t.Errorf("expected %s, got %s", task, l[0].Name)
	}

	appID := 1

	if l[0].Attended {
		t.Errorf("appointment should not be be set to attended yet")
	}

	l.SetVisitedByID(appID)

	if !l[0].Attended {
		t.Errorf("appointment should be set to attended yet")
	}
}

func TestSetVisitedByName(t *testing.T) {
	l := scheduleapp.AppList{}

	task := "Doctor Visit"
	l.AddAppointment(task)

	if l[0].Name != task {
		t.Errorf("expected %s, got %s", task, l[0].Name)
	}

	appName := "Doctor Visit"

	if l[0].Attended {
		t.Errorf("appointment should not be be set to attended yet")
	}

	l.SetVisitedByName(appName)

	if !l[0].Attended {
		t.Errorf("appointment should be set to attended yet")
	}
}

func TestDeleteAppByID(t *testing.T) {
	l := scheduleapp.AppList{}

	apps := []string{
		"Real Estate Agent",
		"Work Meeting",
		"Birthday Party",
	}

	for _, v := range apps {
		l.AddAppointment(v)
	}

	if l[0].Name != apps[0] {
		t.Errorf("expected %s, got %s", apps[0], l[0].Name)
	}

	l.DeleteAppByID(2)

	if apps[2] != l[1].Name {
		t.Errorf("%s expected, but got %s", apps[2], l[1].Name)
	}
}

func TestDeleteAppByName(t *testing.T) {
	l := scheduleapp.AppList{}

	apps := []string{
		"Real Estate Agent",
		"Work Meeting",
		"Birthday Party",
	}

	for _, v := range apps {
		l.AddAppointment(v)
	}

	if l[0].Name != apps[0] {
		t.Errorf("expected %s, got %s", apps[0], l[0].Name)
	}

	l.DeleteAppByName("Real Estate Agent")

	if apps[2] != l[1].Name {
		t.Errorf("%s expected, but got %s", apps[2], l[1].Name)
	}
}

func TestSaveAndRetrieve(t *testing.T) {

	l1 := scheduleapp.AppList{}
	l2 := scheduleapp.AppList{}

	task := "Doctor Visit"
	l1.AddAppointment(task)

	if l1[0].Name != task {
		t.Errorf("expected %s, got %s", task, l1[0].Name)
	}

	temp, err := os.CreateTemp("", "")

	if err != nil {
		t.Errorf("error created temp file: %s", err)
	}

	defer os.Remove(temp.Name())

	if err := l1.SaveApp(temp.Name()); err != nil {
		t.Errorf("error saving file %s", err)
	}

	if err := l2.RetrieveApp(temp.Name()); err != nil {
		t.Errorf("error opening file %s", err)
	}

	if l1[0].Name != l2[0].Name {
		t.Errorf("App %s, should be the same as %s", l2[0].Name, l1[0].Name)
	}

}
