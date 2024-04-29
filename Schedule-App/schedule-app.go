package scheduleapp

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type app struct {
	Name       string
	CreatedAt  time.Time
	Attended   bool
	AttendedAt time.Time
}

type AppList []app

func (a *AppList) AddAppointment(name string) {
	*a = append(*a, app{
		Name:      name,
		CreatedAt: time.Now(),
	})
}

func (a *AppList) SetVisitedByID(i int) error {
	list := *a

	if i <= 0 || i > len(list) {
		return fmt.Errorf("appointment number does not match")
	}

	list[i-1].Attended = true
	list[i-1].AttendedAt = time.Now()

	return nil
}

func (a *AppList) SetVisitedByName(appName string) error {
	list := *a

	for i, v := range list {
		if v.Name == appName {
			list[i].Attended = true
			list[i].AttendedAt = time.Now()
		} else {
			return fmt.Errorf("appointment name does not match")
		}
	}
	return nil
}

func (a *AppList) DeleteAppByID(i int) error {
	list := *a

	if i <= 0 || i > len(list) {
		return fmt.Errorf("appointment id does not match")
	}

	*a = append(list[:i-1], list[i:]...)

	return nil
}

func (a *AppList) DeleteAppByName(name string) error {
	list := *a

	for i, v := range list {
		if v.Name == name {
			*a = append(list[:i], list[i+1:]...)
		} else {
			return fmt.Errorf("appointment name does not match")
		}
	}

	return nil
}

func (a *AppList) SaveApp(fileName string) error {
	js, err := json.Marshal(a)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, js, 0644)
}

func (a *AppList) RetrieveApp(fileName string) error {
	file, err := os.ReadFile(fileName)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, a)
}
