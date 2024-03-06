package mockdb

import (
	"fmt"
	"sync"

	"github.com/go-park-mail-ru/2024_1_Netrunners/internal/domain"
	myerrors "github.com/go-park-mail-ru/2024_1_Netrunners/internal/errors"
)

type UsersMockDB struct {
	storage map[string]domain.User // key: login; value: password
	mutex   *sync.RWMutex
}

func InitUsersMockDB() *UsersMockDB {
	return &UsersMockDB{
		storage: make(map[string]domain.User),
		mutex:   &sync.RWMutex{},
	}
}

func (db *UsersMockDB) CreateUser(user domain.User) error {
	db.mutex.RLock()
	_, ok := db.storage[user.Login]
	db.mutex.RUnlock()

	if ok {
		return fmt.Errorf("user already exists: %w", myerrors.ErrUserAlreadyExists)
	}

	db.mutex.Lock()
	db.storage[user.Login] = user
	db.mutex.Unlock()

	return nil
}

func (db *UsersMockDB) RemoveUser(login string) error {
	db.mutex.RLock()
	_, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return fmt.Errorf("user doesn't exists: %w", myerrors.ErrNoSuchUser)
	}

	db.mutex.Lock()
	delete(db.storage, login)
	db.mutex.Unlock()

	return nil
}

func (db *UsersMockDB) HasUser(login, password string) error {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return fmt.Errorf("user doesn't exists: %w", myerrors.ErrNoSuchUser)
	}

	if user.Password != password {
		return fmt.Errorf("incorrect password: %w", myerrors.ErrIncorrectLoginOrPassword)
	}

	return nil
}

func (db *UsersMockDB) GetUser(login string) (domain.User, error) {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	user.Password = ""

	if !ok {
		return domain.User{}, fmt.Errorf("user doesn't exists: %w", myerrors.ErrNoSuchUser)
	}

	return user, nil
}

func (db *UsersMockDB) ChangeUserPassword(login, newPassword string) error {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return fmt.Errorf("user doesn't exists: %w", myerrors.ErrNoSuchUser)
	}

	db.mutex.Lock()
	user.Password = newPassword
	user.Version++
	db.storage[login] = user
	db.mutex.Unlock()

	return nil
}

func (db *UsersMockDB) ChangeUserName(login, newName string) (domain.User, error) {
	db.mutex.RLock()
	user, ok := db.storage[login]
	db.mutex.RUnlock()

	if !ok {
		return domain.User{}, fmt.Errorf("user doesn't exists: %w", myerrors.ErrNoSuchUser)
	}

	db.mutex.Lock()
	user.Name = newName
	user.Version++
	db.storage[login] = user
	db.mutex.Unlock()

	user.Password = ""

	return user, nil
}
