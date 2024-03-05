package mockdb

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

func TestCreateUser(t *testing.T) {
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

	db := InitUsersMockDB()

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.CreateUser(domain.User{
				Login: currentCase.login,
			})

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.CreateUser(domain.User{
				Login: currentCase.login,
			})

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestRemoveUser(t *testing.T) {
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

	db := InitUsersMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData] = domain.User{
			Login: currentData,
		}
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.RemoveUser(currentCase.login)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.RemoveUser(currentCase.login)

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

	db := InitUsersMockDB()

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
		status   string
		version  uint8
	}{
		{
			"get existed user",
			"Ahmed",
			"ahded",
			"guest",
			1,
		},
		{
			"get existed user",
			"Dima",
			"BestDimaEver",
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

	db := InitUsersMockDB()

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
				Login:   currentCase.login,
				Name:    currentCase.name,
				Status:  currentCase.status,
				Version: currentCase.version,
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

func TestChangeUserPassword(t *testing.T) {
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

	db := InitUsersMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Login] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := db.ChangeUserPassword(currentCase.login, currentCase.newPassword)

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

func TestChangeUserName(t *testing.T) {
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
		status   string
		version  uint8
	}{
		{
			"get existed user",
			"Nikita",
			"Ahmed",
			"Nikita",
			"guest",
			2,
		},
		{
			"get existed user",
			"ArArA",
			"Dima",
			"ArArA",
			"publisher",
			3,
		},
	}

	db := InitUsersMockDB()

	for _, currentData := range data {
		db.mutex.Lock()
		db.storage[currentData.Login] = currentData
		db.mutex.Unlock()
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			user, err := db.ChangeUserName(currentCase.login, currentCase.newName)

			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, true, reflect.DeepEqual(user, domain.User{
				Login:   currentCase.login,
				Name:    currentCase.name,
				Status:  currentCase.status,
				Version: currentCase.version,
			}))
		})
	}
}
