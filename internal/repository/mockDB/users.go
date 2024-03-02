package mockdb

import (
	"errors"
	"sync"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
)

type MockDB struct {
	storage map[string]domain.User // key: login; value: password
	mutex   *sync.RWMutex
}

func InitMockDB() *MockDB {
	return &MockDB{
		storage: make(map[string]domain.User),
		mutex:   &sync.RWMutex{},
	}
}

func (db *MockDB) Create(user domain.User) error {
	db.mutex.RLock()
	_, ok := db.storage[user.Login]
	db.mutex.RUnlock()

	if ok {
		return errors.New("user already exists")
	}

	db.mutex.Lock()
	db.storage[user.Login] = user
	db.mutex.Unlock()

	return nil
}

func (db *MockDB) Remove(login string) error {
	db.mutex.RLock()
	_, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return errors.New("user doesn't exists")
	}

	db.mutex.Lock()
	delete(db.storage, login)
	db.mutex.Unlock()

	return nil
}

func (db *MockDB) HasUser(login, password string) error {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return errors.New("user doesn't exists")
	}

	if user.Password != password {
		return errors.New("incorrect password")
	}

	return nil
}

func (db *MockDB) GetUser(login string) (domain.User, error) {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return domain.User{}, errors.New("user doesn't exists")
	}

	return user, nil
}

func (db *MockDB) ChangePassword(login, newPassword string) error {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return errors.New("user doesn't exists")
	}

	db.mutex.Lock()
	user.Password = newPassword
	user.Version++
	db.storage[login] = user
	db.mutex.Unlock()

	return nil
}

func (db *MockDB) ChangeName(login, newName string) (domain.User, error) {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return domain.User{}, errors.New("user doesn't exists")
	}

	db.mutex.Lock()
	user.Name = newName
	user.Version++
	db.storage[login] = user
	db.mutex.Unlock()

	return user, nil
}
