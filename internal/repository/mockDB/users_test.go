package mockdb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

func TestCreate(t *testing.T) {
	validCases := []struct {
		testName string
		login    string
	}{
		{
			"add new user",
			"Ahmed",
		},
		{
			"add new user",
			"Dima",
		},
		{
			"add new user",
			"Danyenchka",
		},
	}
	invalidCases := []struct {
		testName string
		login    string
	}{
		{
			"add existed user",
			"Ahmed",
		},
		{
			"add existed user",
			"Ahmed",
		},
		{
			"add existed user",
			"Dima",
		},
		{
			"add existed user",
			"Danyenchka",
		},
	}

	db := InitMockDB()

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.Create(domain.User{
				Login: currentCase.login,
			})

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.Create(domain.User{
				Login: currentCase.login,
			})

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestRemove(t *testing.T) {
	data := []string{
		"Ahmed", "Dima", "Danyenchka",
	}

	validCases := []struct {
		testName string
		login    string
	}{
		{
			"remove existed user",
			"Ahmed",
		},
		{
			"remove existed user",
			"Dima",
		},
		{
			"remove existed user",
			"Danyenchka",
		},
	}
	invalidCases := []struct {
		testName string
		login    string
	}{
		{
			"remove unexisted user",
			"Ahmed",
		},
		{
			"remove unexisted user",
			"Ahmed",
		},
		{
			"remove unexisted user",
			"Dima",
		},
		{
			"remove unexisted user",
			"Danyenchka",
		},
	}

	db := InitMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData] = domain.User{
			Login: currentData,
		}
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.Remove(currentCase.login)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.Remove(currentCase.login)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestHasUser(t *testing.T) {
	data := []domain.User{
		{
			Login:    "Ahmed",
			Password: "root",
		},
		{
			Login:    "Dima",
			Password: "AbobA",
		},
	}

	validCases := []struct {
		testName string
		login    string
		password string
	}{
		{
			"check for existed user",
			"Ahmed",
			"root",
		},
		{
			"check for existed user",
			"Dima",
			"AbobA",
		},
	}
	invalidCases := []struct {
		testName string
		login    string
		password string
	}{
		{
			"check for unexisted user",
			"asfasfd",
			"root",
		},
		{
			"check for unexisted user",
			"Dima",
			"aboba",
		},
	}

	db := InitMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Login] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.HasUser(currentCase.login, currentCase.password)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.HasUser(currentCase.login, currentCase.password)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	data := []domain.User{
		{
			Login:    "Ahmed",
			Name:     "ahded",
			Password: "root",
			Status:   "guest",
			Version:  1,
		},
		{
			Login:    "Dima",
			Name:     "BestDimaEver",
			Password: "AbobA",
			Status:   "publisher",
			Version:  2,
		},
	}

	validCases := []struct {
		testName string
		login    string
		name     string
		password string
		status   string
		version  uint8
	}{
		{
			"get existed user",
			"Ahmed",
			"ahded",
			"root",
			"guest",
			1,
		},
		{
			"get existed user",
			"Dima",
			"BestDimaEver",
			"AbobA",
			"publisher",
			2,
		},
	}
	invalidCases := []struct {
		testName string
		login    string
	}{
		{
			"get unexisted user",
			"asfasfd",
		},
		{
			"get unexisted user",
			"Dimaa",
		},
	}

	db := InitMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Login] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			user, err := db.GetUser(currentCase.login)

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, true, reflect.DeepEqual(user, domain.User{
				Login:    currentCase.login,
				Name:     currentCase.name,
				Password: currentCase.password,
				Status:   currentCase.status,
				Version:  currentCase.version,
			}))
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			_, err := db.GetUser(currentCase.login)

			if err == nil {
				t.Error(err)
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	data := []domain.User{
		{
			Login:    "Ahmed",
			Name:     "ahded",
			Password: "root",
			Status:   "guest",
			Version:  1,
		},
		{
			Login:    "Dima",
			Name:     "BestDimaEver",
			Password: "AbobA",
			Status:   "publisher",
			Version:  2,
		},
	}

	validCases := []struct {
		testName    string
		newPassword string
		login       string
		name        string
		password    string
		status      string
		version     uint8
	}{
		{
			"get existed user",
			"newPassword",
			"Ahmed",
			"ahded",
			"newPassword",
			"guest",
			2,
		},
		{
			"get existed user",
			"root",
			"Dima",
			"BestDimaEver",
			"root",
			"publisher",
			3,
		},
	}

	db := InitMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Login] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.ChangePassword(currentCase.login, currentCase.newPassword)

			if err != nil {
				t.Error(err)
			}

			db.mutex.Lock()
			user, ok := db.storage[currentCase.login]
			db.mutex.Unlock()

			if !ok {
				t.Error("user doesnt exist")
			}

			assert.Equal(t, true, reflect.DeepEqual(user, domain.User{
				Login:    currentCase.login,
				Name:     currentCase.name,
				Password: currentCase.password,
				Status:   currentCase.status,
				Version:  currentCase.version,
			}))
		})
	}
}

func TestChangeName(t *testing.T) {
	data := []domain.User{
		{
			Login:    "Ahmed",
			Name:     "ahded",
			Password: "root",
			Status:   "guest",
			Version:  1,
		},
		{
			Login:    "Dima",
			Name:     "BestDimaEver",
			Password: "AbobA",
			Status:   "publisher",
			Version:  2,
		},
	}

	validCases := []struct {
		testName string
		newName  string
		login    string
		name     string
		password string
		status   string
		version  uint8
	}{
		{
			"get existed user",
			"Nikita",
			"Ahmed",
			"Nikita",
			"root",
			"guest",
			2,
		},
		{
			"get existed user",
			"ArArA",
			"Dima",
			"ArArA",
			"AbobA",
			"publisher",
			3,
		},
	}

	db := InitMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Login] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			user, err := db.ChangeName(currentCase.login, currentCase.newName)

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, true, reflect.DeepEqual(user, domain.User{
				Login:    currentCase.login,
				Name:     currentCase.name,
				Password: currentCase.password,
				Status:   currentCase.status,
				Version:  currentCase.version,
			}))
		})
	}
}
