package cache

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSession(t *testing.T) {
	validCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"add new session",
			"Ahmed",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0OTEwMjIsImlzcyI6Ik5FVHJ1bm5lckZMSVgiLCJMb2dpbiI6ImRhbmlsYSIsIlN0YXR1cyI6InJlZ3VsYXIiLCJWZXJzaW9uIjoxfQ.mklk9wl5oFN3tn89ET2Svc1QwqU2qn-HKE5m8gUAxfY",
			1,
		},
		{
			"add new session",
			"Dima",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0OTEwMjIsImlzcyI6Ik5FVHJ1bm5lckZMSVgiLCJMb2dpbiI6ImRhbmlsYSIsIlN0YXR1cyI6InJlZ3VsYXIiLCJWZXJzaW9uIjoxfQ.mklk9wl5oFN3tn89ET2Svc1QwqU2qn-HKE5m8gUAxfY",
			1,
		},
		{
			"add new session",
			"Danyenchka",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0OTEwMjIsImlzcyI6Ik5FVHJ1bm5lckZMSVgiLCJMb2dpbiI6ImRhbmlsYSIsIlN0YXR1cyI6InJlZ3VsYXIiLCJWZXJzaW9uIjoxfQ.mklk9wl5oFN3tn89ET2Svc1QwqU2qn-HKE5m8gUAxfY",
			1,
		},
	}
	invalidCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"add existed session",
			"Ahmed",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0OTEwMjIsImlzcyI6Ik5FVHJ1bm5lckZMSVgiLCJMb2dpbiI6ImRhbmlsYSIsIlN0YXR1cyI6InJlZ3VsYXIiLCJWZXJzaW9uIjoxfQ.mklk9wl5oFN3tn89ET2Svc1QwqU2qn-HKE5m8gUAxfY",
			1,
		},
		{
			"add existed session",
			"Ahmed",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0OTEwMjIsImlzcyI6Ik5FVHJ1bm5lckZMSVgiLCJMb2dpbiI6ImRhbmlsYSIsIlN0YXR1cyI6InJlZ3VsYXIiLCJWZXJzaW9uIjoxfQ.mklk9wl5oFN3tn89ET2Svc1QwqU2qn-HKE5m8gUAxfY",
			1,
		},
		{
			"add existed session",
			"Dima",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0OTEwMjIsImlzcyI6Ik5FVHJ1bm5lckZMSVgiLCJMb2dpbiI6ImRhbmlsYSIsIlN0YXR1cyI6InJlZ3VsYXIiLCJWZXJzaW9uIjoxfQ.mklk9wl5oFN3tn89ET2Svc1QwqU2qn-HKE5m8gUAxfY",
			1,
		},
		{
			"add existed session",
			"Danyenchka",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk0OTEwMjIsImlzcyI6Ik5FVHJ1bm5lckZMSVgiLCJMb2dpbiI6ImRhbmlsYSIsIlN0YXR1cyI6InJlZ3VsYXIiLCJWZXJzaW9uIjoxfQ.mklk9wl5oFN3tn89ET2Svc1QwqU2qn-HKE5m8gUAxfY",
			1,
		},
	}

	storage := NewSessionStorage()

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.Add(currentCase.login, currentCase.token, currentCase.version)

			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.Add(currentCase.login, currentCase.token, currentCase.version)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestDeleteSession(t *testing.T) {
	validCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"delete existed session",
			"Ahmed",
			"a.a.a-a",
			1,
		},
		{
			"delete existed session",
			"Dima",
			"a.a.a-a",
			1,
		},
		{
			"delete existed session",
			"Danyenchka",
			"a.a.a-a",
			1,
		},
	}
	invalidCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"delete non-existed session",
			"Ahmed",
			"a.a.a-a",
			1,
		},
		{
			"delete non-existed session",
			"Ahmed",
			"a.a.a-a",
			1,
		},
		{
			"delete non-existed session",
			"Dima",
			"a.a.a-a",
			1,
		},
		{
			"delete non-existed session",
			"Danyenchka",
			"a.a.a-a",
			1,
		},
	}

	storage := NewSessionStorage()

	for _, currentCase := range validCases {
		err := storage.Add(currentCase.login, currentCase.token, currentCase.version)
		if err != nil {
		}
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.DeleteSession(currentCase.login, currentCase.token)
			if err != nil {
				t.Error(err)
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.DeleteSession(currentCase.login, currentCase.token)
			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestHasSession(t *testing.T) {
	validCases := []struct {
		testName string
		login    string
		token    string
	}{
		{
			"check for existed session",
			"Ahmed",
			"a.a.a-a",
		},
		{
			"check for existed session",
			"Dima",
			"a.a.a-a",
		},
	}
	invalidCases := []struct {
		testName string
		login    string
		token    string
	}{
		{
			"check for unexisted user",
			"asfasfd",
			"a.a.a-a",
		},
		{
			"check for unexisted user",
			"sadas",
			"a.a.a-a",
		},
	}

	storage := NewSessionStorage()

	for _, currentCase := range validCases {
		err := storage.Add(currentCase.login, currentCase.token, 1)
		if err != nil {
		}
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.HasSession(currentCase.login, currentCase.token)

			if err != nil {
				t.Error("no such session")
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.HasSession(currentCase.login, currentCase.token)

			if err == nil {
				t.Error("no error returned")
			}
		})
	}
}

func TestGetVersion(t *testing.T) {
	validCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"get existed version",
			"Ahmed",
			"a.a.a-a",
			1,
		},
		{
			"get existed version",
			"Dima",
			"a.a.a-a",
			1,
		},
		{
			"get existed version",
			"Danyenchka",
			"a.a.a-a",
			1,
		},
	}
	invalidCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"get non-existed version",
			"Ahmed",
			"a.a.a-a",
			255,
		},
		{
			"get non-existed version",
			"Dima",
			"a.a.a-a",
			222,
		},
		{
			"get non-existed version",
			"Danyenchka",
			"a.a.a-a",
			123,
		},
	}

	storage := NewSessionStorage()

	for _, currentCase := range validCases {
		err := storage.Add(currentCase.login, currentCase.token, currentCase.version)
		if err != nil {
		}
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			version, _ := storage.GetVersion(currentCase.login, currentCase.token)
			if version != currentCase.version {
				t.Error("wrong version")
			}

			assert.Equal(t, true, reflect.DeepEqual(version, currentCase.version))
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			version, _ := storage.GetVersion(currentCase.login, currentCase.token)

			if version == currentCase.version {
				t.Error("right version")
			}
		})
	}
}

func TestCheckVersion(t *testing.T) {
	validCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"get existed version",
			"Ahmed",
			"a.a.a-a",
			123,
		},
		{
			"get existed version",
			"Dima",
			"a.a.a-a",
			255,
		},
		{
			"get existed version",
			"Danyenchka",
			"a.a.a-a",
			147,
		},
	}
	invalidCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"get non-existed version",
			"Ahmed",
			"a.a.a-a",
			255,
		},
		{
			"get non-existed version",
			"Dima",
			"a.a.a-a",
			222,
		},
		{
			"get non-existed version",
			"Danyenchka",
			"a.a.a-a",
			123,
		},
		{
			"get non-existed user",
			"Dimochka",
			"a.a.a-a",
			123,
		},
	}

	storage := NewSessionStorage()

	for _, currentCase := range validCases {
		err := storage.Add(currentCase.login, currentCase.token, currentCase.version)
		if err != nil {
		}
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			versionsAreSame, err := storage.CheckVersion(currentCase.login, currentCase.token, currentCase.version)
			if !versionsAreSame || err != nil {
				t.Error("something is wrong")
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			versionsAreSame, err := storage.CheckVersion(currentCase.login, currentCase.token, currentCase.version)

			if versionsAreSame || err == nil {
				t.Error("something is ok")
			}
		})
	}
}

func TestUpdateVersion(t *testing.T) {
	validCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"get existed version",
			"Ahmed",
			"a.a.a-a",
			254,
		},
		{
			"get existed version",
			"Dima",
			"a.a.a-a",
			155,
		},
		{
			"get existed version",
			"Danyenchka",
			"a.a.a-a",
			147,
		},
	}
	invalidCases := []struct {
		testName string
		login    string
		token    string
		version  uint32
	}{
		{
			"get non-existed version",
			"Ahmed",
			"a.a.a-a",
			255,
		},
	}

	storage := NewSessionStorage()

	for _, currentCase := range validCases {
		err := storage.Add(currentCase.login, currentCase.token, currentCase.version)
		if err != nil {
		}
	}

	for _, currentCase := range validCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.Update(currentCase.login, currentCase.token)
			version, _ := storage.GetVersion(currentCase.login, currentCase.token)

			if version-currentCase.version != 1 || err != nil {
				t.Error("something is wrong")
			}
		})
	}

	for _, currentCase := range invalidCases {
		t.Run(currentCase.testName, func(t *testing.T) {
			err := storage.Update(currentCase.login, currentCase.token)
			version, _ := storage.GetVersion(currentCase.login, currentCase.token)
			if version-currentCase.version == 1 || err == nil {
				t.Error("something is ok")
			}
		})
	}
}

func TestSessionStorage_CheckAllUserSessionTokens(t *testing.T) {
	storage := NewSessionStorage()

	// Set up some mock data in the cache
	mockCacheData := map[string]uint32{
		"token1": 1,
		"token2": 2,
	}
	mockLogin := "testUser"
	storage.cacheStorage.Set(mockLogin, mockCacheData, 0)

	err := storage.CheckAllUserSessionTokens(mockLogin)

	assert.NoError(t, err)
}
