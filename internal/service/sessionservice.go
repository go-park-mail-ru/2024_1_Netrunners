package service

import "go.uber.org/zap"

type sessionStorage interface {
	Add(login string, token string, version uint8) (err error)
	DeleteSession(login string, token string) (err error)
	Update(login string, token string) (err error)
	CheckVersion(login string, token string, usersVersion uint8) (hasSession bool, err error)
	GetVersion(login string, token string) (version uint8, err error)
	HasSession(login string, token string) error
	CheckAllUserSessionTokens(login string) error
}

type SessionService struct {
	sessionStorage sessionStorage
	logger         *zap.SugaredLogger
}

func NewSessionService(sessionStorage sessionStorage, logger *zap.SugaredLogger) *SessionService {
	return &SessionService{
		sessionStorage: sessionStorage,
		logger:         logger,
	}
}

func (service *SessionService) Add(login, token, requestID string, version uint8) (err error) {
	err = service.sessionStorage.Add(login, token, version)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at AddSession: %v", requestID, err)
		return err
	}
	return nil
}

func (service *SessionService) DeleteSession(login, token, requestID string) (err error) {
	err = service.sessionStorage.DeleteSession(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at DeleteSession: %v", requestID, err)
		return err
	}
	return nil
}

func (service *SessionService) Update(login, token, requestID string) (err error) {
	err = service.sessionStorage.Update(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at Update: %v", requestID, err)
		return err
	}
	return nil
}

func (service *SessionService) CheckVersion(login, token, requestID string,
	usersVersion uint8) (hasSession bool, err error) {
	hasSession, err = service.sessionStorage.CheckVersion(login, token, usersVersion)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at CheckVersion: %v", requestID, err)
		return hasSession, err
	}
	return hasSession, nil
}

func (service *SessionService) GetVersion(login, token, requestID string) (version uint8, err error) {
	version, err = service.sessionStorage.GetVersion(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at GetVersion: %v", requestID, err)
		return version, err
	}
	return version, nil
}

func (service *SessionService) HasSession(login, token, requestID string) (err error) {
	err = service.sessionStorage.HasSession(login, token)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at HasSession: %v", requestID, err)
		return err
	}
	return nil
}

func (service *SessionService) CheckAllUserSessionTokens(login, requestID string) error {
	err := service.sessionStorage.CheckAllUserSessionTokens(login)
	if err != nil {
		service.logger.Errorf("[reqid=%s] service error at CheckAllUserSessionTokens: %v", requestID, err)
		return err
	}
	return nil
}
